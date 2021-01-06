package mock

import (
	"crypto/rand"
	"encoding/hex"
	"math"
	"math/big"
	"time"

	"github.com/qlcchain/qlc-go-sdk/pkg/random"

	"github.com/qlcchain/go-lsobus/rpc/grpc/proto"

	"github.com/google/uuid"
	qlcSdk "github.com/qlcchain/qlc-go-sdk"
	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"

	pb "github.com/qlcchain/go-lsobus/rpc/grpc/proto"
)

func GetOrderInfoByInternalId(id string) (*qlcSdk.DoDSettleOrderInfo, error) {
	orderInfo, err := OrderInfo()
	if err != nil {
		return nil, err
	}
	return orderInfo, nil
}

func OrderInfo() (*qlcSdk.DoDSettleOrderInfo, error) {
	buyer, seller := DoDSettleUser()
	orderType, err := qlcSdk.ParseDoDSettleOrderType("create")
	if err != nil {
		return nil, err
	}
	orderState, err := qlcSdk.ParseDoDSettleOrderState("complete")
	if err != nil {
		return nil, err
	}
	contractState, err := qlcSdk.ParseDoDSettleContractState("confirmed")
	if err != nil {
		return nil, err
	}
	orderInfo := &qlcSdk.DoDSettleOrderInfo{
		Buyer:         buyer,
		Seller:        seller,
		OrderId:       uuid.New().String(),
		OrderType:     orderType,
		OrderState:    orderState,
		ContractState: contractState,
		Connections:   make([]*qlcSdk.DoDSettleConnectionParam, 0),
		Track:         make([]*qlcSdk.DoDSettleOrderLifeTrack, 0),
	}
	paymentType, err := qlcSdk.ParseDoDSettlePaymentType("invoice")
	if err != nil {
		return nil, err
	}
	billingType, err := qlcSdk.ParseDoDSettleBillingType("DOD")
	if err != nil {
		return nil, err
	}
	serviceClass, err := qlcSdk.ParseDoDSettleServiceClass("gold")
	if err != nil {
		return nil, err
	}
	conn := &qlcSdk.DoDSettleConnectionParam{
		DoDSettleConnectionStaticParam: qlcSdk.DoDSettleConnectionStaticParam{
			BuyerProductId: uuid.New().String(),
			SrcCompanyName: "CBC",
			SrcRegion:      "CHN",
			SrcCity:        "HK",
			SrcDataCenter:  "DCX",
			SrcPort:        "port01",
			DstCompanyName: "CBC",
			DstRegion:      "USA",
			DstCity:        "NYC",
			DstDataCenter:  "DCY",
			DstPort:        "port02",
		},
		DoDSettleConnectionDynamicParam: qlcSdk.DoDSettleConnectionDynamicParam{
			ItemId:         uuid.New().String(),
			ConnectionName: "conn",
			QuoteId:        "1",
			QuoteItemId:    "1",
			Bandwidth:      "100 Mbps",
			Price:          100.00,
			ServiceClass:   serviceClass,
			PaymentType:    paymentType,
			BillingType:    billingType,
			Currency:       "USD",
		},
	}
	orderInfo.Connections = append(orderInfo.Connections, conn)
	contractState1, err := qlcSdk.ParseDoDSettleContractState("request")
	if err != nil {
		return nil, err
	}
	orderState1, err := qlcSdk.ParseDoDSettleOrderState("null")
	if err != nil {
		return nil, err
	}
	track1 := &qlcSdk.DoDSettleOrderLifeTrack{
		ContractState: contractState1,
		OrderState:    orderState1,
		Time:          time.Now().Unix(),
		Hash:          RandomHash(),
	}
	contractState2, err := qlcSdk.ParseDoDSettleContractState("confirmed")
	if err != nil {
		return nil, err
	}
	orderState2, err := qlcSdk.ParseDoDSettleOrderState("null")
	if err != nil {
		return nil, err
	}
	track2 := &qlcSdk.DoDSettleOrderLifeTrack{
		ContractState: contractState2,
		OrderState:    orderState2,
		Time:          time.Now().Unix() + 1,
		Hash:          RandomHash(),
	}
	contractState3, err := qlcSdk.ParseDoDSettleContractState("confirmed")
	if err != nil {
		return nil, err
	}
	orderState3, err := qlcSdk.ParseDoDSettleOrderState("success")
	if err != nil {
		return nil, err
	}
	track3 := &qlcSdk.DoDSettleOrderLifeTrack{
		ContractState: contractState3,
		OrderState:    orderState3,
		Time:          time.Now().Unix() + 2,
		Hash:          RandomHash(),
	}
	contractState4, err := qlcSdk.ParseDoDSettleContractState("confirmed")
	if err != nil {
		return nil, err
	}
	orderState4, err := qlcSdk.ParseDoDSettleOrderState("complete")
	if err != nil {
		return nil, err
	}
	track4 := &qlcSdk.DoDSettleOrderLifeTrack{
		ContractState: contractState4,
		OrderState:    orderState4,
		Time:          time.Now().Unix() + 3,
		Hash:          RandomHash(),
	}
	orderInfo.Track = append(orderInfo.Track, track1, track2, track3, track4)
	return orderInfo, nil
}

