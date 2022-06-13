package keeper_test

import (
	"time"

	"github.com/bitsongofficial/go-bitsong/x/marketplace/types"
	nfttypes "github.com/bitsongofficial/go-bitsong/x/nft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func (suite *KeeperTestSuite) TestLastAuctionIdGetSet() {
	// get default last auction id
	lastAuctionId := suite.app.MarketplaceKeeper.GetLastAuctionId(suite.ctx)
	suite.Require().Equal(lastAuctionId, uint64(0))

	// set last auction id to new value
	newAuctionId := uint64(2)
	suite.app.MarketplaceKeeper.SetLastAuctionId(suite.ctx, newAuctionId)

	// check last auction id update
	lastAuctionId = suite.app.MarketplaceKeeper.GetLastAuctionId(suite.ctx)
	suite.Require().Equal(lastAuctionId, newAuctionId)
}

func (suite *KeeperTestSuite) TestAuctionGetSet() {
	// get auction by not available id
	_, err := suite.app.MarketplaceKeeper.GetAuctionById(suite.ctx, 0)
	suite.Require().Error(err)

	// get all auctions when not available
	allAuctions := suite.app.MarketplaceKeeper.GetAllAuctions(suite.ctx)
	suite.Require().Len(allAuctions, 0)

	now := time.Now().UTC()
	suite.ctx = suite.ctx.WithBlockTime(now.Add(time.Second))

	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	owner2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	auctions := []types.Auction{
		{ // created auction
			Id:               1,
			Authority:        owner.String(),
			NftId:            1,
			PrizeType:        types.AuctionPrizeType_NftOnlyTransfer,
			Duration:         time.Second,
			BidDenom:         "ubtsg",
			PriceFloor:       1000000,
			InstantSalePrice: 2000000,
			TickSize:         10000,
			State:            types.AuctionState_Created,
			LastBidAmount:    0,
			LastBid:          time.Time{},
			EndedAt:          time.Time{},
			EndAuctionAt:     time.Time{},
			Claimed:          0,
		},
		{ // started auction
			Id:               2,
			Authority:        owner.String(),
			NftId:            2,
			PrizeType:        types.AuctionPrizeType_NftOnlyTransfer,
			Duration:         time.Second,
			BidDenom:         "ubtsg",
			PriceFloor:       1000000,
			InstantSalePrice: 2000000,
			TickSize:         10000,
			State:            types.AuctionState_Started,
			LastBidAmount:    0,
			LastBid:          time.Time{},
			EndedAt:          time.Time{},
			EndAuctionAt:     now.Add(time.Second),
			Claimed:          0,
		},
		{ // bid auction
			Id:               3,
			Authority:        owner.String(),
			NftId:            3,
			PrizeType:        types.AuctionPrizeType_NftOnlyTransfer,
			Duration:         time.Second,
			BidDenom:         "ubtsg",
			PriceFloor:       1000000,
			InstantSalePrice: 2000000,
			TickSize:         10000,
			State:            types.AuctionState_Started,
			LastBidAmount:    1000000,
			LastBid:          now,
			EndedAt:          time.Time{},
			EndAuctionAt:     now.Add(time.Second),
			Claimed:          0,
		},
		{ // ended auction
			Id:               4,
			Authority:        owner2.String(),
			NftId:            4,
			PrizeType:        types.AuctionPrizeType_NftOnlyTransfer,
			Duration:         time.Second,
			BidDenom:         "ubtsg",
			PriceFloor:       1000000,
			InstantSalePrice: 2000000,
			TickSize:         10000,
			State:            types.AuctionState_Ended,
			LastBidAmount:    1000000,
			LastBid:          now,
			EndedAt:          now.Add(time.Second * 2),
			EndAuctionAt:     now.Add(time.Second),
			Claimed:          0,
		},
		{ // claimed auction
			Id:               5,
			Authority:        owner2.String(),
			NftId:            5,
			PrizeType:        types.AuctionPrizeType_NftOnlyTransfer,
			Duration:         time.Second,
			BidDenom:         "ubtsg",
			PriceFloor:       1000000,
			InstantSalePrice: 2000000,
			TickSize:         10000,
			State:            types.AuctionState_Ended,
			LastBidAmount:    1000000,
			LastBid:          now,
			EndedAt:          now.Add(time.Second * 2),
			EndAuctionAt:     now.Add(time.Second),
			Claimed:          1,
		},
	}

	for _, auction := range auctions {
		suite.app.MarketplaceKeeper.SetAuction(suite.ctx, auction)
	}

	for _, auction := range auctions {
		c, err := suite.app.MarketplaceKeeper.GetAuctionById(suite.ctx, auction.Id)
		suite.Require().NoError(err)
		suite.Require().Equal(auction, c)
	}

	allAuctions = suite.app.MarketplaceKeeper.GetAllAuctions(suite.ctx)
	suite.Require().Len(allAuctions, 5)
	suite.Require().Equal(auctions, allAuctions)

	// test GetAuctionsByAuthority
	ownerAuctions := suite.app.MarketplaceKeeper.GetAuctionsByAuthority(suite.ctx, owner)
	suite.Require().Len(ownerAuctions, 3)

	// test DeleteAuction
	suite.app.MarketplaceKeeper.DeleteAuction(suite.ctx, auctions[0])
	allAuctions = suite.app.MarketplaceKeeper.GetAllAuctions(suite.ctx)
	suite.Require().Len(allAuctions, 4)
	ownerAuctions = suite.app.MarketplaceKeeper.GetAuctionsByAuthority(suite.ctx, owner)
	suite.Require().Len(ownerAuctions, 2)

	// auctions to end
	toEndAuctions := suite.app.MarketplaceKeeper.GetAuctionsByAuthority(suite.ctx, owner)
	suite.Require().Len(toEndAuctions, 2)
}

