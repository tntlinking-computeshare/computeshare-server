package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ComputeImage holds the schema definition for the ComputeImage entity.
type ComputeImage struct {
	ent.Schema
}

/*
	downloadFiles["ubuntu:20.04"] = Image{
		Name:        "ubuntu-20.04",
		Filename:    "ubuntu-20.04.qcow2.bak",
		DownloadUrl: "https://g.alpha.hamsternet.io/ipfs/QmZnCDgtSBQzHTyv2Ksku4zAxq9t7yUJwWGHUZAj2oX4AB?filename=ubuntu-20.04.qcow2.bak",
		MD5:         "f0432ad697f5762c28980a397c4e8d60",
		OsType:      "linux",
		OsVariant:   "ubuntu20.04",
	}
*/

// Fields of the ComputeImage.
func (ComputeImage) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.String("name").NotEmpty().Comment("显示名"),          // 废弃
		field.String("image").NotEmpty().Comment("镜像名"),         // 废弃
		field.String("tag").NotEmpty().Comment("版本名"),           // 废弃
		field.String("os_type").NotEmpty().Comment("操作系统类型"),    // linux/windows
		field.String("os_variant").NotEmpty().Comment("操作系统版本"), // ubuntu20.04 windows2k19
		field.String("filename").Comment("镜像文件名"),
		field.String("download_url").Comment("镜像下载地址"),
		field.String("md5").Comment("镜像md5"),
	}
}

// Edges of the ComputeImage.
func (ComputeImage) Edges() []ent.Edge {
	return nil
}

func (ComputeImage) Indexes() []ent.Index {
	return []ent.Index{
		// 索引
		index.Fields("id").Unique(),
	}
}
