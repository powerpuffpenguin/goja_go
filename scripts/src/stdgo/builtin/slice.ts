
import { test } from 'qunit';
import * as builtin from 'stdgo/builtin';

export function TestSlice() {
    test(`slice`, (assert) => {
        let b = builtin.IntSlice(5, 10)
        assert.equal(builtin.len(b), 5)
        assert.equal(builtin.cap(b), 10)

        assert.equal(b.Len(), 5)
        assert.equal(b.Cap(), 10)

    })
}
