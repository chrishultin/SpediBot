package pocketbase

import (
	"fmt"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

const (
	NO_RESULTS_ERR = "sql: no rows in result set"

	COLLECTION_CHANNEL_GENERATOR_CONFIG  = "channel_generator_configs"
	COLLECTION_CHANNEL_GENERATOR_CHANNEL = "channel_generator_channels"
)

type Client struct {
	PocketBase *pocketbase.PocketBase
}

func (c *Client) ChannelGeneratorConfigForChannelID(channelID string) (*VoiceChannelGeneratorConfig, error) {
	conf := &VoiceChannelGeneratorConfig{}

	err := c.PocketBase.RecordQuery(COLLECTION_CHANNEL_GENERATOR_CONFIG).
		AndWhere(dbx.NewExp("channelID={:channelID}", dbx.Params{"channelID": channelID})).
		Limit(1).One(conf)

	if err != nil && err.Error() == NO_RESULTS_ERR {
		return nil, nil
	}

	return conf, err
}

func (c *Client) NextFreeChannelGeneratorIndex(channelID string) (int, error) {
	channels, err := c.PocketBase.FindRecordsByFilter(COLLECTION_CHANNEL_GENERATOR_CHANNEL,
		fmt.Sprintf("parentConfig.channelID = %q", channelID), "index", 0, 0)

	if err != nil {
		if err.Error() == NO_RESULTS_ERR {
			return 1, nil
		}
	}

	fmt.Printf("Channels: %+v\n", channels)

	if len(channels) == 0 {
		return 1, nil
	}

	channelIndexMap := map[int]bool{}
	maxIndex := channels[len(channels)-1].GetInt("index")
	for _, channel := range channels {
		channelIndexMap[channel.GetInt("index")] = true
	}

	for i := range maxIndex {
		if i == 0 {
			continue
		}
		if _, ok := channelIndexMap[i]; !ok {
			return i, nil
		}
	}

	return maxIndex + 1, nil
}

func (c *Client) SaveChannelGeneratorChannel(channelID, ownerID string, index int, parentConfig, name string) error {
	collection, err := c.PocketBase.FindCollectionByNameOrId(COLLECTION_CHANNEL_GENERATOR_CHANNEL)
	if err != nil {
		return err
	}
	record := core.NewRecord(collection)

	record.Load(map[string]any{
		"channelID":    channelID,
		"ownerID":      ownerID,
		"index":        index,
		"parentConfig": parentConfig,
		"name":         name,
	})

	return c.PocketBase.Save(record)
}

func (c *Client) TransferChannelGeneratorChannelOwner(channel *VoiceChannelGeneratorChannel, newOwnerID, newName string) error {
	channel.SetOwnerID(newOwnerID)
	channel.SetName(newName)

	return c.PocketBase.Save(channel)
}

func (c *Client) RemoveChannelGeneratorChannel(channelID string) error {
	record, err := c.PocketBase.FindFirstRecordByData(COLLECTION_CHANNEL_GENERATOR_CHANNEL, "channelID", channelID)
	if err != nil {
		return err
	}

	return c.PocketBase.Delete(record)
}

func (c *Client) OwnedChannelGeneratorChannel(channelID, ownerID string) (*VoiceChannelGeneratorChannel, error) {
	channel := &VoiceChannelGeneratorChannel{}
	err := c.PocketBase.RecordQuery(COLLECTION_CHANNEL_GENERATOR_CHANNEL).
		AndWhere(dbx.NewExp("ownerID={:ownerID}", dbx.Params{"ownerID": ownerID})).
		AndWhere(dbx.NewExp("channelID={:channelID}", dbx.Params{"channelID": channelID})).
		Limit(1).One(channel)

	if err != nil {
		if err.Error() != NO_RESULTS_ERR {
			return nil, nil
		}

		return nil, err
	}

	c.PocketBase.ExpandRecord(
		channel.Record,
		[]string{"parentConfig"},
		nil,
	)

	return channel, nil
}