func (suite *KeeperTestSuite) TestCreateAuction() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	user2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	tests := []struct {
		testCase      string
		fee           sdk.Coin
		balance       sdk.Coin
		nftOwner      sdk.AccAddress
		metadataOwner sdk.AccAddress
		auctionType   types.AuctionPrizeType
		nftId         uint64
		expectPass    bool
	}{
		{
			"Not existing nft auction",
			sdk.NewInt64Coin("ubtsg", 0),
			sdk.NewInt64Coin("ubtsg", 0),
			owner,
			owner,
			types.AuctionPrizeType_NftOnlyTransfer,
			0,
			false,
		},
		{
			"Not owned nft auction",
			sdk.NewInt64Coin("ubtsg", 0),
			sdk.NewInt64Coin("ubtsg", 0),
			user2,
			owner,
			types.AuctionPrizeType_NftOnlyTransfer,
			1,
			false,
		},
		{
			"Not owned metadata auction",
			sdk.NewInt64Coin("ubtsg", 0),
			sdk.NewInt64Coin("ubtsg", 0),
			owner,
			user2,
			types.AuctionPrizeType_FullRightsTransfer,
			1,
			false,
		},
		{
			"Not enough balance for auction creation",
			sdk.NewInt64Coin("ubtsg", 2000),
			sdk.NewInt64Coin("ubtsg", 1000),
			owner,
			user2,
			types.AuctionPrizeType_NftOnlyTransfer,
			1,
			false,
		},
		{
			"Successful full rights transfer auction",
			sdk.NewInt64Coin("ubtsg", 0),
			sdk.NewInt64Coin("ubtsg", 0),
			owner,
			owner,
			types.AuctionPrizeType_FullRightsTransfer,
			1,
			true,
		},
		{
			"Successful nft only transfer auction",
			sdk.NewInt64Coin("ubtsg", 0),
			sdk.NewInt64Coin("ubtsg", 0),
			owner,
			user2,
			types.AuctionPrizeType_NftOnlyTransfer,
			1,
			true,
		},
		{
			"Successful fee payment auction",
			sdk.NewInt64Coin("ubtsg", 2000),
			sdk.NewInt64Coin("ubtsg", 2000),
			owner,
			user2,
			types.AuctionPrizeType_NftOnlyTransfer,
			1,
			true,
		},
	}

	for _, tc := range tests {

		// set nft with ownership
		nft := nfttypes.NFT{
			Id:         1,
			Owner:      tc.nftOwner.String(),
			MetadataId: 1,
		}
		suite.app.NFTKeeper.SetNFT(suite.ctx, nft)

		// set metadata with ownership
		metadata := nfttypes.Metadata{
			Id:              1,
			UpdateAuthority: tc.metadataOwner.String(),
		}
		suite.app.NFTKeeper.SetMetadata(suite.ctx, metadata)

		// mint coins if balance should set
		if tc.balance.IsPositive() {
			suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{tc.balance})
			suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, owner, sdk.Coins{tc.balance})
		}
		// set params
		suite.app.MarketplaceKeeper.SetParamSet(suite.ctx, types.Params{
			AuctionCreationPrice: tc.fee,
		})

		// get old balance for future check
		oldBalance := suite.app.BankKeeper.GetBalance(suite.ctx, owner, "ubtsg")

		msg := types.NewMsgCreateAuction(owner, tc.nftId, tc.auctionType, "ubtsg", time.Hour, 1, 1000, 1)
		// execute CreateAuction
		auctionId, err := suite.app.MarketplaceKeeper.CreateAuction(suite.ctx, msg)

		// check error exists on the execution
		if tc.expectPass {
			suite.Require().NoError(err)

			// check balance change
			newBalance := suite.app.BankKeeper.GetBalance(suite.ctx, owner, "ubtsg")
			suite.Require().Equal(newBalance.Amount.Int64()+tc.fee.Amount.Int64(), oldBalance.Amount.Int64())

			// module account
			moduleAddr := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)

			// check nft ownership transfer
			updatedNft, err := suite.app.NFTKeeper.GetNFTById(suite.ctx, msg.NftId)
			suite.Require().NoError(err)
			suite.Require().Equal(updatedNft.Owner, moduleAddr.String())

			// check metadata ownership transfer if full rights transfer auction
			if tc.auctionType == types.AuctionPrizeType_FullRightsTransfer {
				updatedMetadata, err := suite.app.NFTKeeper.GetMetadataById(suite.ctx, nft.MetadataId)
				suite.Require().NoError(err)
				suite.Require().Equal(updatedMetadata.UpdateAuthority, moduleAddr.String())
			}

			// check auction object created
			auction, err := suite.app.MarketplaceKeeper.GetAuctionById(suite.ctx, auctionId)
			suite.Require().NoError(err)
			suite.Require().Equal(auction, types.Auction{
				Id:               auctionId,
				Authority:        msg.Sender,
				NftId:            msg.NftId,
				PrizeType:        msg.PrizeType,
				Duration:         msg.Duration,
				BidDenom:         msg.BidDenom,
				PriceFloor:       msg.PriceFloor,
				InstantSalePrice: msg.InstantSalePrice,
				TickSize:         msg.TickSize,
				State:            types.AuctionState_Created,
				LastBidAmount:    0,
				LastBid:          time.Time{},
				EndedAt:          time.Time{},
				EndAuctionAt:     time.Time{},
				Claimed:          0,
			})
		} else {
			suite.Require().Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestStartAuction() {
	suite.ctx = suite.ctx.WithBlockTime(time.Now().UTC())
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	user2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	tests := []struct {
		testCase     string
		auctionOwner sdk.AccAddress
		auctionState types.AuctionState
		auctionId    uint64
		expectPass   bool
	}{
		{
			"Not existing auction id",
			owner,
			types.AuctionState_Created,
			0,
			false,
		},
		{
			"not auction authority",
			user2,
			types.AuctionState_Created,
			1,
			false,
		},
		{
			"auction already started",
			owner,
			types.AuctionState_Started,
			1,
			false,
		},
		{
			"Successful auction start",
			owner,
			types.AuctionState_Created,
			1,
			true,
		},
	}

	for _, tc := range tests {

		// set auction with ownership
		auction := types.Auction{
			Id:        1,
			Authority: tc.auctionOwner.String(),
			NftId:     1,
			Duration:  time.Second,
			State:     tc.auctionState,
		}
		suite.app.MarketplaceKeeper.SetAuction(suite.ctx, auction)

		// execute StartAuction
		msg := types.NewMsgStartAuction(owner, tc.auctionId)
		err := suite.app.MarketplaceKeeper.StartAuction(suite.ctx, msg)

		// check error exists on the execution
		if tc.expectPass {
			suite.Require().NoError(err)

			// check auction object updated
			auction, err := suite.app.MarketplaceKeeper.GetAuctionById(suite.ctx, tc.auctionId)
			suite.Require().NoError(err)
			suite.Require().Equal(auction.EndAuctionAt, suite.ctx.BlockTime().Add(auction.Duration))
			suite.Require().Equal(auction.State, types.AuctionState_Started)
		} else {
			suite.Require().Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestSetAuctionAuthority() {
	suite.ctx = suite.ctx.WithBlockTime(time.Now().UTC())
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	user2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	tests := []struct {
		testCase     string
		auctionOwner sdk.AccAddress
		auctionId    uint64
		expectPass   bool
	}{
		{
			"Not existing auction id",
			owner,
			0,
			false,
		},
		{
			"not auction authority",
			user2,
			1,
			false,
		},
		{
			"Successful auction authority update",
			owner,
			1,
			true,
		},
	}

	for _, tc := range tests {

		// set auction with ownership
		auction := types.Auction{
			Id:        1,
			Authority: tc.auctionOwner.String(),
			NftId:     1,
			Duration:  time.Second,
			State:     types.AuctionState_Created,
		}
		suite.app.MarketplaceKeeper.SetAuction(suite.ctx, auction)

		// execute SetAuctionAuthority
		msg := types.NewMsgSetAuctionAuthority(owner, tc.auctionId, user2.String())
		err := suite.app.MarketplaceKeeper.SetAuctionAuthority(suite.ctx, msg)

		// check error exists on the execution
		if tc.expectPass {
			suite.Require().NoError(err)

			// check auction authority updated
			auction, err := suite.app.MarketplaceKeeper.GetAuctionById(suite.ctx, tc.auctionId)
			suite.Require().NoError(err)
			suite.Require().Equal(auction.Authority, msg.NewAuthority)
		} else {
			suite.Require().Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestEndAuction() {
	suite.ctx = suite.ctx.WithBlockTime(time.Now().UTC())
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	owner2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	tests := []struct {
		testCase      string
		auctionOwner  sdk.AccAddress
		auctionType   types.AuctionPrizeType
		state         types.AuctionState
		lastBidAmount uint64
		auctionId     uint64
		expectPass    bool
	}{
		{
			"Not existing auction id",
			owner,
			types.AuctionPrizeType_NftOnlyTransfer,
			types.AuctionState_Started,
			0,
			0,
			false,
		},
		{
			"not auction authority",
			owner2,
			types.AuctionPrizeType_NftOnlyTransfer,
			types.AuctionState_Started,
			0,
			1,
			false,
		},
		{
			"already ended auction",
			owner,
			types.AuctionPrizeType_NftOnlyTransfer,
			types.AuctionState_Ended,
			0,
			1,
			false,
		},
		{
			"successful end with winning bid",
			owner,
			types.AuctionPrizeType_NftOnlyTransfer,
			types.AuctionState_Started,
			100,
			1,
			true,
		},
		{
			"return nft back if no bid when nft only transfer",
			owner,
			types.AuctionPrizeType_NftOnlyTransfer,
			types.AuctionState_Started,
			0,
			1,
			true,
		},
		{
			"return nft back if no bid when full rights transfer",
			owner,
			types.AuctionPrizeType_FullRightsTransfer,
			types.AuctionState_Started,
			0,
			1,
			true,
		},
	}

	for _, tc := range tests {
		// module account
		moduleAddr := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)

		// set nft with ownership
		nft := nfttypes.NFT{
			Id:         1,
			Owner:      moduleAddr.String(),
			MetadataId: 1,
		}
		suite.app.NFTKeeper.SetNFT(suite.ctx, nft)

		// set metadata with ownership
		metadata := nfttypes.Metadata{
			Id:              1,
			UpdateAuthority: moduleAddr.String(),
		}
		suite.app.NFTKeeper.SetMetadata(suite.ctx, metadata)

		// set auction with ownership
		auction := types.Auction{
			Id:            1,
			Authority:     tc.auctionOwner.String(),
			NftId:         1,
			Duration:      time.Second,
			PrizeType:     tc.auctionType,
			State:         tc.state,
			LastBidAmount: tc.lastBidAmount,
		}
		suite.app.MarketplaceKeeper.SetAuction(suite.ctx, auction)

		// execute SetAuctionAuthority
		msg := types.NewMsgEndAuction(owner, tc.auctionId)
		err := suite.app.MarketplaceKeeper.EndAuction(suite.ctx, msg)

		// check error exists on the execution
		if tc.expectPass {
			suite.Require().NoError(err)

			// check auction authority updated
			auction, err := suite.app.MarketplaceKeeper.GetAuctionById(suite.ctx, tc.auctionId)
			suite.Require().NoError(err)
			suite.Require().Equal(auction.EndedAt, suite.ctx.BlockTime())
			suite.Require().Equal(auction.State, types.AuctionState_Ended)

			nft, err := suite.app.NFTKeeper.GetNFTById(suite.ctx, nft.Id)
			suite.Require().NoError(err)
			if tc.lastBidAmount == 0 {
				suite.Require().Equal(nft.Owner, auction.Authority)
			} else {
				suite.Require().Equal(nft.Owner, moduleAddr.String())
			}

			if tc.auctionType == types.AuctionPrizeType_FullRightsTransfer {
				metadata, err := suite.app.NFTKeeper.GetMetadataById(suite.ctx, nft.MetadataId)
				suite.Require().NoError(err)
				if tc.lastBidAmount == 0 {
					suite.Require().Equal(metadata.UpdateAuthority, auction.Authority)
				} else {
					suite.Require().Equal(metadata.UpdateAuthority, moduleAddr.String())
				}
			}
		} else {
			suite.Require().Error(err)
		}
	}
}
