import { module } from 'qunit';
import { TestNumber } from './number';
import { TestSlice } from './slice'



module(`stdgo/builtin`, () => {
    TestNumber()
    TestSlice()
})