func ProtoCreateOrderParams() *proto.CreateOrderParam {
	buyer, seller := DoDSettleUser()
	param := &pb.CreateOrderParam{
		Buyer: &pb.User{
			Address: buyer.Address.String(),
			Name:    buyer.Name,
		},
		Seller: &pb.User{
			Address: seller.Address.String(),
			Name:    seller.Name,
		},
		ConnectionParam: make([]*pb.ConnectionParam, 0),
	}

	conn1 := &pb.ConnectionParam{
		StaticParam: &pb.ConnectionStaticParam{
			BuyerProductId: uuid.New().String(),
			ProductId:      uuid.New().String(),
			SrcCompanyName: "CBC",
			SrcRegion:      "CHN",
			SrcCity:        "HK",
			SrcDataCenter:  "DCX",
			SrcPort:        "port01",
			DstCompanyName: "PCCWG",
			DstRegion:      "USA",
			DstCity:        "NYC",
			DstDataCenter:  "DCY",
			DstPort:        "port02",
		},
		DynamicParam: &pb.ConnectionDynamicParam{
			ItemId:         uuid.New().String(),
			ConnectionName: uuid.New().String(),
			OrderId:        uuid.New().String(),
			QuoteId:        uuid.New().String(),
			QuoteItemId:    uuid.New().String(),
			Bandwidth:      "100 Mbps",
			Price:          100.00,
			Addition:       5.00,
			ServiceClass:   "gold",
			PaymentType:    "invoice",
			BillingType:    "DOD",
			BillingUnit:    "month",
			Currency:       "USD",
			StartTime:      time.Now().Unix(),
			EndTime:        time.Now().Unix() + 100,
		},
	}
	param.ConnectionParam = append(param.ConnectionParam, conn1)
	conn2 := &pb.ConnectionParam{
		StaticParam: &pb.ConnectionStaticParam{
			BuyerProductId: uuid.New().String(),
			SrcCompanyName: "CBC",
			SrcRegion:      "CHN",
			SrcCity:        "HK",
			SrcDataCenter:  "DCX",
			SrcPort:        "port1",
			DstCompanyName: "CBC",
			DstRegion:      "USA",
			DstCity:        "NYC",
			DstDataCenter:  "DCY",
			DstPort:        "port2",
		},
		DynamicParam: &pb.ConnectionDynamicParam{
			ItemId:         uuid.New().String(),
			ConnectionName: uuid.New().String(),
			QuoteId:        uuid.New().String(),
			QuoteItemId:    uuid.New().String(),
			Bandwidth:      "600 Mbps",
			BillingUnit:    "month",
			Price:          200.00,
			ServiceClass:   "gold",
			PaymentType:    "invoice",
			BillingType:    "PAYG",
			Currency:       "USD",
		},
	}
	param.ConnectionParam = append(param.ConnectionParam, conn2)
	return param
}

func ProtoChangeOrderParams() *proto.ChangeOrderParam {
	buyer, seller := DoDSettleUser()
	param := &pb.ChangeOrderParam{
		Buyer: &pb.User{
			Address: buyer.Address.String(),
			Name:    buyer.Name,
		},
		Seller: &pb.User{
			Address: seller.Address.String(),
			Name:    seller.Name,
		},
		ChangeConnectionParam: make([]*pb.ChangeConnectionParam, 0),
	}
	conn1 := &pb.ChangeConnectionParam{
		ProductId: uuid.New().String(),
		DynamicParam: &pb.ConnectionDynamicParam{
			ConnectionName: uuid.New().String(),
			OrderId:        uuid.New().String(),
			QuoteId:        uuid.New().String(),
			QuoteItemId:    uuid.New().String(),
			Bandwidth:      "100 Mbps",
			Price:          100.00,
			Addition:       5.00,
			ServiceClass:   "gold",
			PaymentType:    "invoice",
			BillingType:    "DOD",
			BillingUnit:    "month",
			Currency:       "USD",
			StartTime:      time.Now().Unix(),
			EndTime:        time.Now().Unix() + 100,
		},
	}
	param.ChangeConnectionParam = append(param.ChangeConnectionParam, conn1)
	conn2 := &pb.ChangeConnectionParam{
		ProductId: uuid.New().String(),
		DynamicParam: &pb.ConnectionDynamicParam{
			ConnectionName: uuid.New().String(),
			QuoteId:        uuid.New().String(),
			QuoteItemId:    uuid.New().String(),
			Bandwidth:      "600 Mbps",
			BillingUnit:    "month",
			Price:          200.00,
			ServiceClass:   "gold",
			PaymentType:    "invoice",
			BillingType:    "PAYG",
			Currency:       "USD",
		},
	}
	param.ChangeConnectionParam = append(param.ChangeConnectionParam, conn2)
	return param
}

