package op

import (
	"github.com/coschain/contentos-go/common/constants"
	. "github.com/coschain/contentos-go/dandelion"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

type ConvertVestingTester struct {
	acc0, acc1, acc2, acc3, acc4 *DandelionAccount
}

func (tester *ConvertVestingTester) Test(t *testing.T, d *Dandelion) {
	tester.acc0 = d.Account("actor0")
	tester.acc1 = d.Account("actor1")
	tester.acc2 = d.Account("actor2")
	tester.acc3 = d.Account("actor3")
	tester.acc4 = d.Account("actor4")

	//a := assert.New(t)
	//a.NoError(tester.acc2.SendTrxAndProduceBlock(BpRegister(tester.acc2.Name, "", "", tester.acc2.GetPubKey(), defaultProps)))

	t.Run("normal", d.Test(tester.normal))
	t.Run("reset", d.Test(tester.Reset))
	t.Run("too much", d.Test(tester.tooMuch))
	t.Run("too small", d.Test(tester.tooSmall))
	t.Run("mismatch", d.Test(tester.mismatch))
}

func (tester *ConvertVestingTester) normal(t *testing.T, d *Dandelion) {
	a := assert.New(t)
	a.NoError(tester.acc2.SendTrxAndProduceBlock(BpRegister(tester.acc2.Name, "", "", tester.acc2.GetPubKey(), defaultProps)))

	a.NoError(tester.acc0.SendTrxAndProduceBlock(TransferToVesting(tester.acc0.Name, tester.acc0.Name, 10000000)))

	vestingShares0 := d.Account(tester.acc0.Name).GetVestingShares().Value
	balance0 := d.Account(tester.acc0.Name).GetBalance().Value

	headBlock0 := d.GlobalProps().GetHeadBlockNumber()
	a.NoError(tester.acc0.SendTrxAndProduceBlock(ConvertVesting(tester.acc0.Name, 10000000)))
	a.Equal(headBlock0 + uint64(constants.PowerDownBlockInterval), d.Account(tester.acc0.Name).GetNextPowerdownBlockNum())
	eachRate := 10000000 / (constants.ConvertWeeks - 1)
	a.Equal(uint64(eachRate), d.Account(tester.acc0.Name).GetEachPowerdownRate().Value)
	a.NoError(d.ProduceBlocks(constants.PowerDownBlockInterval + 1))
	a.Equal(balance0 + uint64(eachRate), d.Account(tester.acc0.Name).GetBalance().Value)
	a.NoError(d.ProduceBlocks(constants.PowerDownBlockInterval + 1))
	a.Equal(balance0 + uint64(eachRate) * 2, d.Account(tester.acc0.Name).GetBalance().Value)
	a.NoError(d.ProduceBlocks(2000))
	a.Equal(balance0 + 10000000, d.Account(tester.acc0.Name).GetBalance().Value)
	a.Equal(vestingShares0 - 10000000, d.Account(tester.acc0.Name).GetVestingShares().Value)
	a.Equal(uint64(0), d.Account(tester.acc0.Name).GetEachPowerdownRate().Value)
	a.Equal(uint64(math.MaxUint32), d.Account(tester.acc0.Name).GetNextPowerdownBlockNum())
}

func (tester *ConvertVestingTester) Reset(t *testing.T, d *Dandelion) {
	a := assert.New(t)
	a.NoError(tester.acc1.SendTrxAndProduceBlock(TransferToVesting(tester.acc1.Name, tester.acc1.Name, 10000000)))

	vestingShares1 := d.Account(tester.acc1.Name).GetVestingShares().Value
	balance1 := d.Account(tester.acc1.Name).GetBalance().Value
	headBlock1 := d.GlobalProps().GetHeadBlockNumber()

	a.NoError(tester.acc1.SendTrxAndProduceBlock(ConvertVesting(tester.acc1.Name, 10000000)))
	a.Equal(headBlock1 + uint64(constants.PowerDownBlockInterval), d.Account(tester.acc1.Name).GetNextPowerdownBlockNum())
	eachRate1 := 10000000 / (constants.ConvertWeeks - 1)
	a.Equal(uint64(eachRate1), d.Account(tester.acc1.Name).GetEachPowerdownRate().Value)
	a.NoError(d.ProduceBlocks(constants.PowerDownBlockInterval + 1))
	a.Equal(balance1 + uint64(eachRate1), d.Account(tester.acc1.Name).GetBalance().Value)
	a.Equal(vestingShares1 - uint64(eachRate1), d.Account(tester.acc1.Name).GetVestingShares().Value)

	vestingShares2 := d.Account(tester.acc1.Name).GetVestingShares().Value
	balance2 := d.Account(tester.acc1.Name).GetBalance().Value
	headBlock2 := d.GlobalProps().GetHeadBlockNumber()

	a.NoError(tester.acc1.SendTrxAndProduceBlock(ConvertVesting(tester.acc1.Name, 2000000)))
	a.Equal(headBlock2 + uint64(constants.PowerDownBlockInterval), d.Account(tester.acc1.Name).GetNextPowerdownBlockNum())
	eachRate2 := 2000000 / (constants.ConvertWeeks - 1)
	a.Equal(uint64(eachRate2), d.Account(tester.acc1.Name).GetEachPowerdownRate().Value)
	a.NoError(d.ProduceBlocks(constants.PowerDownBlockInterval + 1))
	a.Equal(balance2 + uint64(eachRate2), d.Account(tester.acc1.Name).GetBalance().Value)
	a.Equal(vestingShares2 - uint64(eachRate2), d.Account(tester.acc1.Name).GetVestingShares().Value)
}

func (tester *ConvertVestingTester) tooMuch(t *testing.T, d *Dandelion) {
	a := assert.New(t)
	a.NoError(tester.acc3.SendTrxAndProduceBlock(TransferToVesting(tester.acc3.Name, tester.acc3.Name, 10000000)))
	vestingShares := d.Account(tester.acc3.Name).GetVestingShares().Value
	toConvert := d.Account(tester.acc3.Name).GetToPowerdown().Value

	a.NoError(tester.acc3.SendTrx(ConvertVesting(tester.acc3.Name, vestingShares + 1)))
	a.NoError(d.ProduceBlocks(1))
	a.NoError(tester.acc3.SendTrx(ConvertVesting(tester.acc3.Name, math.MaxUint64)))
	a.NoError(d.ProduceBlocks(1))
	a.Equal(vestingShares, d.Account(tester.acc3.Name).GetVestingShares().Value)
	a.Equal(toConvert, d.Account(tester.acc3.Name).GetToPowerdown().Value)
}

func (tester *ConvertVestingTester) tooSmall(t *testing.T, d *Dandelion) {
	a := assert.New(t)
	a.NoError(tester.acc4.SendTrxAndProduceBlock(TransferToVesting(tester.acc4.Name, tester.acc4.Name, 10000000)))
	vestingShares := d.Account(tester.acc4.Name).GetVestingShares().Value
	toConvert := d.Account(tester.acc4.Name).GetToPowerdown().Value
	a.Error(tester.acc4.SendTrx(ConvertVesting(tester.acc4.Name, 1)))
	a.NoError(d.ProduceBlocks(1))
	a.Equal(vestingShares, d.Account(tester.acc4.Name).GetVestingShares().Value)
	a.Equal(toConvert, d.Account(tester.acc4.Name).GetToPowerdown().Value)
}

func (tester *ConvertVestingTester) mismatch(t *testing.T, d *Dandelion) {
	a := assert.New(t)
	vestingShares := d.Account(tester.acc0.Name).GetVestingShares().Value
	toConvert := d.Account(tester.acc0.Name).GetToPowerdown().Value
	a.Error(tester.acc0.SendTrx(ConvertVesting(tester.acc1.Name, 1000000)))
	a.NoError(d.ProduceBlocks(1))
	a.Equal(vestingShares, d.Account(tester.acc0.Name).GetVestingShares().Value)
	a.Equal(toConvert, d.Account(tester.acc0.Name).GetToPowerdown().Value)
}