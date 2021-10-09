
import { test } from 'qunit';
import * as builtin from 'stdgo/builtin';

export function TestNumber() {
    test(`const`, (assert) => {
        assert.equal(builtin.MaxInt64.String(), '9223372036854775807')
        assert.equal(builtin.MaxInt32.String(), '2147483647')
        assert.equal(builtin.MaxInt16.String(), '32767')
        assert.equal(builtin.MaxInt8.String(), '127')
        assert.equal(builtin.MaxUint64.String(), '18446744073709551615')
        assert.equal(builtin.MaxUint32.String(), '4294967295')
        assert.equal(builtin.MaxUint16.String(), '65535')
        assert.equal(builtin.MaxUint8.String(), '255')
        assert.equal(builtin.MaxFloat64.String(), '1.7976931348623157e+308')
        assert.equal(builtin.MaxFloat32.String(), '3.4028235e+38')
        assert.equal(builtin.MinInt64.String(), '-9223372036854775808')
        assert.equal(builtin.MinInt32.String(), '-2147483648')
        assert.equal(builtin.MinInt16.String(), '-32768')
        assert.equal(builtin.MinInt8.String(), '-128')


        assert.equal(builtin.MaxInt64.Sub(1).String(), "9223372036854775806")
        assert.equal(builtin.MaxInt32.Sub(1).String(), "2147483646")
        assert.equal(builtin.MaxInt16.Sub(1).String(), "32766")
        assert.equal(builtin.MaxInt8.Sub(1).String(), "126")
        assert.equal(builtin.MinInt64.Add(1).String(), "-9223372036854775807")
        assert.equal(builtin.MinInt32.Add(1).String(), "-2147483647")
        assert.equal(builtin.MinInt16.Add(1).String(), "-32767")
        assert.equal(builtin.MinInt8.Add(1).String(), "-127")

        assert.equal(builtin.MaxUint64.Sub(builtin.Int(2)).String(), "18446744073709551613")
        assert.equal(builtin.MaxUint32.Sub(builtin.Int8(2)).String(), "4294967293")
        assert.equal(builtin.MaxUint16.Sub(builtin.Uint(2)).String(), "65533")
        assert.equal(builtin.MaxUint8.Sub(builtin.Float32(2)).String(), "253")

        assert.true(builtin.isInt64(builtin.MaxInt64))
        assert.true(builtin.isInt32(builtin.MaxInt32))
        assert.true(builtin.isInt16(builtin.MaxInt16))
        assert.true(builtin.isInt8(builtin.MaxInt8))
        assert.true(builtin.isUint64(builtin.MaxUint64))
        assert.true(builtin.isUint32(builtin.MaxUint32))
        assert.true(builtin.isUint16(builtin.MaxUint16))
        assert.true(builtin.isUint8(builtin.MaxUint8))
        assert.true(builtin.isFloat64(builtin.MaxFloat64))
        assert.true(builtin.isFloat32(builtin.MaxFloat32))
        assert.true(builtin.isInt64(builtin.MinInt64))
        assert.true(builtin.isInt32(builtin.MinInt32))
        assert.true(builtin.isInt16(builtin.MinInt16))
        assert.true(builtin.isInt8(builtin.MinInt8))

        assert.true(typeof builtin.IntSize === "number")
        assert.true(typeof builtin.MaxInt64 === "object")
        assert.true(typeof builtin.MaxInt32 === "object")
        assert.true(typeof builtin.MaxInt16 === "object")
        assert.true(typeof builtin.MaxInt8 === "object")
        assert.true(typeof builtin.MaxUint64 === "object")
        assert.true(typeof builtin.MaxUint32 === "object")
        assert.true(typeof builtin.MaxUint16 === "object")
        assert.true(typeof builtin.MaxUint8 === "object")
        assert.true(typeof builtin.MaxFloat64 === "object")
        assert.true(typeof builtin.MaxFloat32 === "object")
        assert.true(typeof builtin.MinInt64 === "object")
        assert.true(typeof builtin.MinInt32 === "object")
        assert.true(typeof builtin.MinInt16 === "object")
        assert.true(typeof builtin.MinInt8 === "object")
    })
    test(`value`, (assert) => {
        if (builtin.IntSize == 64) {
            assert.equal(builtin.Uint('18446744073709551615').String(), '18446744073709551615')
        } else if (builtin.IntSize == 32) {
            assert.equal(builtin.Uint('4294967295').String(), '4294967295')
        }
    })
}
