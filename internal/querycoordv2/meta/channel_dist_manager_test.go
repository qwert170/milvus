package meta

import (
	"testing"

	"github.com/milvus-io/milvus/internal/proto/datapb"
	"github.com/milvus-io/milvus/internal/util/typeutil"
	"github.com/stretchr/testify/suite"
)

type ChannelDistManagerSuite struct {
	suite.Suite
	dist       *ChannelDistManager
	collection int64
	nodes      []int64
	channels   map[string]*DmChannel
}

func (suite *ChannelDistManagerSuite) SetupSuite() {
	// Replica 0: 0, 2
	// Replica 1: 1
	suite.collection = 10
	suite.nodes = []int64{0, 1, 2}
	suite.channels = map[string]*DmChannel{
		"dmc0": DmChannelFromVChannel(&datapb.VchannelInfo{
			CollectionID: suite.collection,
			ChannelName:  "dmc0",
		}),
		"dmc1": DmChannelFromVChannel(&datapb.VchannelInfo{
			CollectionID: suite.collection,
			ChannelName:  "dmc1",
		}),
	}
}

func (suite *ChannelDistManagerSuite) SetupTest() {
	suite.dist = NewChannelDistManager()
	// Distribution:
	// node 0 contains channel dmc0
	// node 1 contains channel dmc0, dmc1
	// node 2 contains channel dmc1
	suite.dist.Update(suite.nodes[0], suite.channels["dmc0"].Clone())
	suite.dist.Update(suite.nodes[1], suite.channels["dmc0"].Clone(), suite.channels["dmc1"].Clone())
	suite.dist.Update(suite.nodes[2], suite.channels["dmc1"].Clone())
}

func (suite *ChannelDistManagerSuite) TestGetBy() {
	dist := suite.dist

	// Test GetAll
	channels := dist.GetAll()
	suite.Len(channels, 4)

	// Test GetByNode
	for _, node := range suite.nodes {
		channels := dist.GetByNode(node)
		suite.AssertNode(channels, node)
	}

	// Test GetByCollection
	channels = dist.GetByCollection(suite.collection)
	suite.Len(channels, 4)
	suite.AssertCollection(channels, suite.collection)
	channels = dist.GetByCollection(-1)
	suite.Len(channels, 0)

	// Test GetByNodeAndCollection
	// 1. Valid node and valid collection
	for _, node := range suite.nodes {
		channels := dist.GetByCollectionAndNode(suite.collection, node)
		suite.AssertNode(channels, node)
		suite.AssertCollection(channels, suite.collection)
	}

	// 2. Valid node and invalid collection
	channels = dist.GetByCollectionAndNode(-1, suite.nodes[1])
	suite.Len(channels, 0)

	// 3. Invalid node and valid collection
	channels = dist.GetByCollectionAndNode(suite.collection, -1)
	suite.Len(channels, 0)
}

func (suite *ChannelDistManagerSuite) TestGetShardLeader() {
	replicas := []*Replica{
		{Nodes: typeutil.NewUniqueSet(suite.nodes[0], suite.nodes[2])},
		{Nodes: typeutil.NewUniqueSet(suite.nodes[1])},
	}

	// Test on replica 0
	leader0, ok := suite.dist.GetShardLeader(replicas[0], "dmc0")
	suite.True(ok)
	suite.Equal(suite.nodes[0], leader0)
	leader1, ok := suite.dist.GetShardLeader(replicas[0], "dmc1")
	suite.True(ok)
	suite.Equal(suite.nodes[2], leader1)

	// Test on replica 1
	leader0, ok = suite.dist.GetShardLeader(replicas[1], "dmc0")
	suite.True(ok)
	suite.Equal(suite.nodes[1], leader0)
	leader1, ok = suite.dist.GetShardLeader(replicas[1], "dmc1")
	suite.True(ok)
	suite.Equal(suite.nodes[1], leader1)

	// Test no shard leader for given channel
	_, ok = suite.dist.GetShardLeader(replicas[0], "invalid-shard")
	suite.False(ok)

	// Test on replica 0
	leaders := suite.dist.GetShardLeadersByReplica(replicas[0])
	suite.Len(leaders, 2)
	suite.Equal(leaders["dmc0"], suite.nodes[0])
	suite.Equal(leaders["dmc1"], suite.nodes[2])

	// Test on replica 1
	leaders = suite.dist.GetShardLeadersByReplica(replicas[1])
	suite.Len(leaders, 2)
	suite.Equal(leaders["dmc0"], suite.nodes[1])
	suite.Equal(leaders["dmc1"], suite.nodes[1])
}

func (suite *ChannelDistManagerSuite) AssertNames(channels []*DmChannel, names ...string) bool {
	for _, channel := range channels {
		hasChannel := false
		for _, name := range names {
			if channel.ChannelName == name {
				hasChannel = true
				break
			}
		}
		if !suite.True(hasChannel, "channel %v not in the given expected list %+v", channel.ChannelName, names) {
			return false
		}
	}
	return true
}

func (suite *ChannelDistManagerSuite) AssertNode(channels []*DmChannel, node int64) bool {
	for _, channel := range channels {
		if !suite.Equal(node, channel.Node) {
			return false
		}
	}
	return true
}

func (suite *ChannelDistManagerSuite) AssertCollection(channels []*DmChannel, collection int64) bool {
	for _, channel := range channels {
		if !suite.Equal(collection, channel.GetCollectionID()) {
			return false
		}
	}
	return true
}

func TestChannelDistManager(t *testing.T) {
	suite.Run(t, new(ChannelDistManagerSuite))
}