func ProtoTerminateOrderParams() *proto.TerminateOrderParam {
	buyer, seller := DoDSettleUser()
	param := &pb.TerminateOrderParam{
		Buyer: &pb.User{
			Address: buyer.Address.String(),
			Name:    buyer.Name,
		},
		Seller: &pb.User{
			Address: seller.Address.String(),
			Name:    seller.Name,
		},
		TerminateConnectionParam: make([]*pb.TerminateConnectionParam, 0),
	}
	conn1 := &pb.TerminateConnectionParam{
		ProductId: uuid.New().String(),
		DynamicParam: &pb.ConnectionDynamicParam{
			ConnectionName: uuid.New().String(),
			OrderId:        uuid.New().String(),
			QuoteId:        uuid.New().String(),
			QuoteItemId:    uuid.New().String(),
			Bandwidth:      "100 Mbps",
			Price:          100.00,
			Addition:       5.00,
			ServiceClass:   "gold",
			PaymentType:    "invoice",
			BillingType:    "DOD",
			BillingUnit:    "month",
			Currency:       "USD",
			StartTime:      time.Now().Unix(),
			EndTime:        time.Now().Unix() + 100,
		},
	}
	param.TerminateConnectionParam = append(param.TerminateConnectionParam, conn1)
	conn2 := &pb.TerminateConnectionParam{
		ProductId: uuid.New().String(),
		DynamicParam: &pb.ConnectionDynamicParam{
			ConnectionName: uuid.New().String(),
			QuoteId:        uuid.New().String(),
			QuoteItemId:    uuid.New().String(),
			Bandwidth:      "600 Mbps",
			BillingUnit:    "month",
			Price:          200.00,
			ServiceClass:   "gold",
			PaymentType:    "invoice",
			BillingType:    "PAYG",
			Currency:       "USD",
		},
	}
	param.TerminateConnectionParam = append(param.TerminateConnectionParam, conn2)
	return param
}

func DoDSettleUser() (*qlcSdk.DoDSettleUser, *qlcSdk.DoDSettleUser) {
	buyerAddr, _, _ := pkg.GenerateAddress()
	sellerAddr, _, _ := pkg.GenerateAddress()
	buyer := &qlcSdk.DoDSettleUser{
		Address: buyerAddr,
		Name:    "CBC",
	}
	seller := &qlcSdk.DoDSettleUser{
		Address: sellerAddr,
		Name:    "PCCWG",
	}
	return buyer, seller
}

func RandomHash() pkg.Hash {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	s := hex.EncodeToString(b)
	hash, _ := pkg.NewHash(s)
	return hash
}

func GetCreateOrderBlock(op *qlcSdk.DoDSettleCreateOrderParam, sign qlcSdk.Signature) (*pkg.StateBlock, error) {
	blk := StateBlockWithoutWork()
	var err error
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return blk, nil
}

func GetCreateOrderRewardBlock(op *qlcSdk.DoDSettleResponseParam, sign qlcSdk.Signature) (*pkg.StateBlock, error) {
	blk := StateBlockWithoutWork()
	var err error
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return blk, nil
}

func GetChangeOrderBlock(op *qlcSdk.DoDSettleChangeOrderParam, sign qlcSdk.Signature) (*pkg.StateBlock, error) {
	blk := StateBlockWithoutWork()
	var err error
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return blk, nil
}

func GetChangeOrderRewardBlock(op *qlcSdk.DoDSettleResponseParam, sign qlcSdk.Signature) (*pkg.StateBlock, error) {
	blk := StateBlockWithoutWork()
	var err error
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return blk, nil
}

func GetTerminateOrderBlock(op *qlcSdk.DoDSettleTerminateOrderParam, sign qlcSdk.Signature) (*pkg.StateBlock, error) {
	blk := StateBlockWithoutWork()
	var err error
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return blk, nil
}

func GetTerminateOrderRewardBlock(op *qlcSdk.DoDSettleResponseParam, sign qlcSdk.Signature) (*pkg.StateBlock, error) {
	blk := StateBlockWithoutWork()
	var err error
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return blk, nil
}

func GetUpdateOrderInfoBlock(op *qlcSdk.DoDSettleUpdateOrderInfoParam, sign qlcSdk.Signature) (*pkg.StateBlock, error) {
	blk := StateBlockWithoutWork()
	var err error
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return blk, nil
}

