package x86

import (
	"cmd/internal/obj"
	"reflect"
	"testing"
)

func BenchmarkYmovtabMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, ok := ymovtab[1]
		if ok {
			ok = false
		}
	}
}

func BenchmarkYmovtabarry(b *testing.B) {
	for j := 0; j < b.N; j++ {
		for i := 0; i < len(ymovtabOrigin); i++ {
			if ymovtabOrigin[i].as == 1 {

			}
		}
	}
}

func TestSame(t *testing.T) {
	genM := make(map[obj.As][]movtab)
	for _, item := range ymovtabOrigin {
		genM[item.as] = append(genM[item.as], item)
	}
	newCount := 0
	for as, slice := range ymovtab {
		newCount += len(slice)
		oldSlice := genM[as]
		for i, newItem := range slice {
			oldItem := oldSlice[i]
			if !movtabEqual(oldItem, newItem) {
				t.Fatal("not equal", as)
			}
		}
	}
	if newCount != len(ymovtabOrigin) {
		t.Fatal("count not equal")
	}
}

func movtabEqual(a, b movtab) bool {
	return reflect.DeepEqual(a, b)
}

var ymovtabOrigin = []movtab{
	// push
	{APUSHL, Ycs, Ynone, Ynone, movLit, [4]uint8{0x0e, 0}},
	{APUSHL, Yss, Ynone, Ynone, movLit, [4]uint8{0x16, 0}},
	{APUSHL, Yds, Ynone, Ynone, movLit, [4]uint8{0x1e, 0}},
	{APUSHL, Yes, Ynone, Ynone, movLit, [4]uint8{0x06, 0}},
	{APUSHL, Yfs, Ynone, Ynone, movLit, [4]uint8{0x0f, 0xa0, 0}},
	{APUSHL, Ygs, Ynone, Ynone, movLit, [4]uint8{0x0f, 0xa8, 0}},
	{APUSHQ, Yfs, Ynone, Ynone, movLit, [4]uint8{0x0f, 0xa0, 0}},
	{APUSHQ, Ygs, Ynone, Ynone, movLit, [4]uint8{0x0f, 0xa8, 0}},
	{APUSHW, Ycs, Ynone, Ynone, movLit, [4]uint8{Pe, 0x0e, 0}},
	{APUSHW, Yss, Ynone, Ynone, movLit, [4]uint8{Pe, 0x16, 0}},
	{APUSHW, Yds, Ynone, Ynone, movLit, [4]uint8{Pe, 0x1e, 0}},
	{APUSHW, Yes, Ynone, Ynone, movLit, [4]uint8{Pe, 0x06, 0}},
	{APUSHW, Yfs, Ynone, Ynone, movLit, [4]uint8{Pe, 0x0f, 0xa0, 0}},
	{APUSHW, Ygs, Ynone, Ynone, movLit, [4]uint8{Pe, 0x0f, 0xa8, 0}},

	// pop
	{APOPL, Ynone, Ynone, Yds, movLit, [4]uint8{0x1f, 0}},
	{APOPL, Ynone, Ynone, Yes, movLit, [4]uint8{0x07, 0}},
	{APOPL, Ynone, Ynone, Yss, movLit, [4]uint8{0x17, 0}},
	{APOPL, Ynone, Ynone, Yfs, movLit, [4]uint8{0x0f, 0xa1, 0}},
	{APOPL, Ynone, Ynone, Ygs, movLit, [4]uint8{0x0f, 0xa9, 0}},
	{APOPQ, Ynone, Ynone, Yfs, movLit, [4]uint8{0x0f, 0xa1, 0}},
	{APOPQ, Ynone, Ynone, Ygs, movLit, [4]uint8{0x0f, 0xa9, 0}},
	{APOPW, Ynone, Ynone, Yds, movLit, [4]uint8{Pe, 0x1f, 0}},
	{APOPW, Ynone, Ynone, Yes, movLit, [4]uint8{Pe, 0x07, 0}},
	{APOPW, Ynone, Ynone, Yss, movLit, [4]uint8{Pe, 0x17, 0}},
	{APOPW, Ynone, Ynone, Yfs, movLit, [4]uint8{Pe, 0x0f, 0xa1, 0}},
	{APOPW, Ynone, Ynone, Ygs, movLit, [4]uint8{Pe, 0x0f, 0xa9, 0}},

	// mov seg
	{AMOVW, Yes, Ynone, Yml, movRegMem, [4]uint8{0x8c, 0, 0, 0}},
	{AMOVW, Ycs, Ynone, Yml, movRegMem, [4]uint8{0x8c, 1, 0, 0}},
	{AMOVW, Yss, Ynone, Yml, movRegMem, [4]uint8{0x8c, 2, 0, 0}},
	{AMOVW, Yds, Ynone, Yml, movRegMem, [4]uint8{0x8c, 3, 0, 0}},
	{AMOVW, Yfs, Ynone, Yml, movRegMem, [4]uint8{0x8c, 4, 0, 0}},
	{AMOVW, Ygs, Ynone, Yml, movRegMem, [4]uint8{0x8c, 5, 0, 0}},
	{AMOVW, Yml, Ynone, Yes, movMemReg, [4]uint8{0x8e, 0, 0, 0}},
	{AMOVW, Yml, Ynone, Ycs, movMemReg, [4]uint8{0x8e, 1, 0, 0}},
	{AMOVW, Yml, Ynone, Yss, movMemReg, [4]uint8{0x8e, 2, 0, 0}},
	{AMOVW, Yml, Ynone, Yds, movMemReg, [4]uint8{0x8e, 3, 0, 0}},
	{AMOVW, Yml, Ynone, Yfs, movMemReg, [4]uint8{0x8e, 4, 0, 0}},
	{AMOVW, Yml, Ynone, Ygs, movMemReg, [4]uint8{0x8e, 5, 0, 0}},

	// mov cr
	{AMOVL, Ycr0, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x20, 0, 0}},
	{AMOVL, Ycr2, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x20, 2, 0}},
	{AMOVL, Ycr3, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x20, 3, 0}},
	{AMOVL, Ycr4, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x20, 4, 0}},
	{AMOVL, Ycr8, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x20, 8, 0}},
	{AMOVQ, Ycr0, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x20, 0, 0}},
	{AMOVQ, Ycr2, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x20, 2, 0}},
	{AMOVQ, Ycr3, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x20, 3, 0}},
	{AMOVQ, Ycr4, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x20, 4, 0}},
	{AMOVQ, Ycr8, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x20, 8, 0}},
	{AMOVL, Yrl, Ynone, Ycr0, movMemReg2op, [4]uint8{0x0f, 0x22, 0, 0}},
	{AMOVL, Yrl, Ynone, Ycr2, movMemReg2op, [4]uint8{0x0f, 0x22, 2, 0}},
	{AMOVL, Yrl, Ynone, Ycr3, movMemReg2op, [4]uint8{0x0f, 0x22, 3, 0}},
	{AMOVL, Yrl, Ynone, Ycr4, movMemReg2op, [4]uint8{0x0f, 0x22, 4, 0}},
	{AMOVL, Yrl, Ynone, Ycr8, movMemReg2op, [4]uint8{0x0f, 0x22, 8, 0}},
	{AMOVQ, Yrl, Ynone, Ycr0, movMemReg2op, [4]uint8{0x0f, 0x22, 0, 0}},
	{AMOVQ, Yrl, Ynone, Ycr2, movMemReg2op, [4]uint8{0x0f, 0x22, 2, 0}},
	{AMOVQ, Yrl, Ynone, Ycr3, movMemReg2op, [4]uint8{0x0f, 0x22, 3, 0}},
	{AMOVQ, Yrl, Ynone, Ycr4, movMemReg2op, [4]uint8{0x0f, 0x22, 4, 0}},
	{AMOVQ, Yrl, Ynone, Ycr8, movMemReg2op, [4]uint8{0x0f, 0x22, 8, 0}},

	// mov dr
	{AMOVL, Ydr0, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x21, 0, 0}},
	{AMOVL, Ydr6, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x21, 6, 0}},
	{AMOVL, Ydr7, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x21, 7, 0}},
	{AMOVQ, Ydr0, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x21, 0, 0}},
	{AMOVQ, Ydr2, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x21, 2, 0}},
	{AMOVQ, Ydr3, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x21, 3, 0}},
	{AMOVQ, Ydr6, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x21, 6, 0}},
	{AMOVQ, Ydr7, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x21, 7, 0}},
	{AMOVL, Yrl, Ynone, Ydr0, movMemReg2op, [4]uint8{0x0f, 0x23, 0, 0}},
	{AMOVL, Yrl, Ynone, Ydr6, movMemReg2op, [4]uint8{0x0f, 0x23, 6, 0}},
	{AMOVL, Yrl, Ynone, Ydr7, movMemReg2op, [4]uint8{0x0f, 0x23, 7, 0}},
	{AMOVQ, Yrl, Ynone, Ydr0, movMemReg2op, [4]uint8{0x0f, 0x23, 0, 0}},
	{AMOVQ, Yrl, Ynone, Ydr2, movMemReg2op, [4]uint8{0x0f, 0x23, 2, 0}},
	{AMOVQ, Yrl, Ynone, Ydr3, movMemReg2op, [4]uint8{0x0f, 0x23, 3, 0}},
	{AMOVQ, Yrl, Ynone, Ydr6, movMemReg2op, [4]uint8{0x0f, 0x23, 6, 0}},
	{AMOVQ, Yrl, Ynone, Ydr7, movMemReg2op, [4]uint8{0x0f, 0x23, 7, 0}},

	// mov tr
	{AMOVL, Ytr6, Ynone, Yml, movRegMem2op, [4]uint8{0x0f, 0x24, 6, 0}},
	{AMOVL, Ytr7, Ynone, Yml, movRegMem2op, [4]uint8{0x0f, 0x24, 7, 0}},
	{AMOVL, Yml, Ynone, Ytr6, movMemReg2op, [4]uint8{0x0f, 0x26, 6, 0xff}},
	{AMOVL, Yml, Ynone, Ytr7, movMemReg2op, [4]uint8{0x0f, 0x26, 7, 0xff}},

	// lgdt, sgdt, lidt, sidt
	{AMOVL, Ym, Ynone, Ygdtr, movMemReg2op, [4]uint8{0x0f, 0x01, 2, 0}},
	{AMOVL, Ygdtr, Ynone, Ym, movRegMem2op, [4]uint8{0x0f, 0x01, 0, 0}},
	{AMOVL, Ym, Ynone, Yidtr, movMemReg2op, [4]uint8{0x0f, 0x01, 3, 0}},
	{AMOVL, Yidtr, Ynone, Ym, movRegMem2op, [4]uint8{0x0f, 0x01, 1, 0}},
	{AMOVQ, Ym, Ynone, Ygdtr, movMemReg2op, [4]uint8{0x0f, 0x01, 2, 0}},
	{AMOVQ, Ygdtr, Ynone, Ym, movRegMem2op, [4]uint8{0x0f, 0x01, 0, 0}},
	{AMOVQ, Ym, Ynone, Yidtr, movMemReg2op, [4]uint8{0x0f, 0x01, 3, 0}},
	{AMOVQ, Yidtr, Ynone, Ym, movRegMem2op, [4]uint8{0x0f, 0x01, 1, 0}},

	// lldt, sldt
	{AMOVW, Yml, Ynone, Yldtr, movMemReg2op, [4]uint8{0x0f, 0x00, 2, 0}},
	{AMOVW, Yldtr, Ynone, Yml, movRegMem2op, [4]uint8{0x0f, 0x00, 0, 0}},

	// lmsw, smsw
	{AMOVW, Yml, Ynone, Ymsw, movMemReg2op, [4]uint8{0x0f, 0x01, 6, 0}},
	{AMOVW, Ymsw, Ynone, Yml, movRegMem2op, [4]uint8{0x0f, 0x01, 4, 0}},

	// ltr, str
	{AMOVW, Yml, Ynone, Ytask, movMemReg2op, [4]uint8{0x0f, 0x00, 3, 0}},
	{AMOVW, Ytask, Ynone, Yml, movRegMem2op, [4]uint8{0x0f, 0x00, 1, 0}},

	/* load full pointer - unsupported
	{AMOVL, Yml, Ycol, movFullPtr, [4]uint8{0, 0, 0, 0}},
	{AMOVW, Yml, Ycol, movFullPtr, [4]uint8{Pe, 0, 0, 0}},
	*/

	// double shift
	{ASHLL, Yi8, Yrl, Yml, movDoubleShift, [4]uint8{0xa4, 0xa5, 0, 0}},
	{ASHLL, Ycl, Yrl, Yml, movDoubleShift, [4]uint8{0xa4, 0xa5, 0, 0}},
	{ASHLL, Ycx, Yrl, Yml, movDoubleShift, [4]uint8{0xa4, 0xa5, 0, 0}},
	{ASHRL, Yi8, Yrl, Yml, movDoubleShift, [4]uint8{0xac, 0xad, 0, 0}},
	{ASHRL, Ycl, Yrl, Yml, movDoubleShift, [4]uint8{0xac, 0xad, 0, 0}},
	{ASHRL, Ycx, Yrl, Yml, movDoubleShift, [4]uint8{0xac, 0xad, 0, 0}},
	{ASHLQ, Yi8, Yrl, Yml, movDoubleShift, [4]uint8{Pw, 0xa4, 0xa5, 0}},
	{ASHLQ, Ycl, Yrl, Yml, movDoubleShift, [4]uint8{Pw, 0xa4, 0xa5, 0}},
	{ASHLQ, Ycx, Yrl, Yml, movDoubleShift, [4]uint8{Pw, 0xa4, 0xa5, 0}},
	{ASHRQ, Yi8, Yrl, Yml, movDoubleShift, [4]uint8{Pw, 0xac, 0xad, 0}},
	{ASHRQ, Ycl, Yrl, Yml, movDoubleShift, [4]uint8{Pw, 0xac, 0xad, 0}},
	{ASHRQ, Ycx, Yrl, Yml, movDoubleShift, [4]uint8{Pw, 0xac, 0xad, 0}},
	{ASHLW, Yi8, Yrl, Yml, movDoubleShift, [4]uint8{Pe, 0xa4, 0xa5, 0}},
	{ASHLW, Ycl, Yrl, Yml, movDoubleShift, [4]uint8{Pe, 0xa4, 0xa5, 0}},
	{ASHLW, Ycx, Yrl, Yml, movDoubleShift, [4]uint8{Pe, 0xa4, 0xa5, 0}},
	{ASHRW, Yi8, Yrl, Yml, movDoubleShift, [4]uint8{Pe, 0xac, 0xad, 0}},
	{ASHRW, Ycl, Yrl, Yml, movDoubleShift, [4]uint8{Pe, 0xac, 0xad, 0}},
	{ASHRW, Ycx, Yrl, Yml, movDoubleShift, [4]uint8{Pe, 0xac, 0xad, 0}},

	// load TLS base
	{AMOVL, Ytls, Ynone, Yrl, movTLSReg, [4]uint8{0, 0, 0, 0}},
	{AMOVQ, Ytls, Ynone, Yrl, movTLSReg, [4]uint8{0, 0, 0, 0}},
}
