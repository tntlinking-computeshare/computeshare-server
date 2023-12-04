// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/agent"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/computeimage"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/computeinstance"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/computespec"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/domainbinding"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/gateway"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/gatewayport"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/networkmapping"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/s3bucket"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/s3user"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/schema"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/script"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/scriptexecutionrecord"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/storage"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/task"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	agentFields := schema.Agent{}.Fields()
	_ = agentFields
	// agentDescPeerID is the schema descriptor for peer_id field.
	agentDescPeerID := agentFields[1].Descriptor()
	// agent.PeerIDValidator is a validator for the "peer_id" field. It is called by the builders before save.
	agent.PeerIDValidator = agentDescPeerID.Validators[0].(func(string) error)
	// agentDescActive is the schema descriptor for active field.
	agentDescActive := agentFields[2].Descriptor()
	// agent.DefaultActive holds the default value on creation for the active field.
	agent.DefaultActive = agentDescActive.Default.(bool)
	// agentDescLastUpdateTime is the schema descriptor for last_update_time field.
	agentDescLastUpdateTime := agentFields[3].Descriptor()
	// agent.DefaultLastUpdateTime holds the default value on creation for the last_update_time field.
	agent.DefaultLastUpdateTime = agentDescLastUpdateTime.Default.(func() time.Time)
	// agent.UpdateDefaultLastUpdateTime holds the default value on update for the last_update_time field.
	agent.UpdateDefaultLastUpdateTime = agentDescLastUpdateTime.UpdateDefault.(func() time.Time)
	// agentDescID is the schema descriptor for id field.
	agentDescID := agentFields[0].Descriptor()
	// agent.DefaultID holds the default value on creation for the id field.
	agent.DefaultID = agentDescID.Default.(func() uuid.UUID)
	computeimageFields := schema.ComputeImage{}.Fields()
	_ = computeimageFields
	// computeimageDescName is the schema descriptor for name field.
	computeimageDescName := computeimageFields[1].Descriptor()
	// computeimage.NameValidator is a validator for the "name" field. It is called by the builders before save.
	computeimage.NameValidator = computeimageDescName.Validators[0].(func(string) error)
	// computeimageDescImage is the schema descriptor for image field.
	computeimageDescImage := computeimageFields[2].Descriptor()
	// computeimage.ImageValidator is a validator for the "image" field. It is called by the builders before save.
	computeimage.ImageValidator = computeimageDescImage.Validators[0].(func(string) error)
	// computeimageDescTag is the schema descriptor for tag field.
	computeimageDescTag := computeimageFields[3].Descriptor()
	// computeimage.TagValidator is a validator for the "tag" field. It is called by the builders before save.
	computeimage.TagValidator = computeimageDescTag.Validators[0].(func(string) error)
	computeinstanceFields := schema.ComputeInstance{}.Fields()
	_ = computeinstanceFields
	// computeinstanceDescOwner is the schema descriptor for owner field.
	computeinstanceDescOwner := computeinstanceFields[1].Descriptor()
	// computeinstance.OwnerValidator is a validator for the "owner" field. It is called by the builders before save.
	computeinstance.OwnerValidator = computeinstanceDescOwner.Validators[0].(func(string) error)
	// computeinstanceDescName is the schema descriptor for name field.
	computeinstanceDescName := computeinstanceFields[2].Descriptor()
	// computeinstance.NameValidator is a validator for the "name" field. It is called by the builders before save.
	computeinstance.NameValidator = computeinstanceDescName.Validators[0].(func(string) error)
	// computeinstanceDescCore is the schema descriptor for core field.
	computeinstanceDescCore := computeinstanceFields[3].Descriptor()
	// computeinstance.CoreValidator is a validator for the "core" field. It is called by the builders before save.
	computeinstance.CoreValidator = computeinstanceDescCore.Validators[0].(func(string) error)
	// computeinstanceDescMemory is the schema descriptor for memory field.
	computeinstanceDescMemory := computeinstanceFields[4].Descriptor()
	// computeinstance.MemoryValidator is a validator for the "memory" field. It is called by the builders before save.
	computeinstance.MemoryValidator = computeinstanceDescMemory.Validators[0].(func(string) error)
	// computeinstanceDescImage is the schema descriptor for image field.
	computeinstanceDescImage := computeinstanceFields[5].Descriptor()
	// computeinstance.ImageValidator is a validator for the "image" field. It is called by the builders before save.
	computeinstance.ImageValidator = computeinstanceDescImage.Validators[0].(func(string) error)
	// computeinstanceDescID is the schema descriptor for id field.
	computeinstanceDescID := computeinstanceFields[0].Descriptor()
	// computeinstance.DefaultID holds the default value on creation for the id field.
	computeinstance.DefaultID = computeinstanceDescID.Default.(func() uuid.UUID)
	computespecFields := schema.ComputeSpec{}.Fields()
	_ = computespecFields
	// computespecDescCore is the schema descriptor for core field.
	computespecDescCore := computespecFields[1].Descriptor()
	// computespec.CoreValidator is a validator for the "core" field. It is called by the builders before save.
	computespec.CoreValidator = computespecDescCore.Validators[0].(func(string) error)
	// computespecDescMemory is the schema descriptor for memory field.
	computespecDescMemory := computespecFields[2].Descriptor()
	// computespec.MemoryValidator is a validator for the "memory" field. It is called by the builders before save.
	computespec.MemoryValidator = computespecDescMemory.Validators[0].(func(string) error)
	domainbindingFields := schema.DomainBinding{}.Fields()
	_ = domainbindingFields
	// domainbindingDescName is the schema descriptor for name field.
	domainbindingDescName := domainbindingFields[4].Descriptor()
	// domainbinding.NameValidator is a validator for the "name" field. It is called by the builders before save.
	domainbinding.NameValidator = func() func(string) error {
		validators := domainbindingDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// domainbindingDescDomain is the schema descriptor for domain field.
	domainbindingDescDomain := domainbindingFields[5].Descriptor()
	// domainbinding.DomainValidator is a validator for the "domain" field. It is called by the builders before save.
	domainbinding.DomainValidator = domainbindingDescDomain.Validators[0].(func(string) error)
	// domainbindingDescID is the schema descriptor for id field.
	domainbindingDescID := domainbindingFields[0].Descriptor()
	// domainbinding.DefaultID holds the default value on creation for the id field.
	domainbinding.DefaultID = domainbindingDescID.Default.(func() uuid.UUID)
	gatewayFields := schema.Gateway{}.Fields()
	_ = gatewayFields
	// gatewayDescName is the schema descriptor for name field.
	gatewayDescName := gatewayFields[1].Descriptor()
	// gateway.NameValidator is a validator for the "name" field. It is called by the builders before save.
	gateway.NameValidator = func() func(string) error {
		validators := gatewayDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// gatewayDescID is the schema descriptor for id field.
	gatewayDescID := gatewayFields[0].Descriptor()
	// gateway.DefaultID holds the default value on creation for the id field.
	gateway.DefaultID = gatewayDescID.Default.(func() uuid.UUID)
	gatewayportFields := schema.GatewayPort{}.Fields()
	_ = gatewayportFields
	// gatewayportDescIsUse is the schema descriptor for is_use field.
	gatewayportDescIsUse := gatewayportFields[3].Descriptor()
	// gatewayport.DefaultIsUse holds the default value on creation for the is_use field.
	gatewayport.DefaultIsUse = gatewayportDescIsUse.Default.(bool)
	// gatewayportDescID is the schema descriptor for id field.
	gatewayportDescID := gatewayportFields[0].Descriptor()
	// gatewayport.DefaultID holds the default value on creation for the id field.
	gatewayport.DefaultID = gatewayportDescID.Default.(func() uuid.UUID)
	networkmappingFields := schema.NetworkMapping{}.Fields()
	_ = networkmappingFields
	// networkmappingDescName is the schema descriptor for name field.
	networkmappingDescName := networkmappingFields[1].Descriptor()
	// networkmapping.NameValidator is a validator for the "name" field. It is called by the builders before save.
	networkmapping.NameValidator = func() func(string) error {
		validators := networkmappingDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// networkmappingDescStatus is the schema descriptor for status field.
	networkmappingDescStatus := networkmappingFields[5].Descriptor()
	// networkmapping.DefaultStatus holds the default value on creation for the status field.
	networkmapping.DefaultStatus = networkmappingDescStatus.Default.(int)
	// networkmappingDescID is the schema descriptor for id field.
	networkmappingDescID := networkmappingFields[0].Descriptor()
	// networkmapping.DefaultID holds the default value on creation for the id field.
	networkmapping.DefaultID = networkmappingDescID.Default.(func() uuid.UUID)
	s3bucketFields := schema.S3Bucket{}.Fields()
	_ = s3bucketFields
	// s3bucketDescBucket is the schema descriptor for bucket field.
	s3bucketDescBucket := s3bucketFields[1].Descriptor()
	// s3bucket.BucketValidator is a validator for the "bucket" field. It is called by the builders before save.
	s3bucket.BucketValidator = s3bucketDescBucket.Validators[0].(func(string) error)
	// s3bucketDescID is the schema descriptor for id field.
	s3bucketDescID := s3bucketFields[0].Descriptor()
	// s3bucket.DefaultID holds the default value on creation for the id field.
	s3bucket.DefaultID = s3bucketDescID.Default.(func() uuid.UUID)
	s3userFields := schema.S3User{}.Fields()
	_ = s3userFields
	// s3userDescAccessKey is the schema descriptor for access_key field.
	s3userDescAccessKey := s3userFields[2].Descriptor()
	// s3user.AccessKeyValidator is a validator for the "access_key" field. It is called by the builders before save.
	s3user.AccessKeyValidator = s3userDescAccessKey.Validators[0].(func(string) error)
	// s3userDescSecretKey is the schema descriptor for secret_key field.
	s3userDescSecretKey := s3userFields[3].Descriptor()
	// s3user.SecretKeyValidator is a validator for the "secret_key" field. It is called by the builders before save.
	s3user.SecretKeyValidator = s3userDescSecretKey.Validators[0].(func(string) error)
	// s3userDescID is the schema descriptor for id field.
	s3userDescID := s3userFields[0].Descriptor()
	// s3user.DefaultID holds the default value on creation for the id field.
	s3user.DefaultID = s3userDescID.Default.(func() uuid.UUID)
	scriptFields := schema.Script{}.Fields()
	_ = scriptFields
	// scriptDescTaskNumber is the schema descriptor for task_number field.
	scriptDescTaskNumber := scriptFields[2].Descriptor()
	// script.TaskNumberValidator is a validator for the "task_number" field. It is called by the builders before save.
	script.TaskNumberValidator = scriptDescTaskNumber.Validators[0].(func(int32) error)
	// scriptDescScriptName is the schema descriptor for script_name field.
	scriptDescScriptName := scriptFields[3].Descriptor()
	// script.ScriptNameValidator is a validator for the "script_name" field. It is called by the builders before save.
	script.ScriptNameValidator = scriptDescScriptName.Validators[0].(func(string) error)
	// scriptDescScriptContent is the schema descriptor for script_content field.
	scriptDescScriptContent := scriptFields[5].Descriptor()
	// script.ScriptContentValidator is a validator for the "script_content" field. It is called by the builders before save.
	script.ScriptContentValidator = scriptDescScriptContent.Validators[0].(func(string) error)
	// scriptDescCreateTime is the schema descriptor for create_time field.
	scriptDescCreateTime := scriptFields[6].Descriptor()
	// script.DefaultCreateTime holds the default value on creation for the create_time field.
	script.DefaultCreateTime = scriptDescCreateTime.Default.(time.Time)
	// scriptDescUpdateTime is the schema descriptor for update_time field.
	scriptDescUpdateTime := scriptFields[7].Descriptor()
	// script.DefaultUpdateTime holds the default value on creation for the update_time field.
	script.DefaultUpdateTime = scriptDescUpdateTime.Default.(time.Time)
	scriptexecutionrecordFields := schema.ScriptExecutionRecord{}.Fields()
	_ = scriptexecutionrecordFields
	// scriptexecutionrecordDescFkScriptID is the schema descriptor for fk_script_id field.
	scriptexecutionrecordDescFkScriptID := scriptexecutionrecordFields[2].Descriptor()
	// scriptexecutionrecord.FkScriptIDValidator is a validator for the "fk_script_id" field. It is called by the builders before save.
	scriptexecutionrecord.FkScriptIDValidator = scriptexecutionrecordDescFkScriptID.Validators[0].(func(int32) error)
	// scriptexecutionrecordDescScriptContent is the schema descriptor for script_content field.
	scriptexecutionrecordDescScriptContent := scriptexecutionrecordFields[3].Descriptor()
	// scriptexecutionrecord.ScriptContentValidator is a validator for the "script_content" field. It is called by the builders before save.
	scriptexecutionrecord.ScriptContentValidator = scriptexecutionrecordDescScriptContent.Validators[0].(func(string) error)
	// scriptexecutionrecordDescTaskNumber is the schema descriptor for task_number field.
	scriptexecutionrecordDescTaskNumber := scriptexecutionrecordFields[4].Descriptor()
	// scriptexecutionrecord.TaskNumberValidator is a validator for the "task_number" field. It is called by the builders before save.
	scriptexecutionrecord.TaskNumberValidator = scriptexecutionrecordDescTaskNumber.Validators[0].(func(int32) error)
	// scriptexecutionrecordDescScriptName is the schema descriptor for script_name field.
	scriptexecutionrecordDescScriptName := scriptexecutionrecordFields[5].Descriptor()
	// scriptexecutionrecord.ScriptNameValidator is a validator for the "script_name" field. It is called by the builders before save.
	scriptexecutionrecord.ScriptNameValidator = scriptexecutionrecordDescScriptName.Validators[0].(func(string) error)
	// scriptexecutionrecordDescCreateTime is the schema descriptor for create_time field.
	scriptexecutionrecordDescCreateTime := scriptexecutionrecordFields[9].Descriptor()
	// scriptexecutionrecord.DefaultCreateTime holds the default value on creation for the create_time field.
	scriptexecutionrecord.DefaultCreateTime = scriptexecutionrecordDescCreateTime.Default.(time.Time)
	// scriptexecutionrecordDescUpdateTime is the schema descriptor for update_time field.
	scriptexecutionrecordDescUpdateTime := scriptexecutionrecordFields[10].Descriptor()
	// scriptexecutionrecord.DefaultUpdateTime holds the default value on creation for the update_time field.
	scriptexecutionrecord.DefaultUpdateTime = scriptexecutionrecordDescUpdateTime.Default.(time.Time)
	storageFields := schema.Storage{}.Fields()
	_ = storageFields
	// storageDescOwner is the schema descriptor for owner field.
	storageDescOwner := storageFields[1].Descriptor()
	// storage.OwnerValidator is a validator for the "owner" field. It is called by the builders before save.
	storage.OwnerValidator = func() func(string) error {
		validators := storageDescOwner.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(owner string) error {
			for _, fn := range fns {
				if err := fn(owner); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// storageDescType is the schema descriptor for type field.
	storageDescType := storageFields[2].Descriptor()
	// storage.DefaultType holds the default value on creation for the type field.
	storage.DefaultType = storageDescType.Default.(int32)
	// storageDescName is the schema descriptor for name field.
	storageDescName := storageFields[3].Descriptor()
	// storage.NameValidator is a validator for the "name" field. It is called by the builders before save.
	storage.NameValidator = func() func(string) error {
		validators := storageDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// storageDescCid is the schema descriptor for cid field.
	storageDescCid := storageFields[4].Descriptor()
	// storage.CidValidator is a validator for the "cid" field. It is called by the builders before save.
	storage.CidValidator = storageDescCid.Validators[0].(func(string) error)
	// storageDescLastModify is the schema descriptor for last_modify field.
	storageDescLastModify := storageFields[6].Descriptor()
	// storage.DefaultLastModify holds the default value on creation for the last_modify field.
	storage.DefaultLastModify = storageDescLastModify.Default.(func() time.Time)
	// storageDescParentID is the schema descriptor for parent_id field.
	storageDescParentID := storageFields[7].Descriptor()
	// storage.ParentIDValidator is a validator for the "parent_id" field. It is called by the builders before save.
	storage.ParentIDValidator = storageDescParentID.Validators[0].(func(string) error)
	// storageDescID is the schema descriptor for id field.
	storageDescID := storageFields[0].Descriptor()
	// storage.DefaultID holds the default value on creation for the id field.
	storage.DefaultID = storageDescID.Default.(func() uuid.UUID)
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescAgentID is the schema descriptor for agent_id field.
	taskDescAgentID := taskFields[1].Descriptor()
	// task.AgentIDValidator is a validator for the "agent_id" field. It is called by the builders before save.
	task.AgentIDValidator = func() func(string) error {
		validators := taskDescAgentID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(agent_id string) error {
			for _, fn := range fns {
				if err := fn(agent_id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// taskDescCmd is the schema descriptor for cmd field.
	taskDescCmd := taskFields[2].Descriptor()
	// task.DefaultCmd holds the default value on creation for the cmd field.
	task.DefaultCmd = taskDescCmd.Default.(int32)
	// taskDescParams is the schema descriptor for params field.
	taskDescParams := taskFields[3].Descriptor()
	// task.ParamsValidator is a validator for the "params" field. It is called by the builders before save.
	task.ParamsValidator = taskDescParams.Validators[0].(func(string) error)
	// taskDescCreateTime is the schema descriptor for create_time field.
	taskDescCreateTime := taskFields[5].Descriptor()
	// task.DefaultCreateTime holds the default value on creation for the create_time field.
	task.DefaultCreateTime = taskDescCreateTime.Default.(func() time.Time)
	// taskDescID is the schema descriptor for id field.
	taskDescID := taskFields[0].Descriptor()
	// task.DefaultID holds the default value on creation for the id field.
	task.DefaultID = taskDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCountryCallCoding is the schema descriptor for country_call_coding field.
	userDescCountryCallCoding := userFields[1].Descriptor()
	// user.CountryCallCodingValidator is a validator for the "country_call_coding" field. It is called by the builders before save.
	user.CountryCallCodingValidator = func() func(string) error {
		validators := userDescCountryCallCoding.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(country_call_coding string) error {
			for _, fn := range fns {
				if err := fn(country_call_coding); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescTelephoneNumber is the schema descriptor for telephone_number field.
	userDescTelephoneNumber := userFields[2].Descriptor()
	// user.TelephoneNumberValidator is a validator for the "telephone_number" field. It is called by the builders before save.
	user.TelephoneNumberValidator = func() func(string) error {
		validators := userDescTelephoneNumber.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(telephone_number string) error {
			for _, fn := range fns {
				if err := fn(telephone_number); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[3].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescCreateDate is the schema descriptor for create_date field.
	userDescCreateDate := userFields[4].Descriptor()
	// user.DefaultCreateDate holds the default value on creation for the create_date field.
	user.DefaultCreateDate = userDescCreateDate.Default.(func() time.Time)
	// userDescLastLoginDate is the schema descriptor for last_login_date field.
	userDescLastLoginDate := userFields[5].Descriptor()
	// user.DefaultLastLoginDate holds the default value on creation for the last_login_date field.
	user.DefaultLastLoginDate = userDescLastLoginDate.Default.(func() time.Time)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[6].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescPwdConfig is the schema descriptor for pwd_config field.
	userDescPwdConfig := userFields[8].Descriptor()
	// user.DefaultPwdConfig holds the default value on creation for the pwd_config field.
	user.DefaultPwdConfig = userDescPwdConfig.Default.(bool)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