func GetUpdateProductInfoBlock(
	op *qlcSdk.DoDSettleUpdateProductInfoParam, sign qlcSdk.Signature,
) (*pkg.StateBlock, error) {
	blk := StateBlockWithoutWork()
	var err error
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return blk, nil
}

func GetPendingRequest(addr pkg.Address) []*qlcSdk.DoDPendingRequestRsp {
	pend := make([]*qlcSdk.DoDPendingRequestRsp, 0)
	order1, _ := OrderInfo()
	pend1 := &qlcSdk.DoDPendingRequestRsp{
		Hash:  RandomHash(),
		Order: order1,
	}
	order2, _ := OrderInfo()
	order2.OrderType, _ = qlcSdk.ParseDoDSettleOrderType("change")
	pend2 := &qlcSdk.DoDPendingRequestRsp{
		Hash:  RandomHash(),
		Order: order2,
	}
	order3, _ := OrderInfo()
	order3.OrderType, _ = qlcSdk.ParseDoDSettleOrderType("terminate")
	pend3 := &qlcSdk.DoDPendingRequestRsp{
		Hash:  RandomHash(),
		Order: order3,
	}
	pend = append(pend, pend1, pend2, pend3)
	return pend
}

func GetPendingResourceCheck(addr pkg.Address) []*qlcSdk.DoDPendingResourceCheckInfo {
	infos := make([]*qlcSdk.DoDPendingResourceCheckInfo, 0)
	info1 := &qlcSdk.DoDPendingResourceCheckInfo{
		SendHash:   RandomHash(),
		OrderId:    uuid.New().String(),
		InternalId: RandomHash(),
		Products:   make([]*qlcSdk.DoDSettleProductInfo, 0),
	}
	product1 := &qlcSdk.DoDSettleProductInfo{
		ProductId: uuid.New().String(),
		Active:    false,
	}
	info1.Products = append(info1.Products, product1)
	info2 := &qlcSdk.DoDPendingResourceCheckInfo{
		SendHash:   RandomHash(),
		OrderId:    uuid.New().String(),
		InternalId: RandomHash(),
		Products:   make([]*qlcSdk.DoDSettleProductInfo, 0),
	}
	product2 := &qlcSdk.DoDSettleProductInfo{
		ProductId: uuid.New().String(),
		Active:    false,
	}
	info2.Products = append(info1.Products, product2)
	infos = append(infos, info1, info2)
	return infos
}

func GetPendingResourceCheckForProductId(addr pkg.Address) []*qlcSdk.DoDPendingResourceCheckInfo {
	infos := make([]*qlcSdk.DoDPendingResourceCheckInfo, 0)
	info1 := &qlcSdk.DoDPendingResourceCheckInfo{
		SendHash:   RandomHash(),
		OrderId:    uuid.New().String(),
		InternalId: RandomHash(),
		Products:   make([]*qlcSdk.DoDSettleProductInfo, 0),
	}
	//product1 := &qlcSdk.DoDSettleProductInfo{
	//	ProductId: uuid.New().String(),
	//	Active:    false,
	//}
	//info1.Products = append(info1.Products, product1)
	info2 := &qlcSdk.DoDPendingResourceCheckInfo{
		SendHash:   RandomHash(),
		OrderId:    uuid.New().String(),
		InternalId: RandomHash(),
		Products:   make([]*qlcSdk.DoDSettleProductInfo, 0),
	}
	//product2 := &qlcSdk.DoDSettleProductInfo{
	//	ProductId: uuid.New().String(),
	//	Active:    false,
	//}
	//info2.Products = append(info1.Products, product2)
	infos = append(infos, info1, info2)
	return infos
}

func GetUpdateOrderInfoRewardBlock(
	param *qlcSdk.DoDSettleResponseParam, sign qlcSdk.Signature,
) (*pkg.StateBlock, error) {
	blk := StateBlockWithoutWork()
	var err error
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return blk, nil
}

func StateBlockWithoutWork() *pkg.StateBlock {
	sb := new(pkg.StateBlock)
	a, _, _ := pkg.GenerateAddress()
	i, _ := random.Intn(math.MaxInt16)
	sb.Type = pkg.ContractSend
	sb.Balance = pkg.Balance{Int: big.NewInt(int64(i))}
	sb.Vote = pkg.NewBalance(0)
	sb.Network = pkg.NewBalance(0)
	sb.Oracle = pkg.NewBalance(0)
	sb.Storage = pkg.NewBalance(0)
	sb.Address = a
	sb.Previous = RandomHash()
	sb.Representative = a
	sb.Timestamp = time.Now().Unix()
	sb.Link = RandomHash()
	sb.Message = RandomHash()
	return sb
}
