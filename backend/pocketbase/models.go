package pocketbase

import "github.com/pocketbase/pocketbase/core"

var (
	_ core.RecordProxy = (*VoiceChannelGeneratorConfig)(nil)
	_ core.RecordProxy = (*VoiceChannelGeneratorChannel)(nil)
	_ core.RecordProxy = (*VoiceRoleConfig)(nil)
)

type VoiceChannelGeneratorConfig struct {
	core.BaseRecordProxy
}

func (v *VoiceChannelGeneratorConfig) GetNameFormat() string {
	return v.GetString("nameFormat")
}

func (v *VoiceChannelGeneratorConfig) GetChannelID() string {
	return v.GetString("channelID")
}

func (v *VoiceChannelGeneratorConfig) GetServerID() string {
	return v.GetString("serverID")
}

type VoiceChannelGeneratorChannel struct {
	core.BaseRecordProxy
}

func (v *VoiceChannelGeneratorChannel) GetIndex() int {
	return v.GetInt("index")
}

func (v *VoiceChannelGeneratorChannel) SetIndex(i int) {
	v.Set("index", i)
}

func (v *VoiceChannelGeneratorChannel) SetChannelID(id string) {
	v.Set("channelID", id)
}

func (v *VoiceChannelGeneratorChannel) SetOwnerID(id string) {
	v.Set("ownerID", id)
}

func (v *VoiceChannelGeneratorChannel) SetParentConfig(config *VoiceChannelGeneratorConfig) {
	v.Set("parentConfig", config.Id)
}

func (v *VoiceChannelGeneratorChannel) SetName(name string) {
	v.Set("name", name)
}

func (v *VoiceChannelGeneratorChannel) GetName() string {
	return v.GetString("name")
}

type VoiceRoleConfig struct {
	core.BaseRecordProxy
}
