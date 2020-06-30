package yttracker

import (
	pb "github.com/yottachain/yotta-miner-tracker/pbtracker"
)

const (
	//Function tag
	Function = "function"
	//MinerID tag
	MinerID = "minerID"
	//AccountName tag
	AccountName = "account"
)

//Auth crenditial info
type Auth struct {
	//Account name in BP
	Account string `bson:"_id"`
	//PublicKey public key of account
	PublicKey string `bson:"publickey"`
}

// Node instance
type Node struct {
	//data node index
	ID int32 `bson:"_id"`
	//data node ID, generated from PubKey
	NodeID string `bson:"nodeid"`
	//public key of data node
	PubKey string `bson:"pubkey"`
	//owner account of this miner
	Owner string `bson:"owner"`
	//profit account of this miner
	ProfitAcc string `bson:"profitAcc"`
	//ID of associated miner pool
	PoolID string `bson:"poolID"`
	//Owner of associated miner pool
	PoolOwner string `bson:"poolOwner"`
	//quota allocated by associated miner pool
	Quota int64 `bson:"quota"`
	//listening addresses of data node
	Addrs []string `bson:"addrs"`
	//CPU usage of data node
	CPU int32 `bson:"cpu"`
	//memory usage of data node
	Memory int32 `bson:"memory"`
	//bandwidth usage of data node
	Bandwidth int32 `bson:"bandwidth"`
	//max space of data node
	MaxDataSpace int64 `bson:"maxDataSpace"`
	//space assigned to YTFS
	AssignedSpace int64 `bson:"assignedSpace"`
	//pre-allocated space of data node
	ProductiveSpace int64 `bson:"productiveSpace"`
	//used space of data node
	UsedSpace int64 `bson:"usedSpace"`
	//used spaces on each SN
	Uspaces map[string]int64 `bson:"uspaces"`
	//weight for allocate data node
	Weight float64 `bson:"weight"`
	//Is node valid
	Valid int32 `bson:"valid"`
	//Is relay node
	Relay int32 `bson:"relay"`
	//status code: 0 - registered 1 - active
	Status int32 `bson:"status"`
	//timestamp of status updating operation
	Timestamp int64 `bson:"timestamp"`
	//version number of miner
	Version int32 `bson:"version"`
	//Rebuilding if node is under rebuilding
	Rebuilding int32 `bson:"rebuilding"`
	//RealSpace real space of miner
	RealSpace int64 `bson:"realSpace"`
	//Tx
	Tx int64 `bson:"tx"`
	//Rx
	Rx int64 `bson:"rx"`
	//Ext
	Ext string `bson:"-"`
	//ManualWeight
	ManualWeight int32 `bson:"manualWeight"`
}

//NewNode create a node struct
func NewNode(id int32, nodeid string, pubkey string, owner string, profitAcc string, poolID string, poolOwner string, quota int64, addrs []string, cpu int32, memory int32, bandwidth int32, maxDataSpace int64, assignedSpace int64, productiveSpace int64, usedSpace int64, weight float64, valid int32, relay int32, status int32, timestamp int64, version int32, rebuilding int32, realSpace int64, tx int64, rx int64, ext string, manualWeight int32) *Node {
	return &Node{ID: id, NodeID: nodeid, PubKey: pubkey, Owner: owner, ProfitAcc: profitAcc, PoolID: poolID, PoolOwner: poolOwner, Quota: quota, Addrs: addrs, CPU: cpu, Memory: memory, Bandwidth: bandwidth, MaxDataSpace: maxDataSpace, AssignedSpace: assignedSpace, ProductiveSpace: productiveSpace, UsedSpace: usedSpace, Weight: weight, Valid: valid, Relay: relay, Status: status, Timestamp: timestamp, Version: version, Rebuilding: rebuilding, RealSpace: realSpace, Tx: tx, Rx: rx, Ext: ext, ManualWeight: manualWeight}
}

//relative DB and collection name
var (
	MinerTrackerDB = "minertracker"
	NodeTab        = "Node"
	AuthTab        = "Auth"
)

// Convert convert Node strcut to NodeMsg
func (node *Node) Convert() *pb.NodeMsg {
	return &pb.NodeMsg{
		ID:              node.ID,
		NodeID:          node.NodeID,
		PubKey:          node.PubKey,
		Owner:           node.Owner,
		ProfitAcc:       node.ProfitAcc,
		PoolID:          node.PoolID,
		PoolOwner:       node.PoolOwner,
		Quota:           node.Quota,
		Addrs:           node.Addrs,
		CPU:             node.CPU,
		Memory:          node.Memory,
		Bandwidth:       node.Bandwidth,
		MaxDataSpace:    node.MaxDataSpace,
		AssignedSpace:   node.AssignedSpace,
		ProductiveSpace: node.ProductiveSpace,
		UsedSpace:       node.UsedSpace,
		Uspaces:         node.Uspaces,
		Weight:          node.Weight,
		Valid:           node.Valid,
		Relay:           node.Relay,
		Status:          node.Status,
		Timestamp:       node.Timestamp,
		Version:         node.Version,
		Rebuilding:      node.Rebuilding,
		RealSpace:       node.RealSpace,
		Tx:              node.Tx,
		Rx:              node.Rx,
		Ext:             node.Ext,
		ManualWeight:    node.ManualWeight,
	}
}

// Fillby convert NodeMsg to Node struct
func (node *Node) Fillby(msg *pb.NodeMsg) {
	node.ID = msg.ID
	node.NodeID = msg.NodeID
	node.PubKey = msg.PubKey
	node.Owner = msg.Owner
	node.ProfitAcc = msg.ProfitAcc
	node.PoolID = msg.PoolID
	node.PoolOwner = msg.PoolOwner
	node.Quota = msg.Quota
	node.Addrs = msg.Addrs
	node.CPU = msg.CPU
	node.Memory = msg.Memory
	node.Bandwidth = msg.Bandwidth
	node.MaxDataSpace = msg.MaxDataSpace
	node.AssignedSpace = msg.AssignedSpace
	node.ProductiveSpace = msg.ProductiveSpace
	node.UsedSpace = msg.UsedSpace
	node.Uspaces = msg.Uspaces
	node.Weight = msg.Weight
	node.Valid = msg.Valid
	node.Relay = msg.Relay
	node.Status = msg.Status
	node.Timestamp = msg.Timestamp
	node.Version = msg.Version
	node.Rebuilding = msg.Rebuilding
	node.RealSpace = msg.RealSpace
	node.Tx = msg.Tx
	node.Rx = msg.Rx
	node.Ext = msg.Ext
	node.ManualWeight = msg.ManualWeight
}