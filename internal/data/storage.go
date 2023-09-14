package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	pb "github.com/mohaijiang/computeshare-server/api/compute/v1"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/storage"
	"path/filepath"
	"regexp"
	"strconv"
)

type storageRepo struct {
	data *Data
	log  *log.Helper
}

func NewStorageRepo(data *Data, logger log.Logger) biz.StorageRepo {
	return &storageRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ur *storageRepo) ListStorage(ctx context.Context, owner string, parentId string) ([]*biz.Storage, error) {
	parentIdPredicate := storage.ParentID(parentId)
	ps, err := ur.data.db.Storage.Query().
		Where(storage.OwnerEQ(owner), parentIdPredicate).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Storage, 0)
	for _, p := range ps {
		rv = append(rv, &biz.Storage{
			ID:         p.ID,
			Name:       p.Name,
			Type:       p.Type,
			Owner:      p.Owner,
			Cid:        *p.Cid,
			LastModify: p.LastModify,
			ParentID:   parentId,
		})
	}
	return rv, nil
}
func (ur *storageRepo) GetStorage(ctx context.Context, id uuid.UUID) (*biz.Storage, error) {
	p, err := ur.data.db.Storage.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &biz.Storage{
		ID:         p.ID,
		Name:       p.Name,
		Type:       p.Type,
		Owner:      p.Owner,
		Cid:        *p.Cid,
		LastModify: p.LastModify,
		ParentID:   p.ParentID,
	}, nil
}

func (ur *storageRepo) CreateStorage(ctx context.Context, entity *biz.Storage) error {

	// 判断有该目录有无重复的文件夹
	exists := ur.data.db.Storage.Query().Where(
		storage.OwnerEQ(entity.Owner),
		storage.ParentIDEQ(entity.ParentID),
		storage.NameEQ(entity.Name),
		storage.TypeEQ(entity.Type),
	).ExistX(ctx)

	if exists && entity.Type == int32(pb.FileType_DIR) {
		return fmt.Errorf("%s dir is exists", entity.Name)
	}
	name := entity.Name
	if exists && entity.Type == int32(pb.FileType_FILE) {
		for {
			if !exists {
				break
			}
			// 使用filepath包的Ext来获取文件的后缀
			fileExt := filepath.Ext(name)

			// 使用filepath包的Base来获取不包含后缀的文件名
			fileName := name[:len(name)-len(fileExt)]
			// 使用正则表达式来匹配以(N)结尾的字符串
			re := regexp.MustCompile(`\((\d+)\)$`)
			match := re.FindStringSubmatch(fileName)

			if len(match) == 2 {
				// 第一个匹配项是整数N的字符串表示
				nStr := match[1]

				// 将nStr解析为整数
				n, err := strconv.Atoi(nStr)
				if err != nil {
					fmt.Println("无法将字符串转换为整数:", err)
					fileName = fmt.Sprintf("%s(%d)", re.ReplaceAllString(fileName, ""), n+1)
				} else {
					fmt.Println("找到匹配，N =", n)
					fileName = re.ReplaceAllString(fileName, fmt.Sprintf("(%d)", n+1))
				}
			} else {
				fmt.Println("没有找到匹配")
				fileName = fmt.Sprintf("%s(1)", fileName)
			}

			fmt.Println("文件名前缀:", fileName)
			fmt.Println("文件后缀:", fileExt)
			name = fmt.Sprintf("%s%s", fileName, fileExt)
			exists = ur.data.db.Storage.Query().Where(
				storage.OwnerEQ(entity.Owner),
				storage.ParentIDEQ(entity.ParentID),
				storage.NameEQ(name),
				storage.TypeEQ(entity.Type),
			).ExistX(ctx)
		}

	}

	client := ur.data.db.Storage.
		Create().
		SetOwner(entity.Owner).
		SetName(name).
		SetSize(entity.Size).
		SetType(entity.Type).
		SetLastModify(entity.LastModify).
		SetCid(entity.Cid).
		SetParentID(entity.ParentID)

	result, err := client.Save(ctx)

	if err != nil {
		return err
	}

	entity.ID = result.ID
	entity.Name = name
	return err
}
func (ur *storageRepo) UpdateStorage(ctx context.Context, id uuid.UUID, storage *biz.Storage) error {
	p, err := ur.data.db.Storage.Get(ctx, id)
	if err != nil {
		return err
	}
	_, err = p.Update().
		SetOwner(storage.Owner).
		SetName(storage.Name).
		SetCid(storage.Cid).
		SetType(storage.Type).
		SetLastModify(storage.LastModify).
		SetParentID(storage.ParentID).
		Save(ctx)
	return err
}
func (ur *storageRepo) DeleteStorage(ctx context.Context, id uuid.UUID) error {
	result, err := ur.data.db.Storage.Get(ctx, id)
	if err != nil {
		return err
	}

	children, err := ur.data.db.Storage.Query().Where(storage.ParentID(result.ID.String())).All(ctx)
	if err != nil {
		return err
	}

	for _, node := range children {
		err = ur.DeleteStorage(ctx, node.ID)
		if err != nil {
			return err
		}
	}
	return ur.data.db.Storage.DeleteOneID(id).Exec(ctx)
}
