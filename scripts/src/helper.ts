import * as QUnit from 'qunit'

export class Helper {
    /**
     * 是否需要打印 細節
     */
    details = false
    private static instance_: Helper
    static get instance(): Helper {
        if (!Helper.instance_) {
            Helper.instance_ = new Helper()
        }
        return Helper.instance_
    }
    private constructor() {
        // 開始 測試
        QUnit.begin((details) => {
            if (!this.details) {
                return
            }
            console.info(`begin ${details.totalTests} testing ...`)
        })

        // 測試 完成
        QUnit.done((details) => {
            // 匯總 測試信息 並打印
            let message = `done ${details.total} assert,`
            if (details.passed > 0) {
                message += ` ${details.passed} passed,`
            }
            if (details.failed > 0) {
                message += ` ${details.failed} failed,`
            }
            message += ` used ${details.runtime} milliseconds.`
            console.log(message)
        })
        // 開始 模塊測試
        QUnit.moduleStart((details) => {
            if (!this.details) {
                return
            }
            console.log(`- module ${details.name}`)
        })
        // 模塊測試 結束
        QUnit.moduleDone((details) => {
            if (!this.details) {
                return
            }
            // 匯總 測試信息 並打印
            let message
            if (details.name.length == 0) {
                message = `- module : done ${details.total} assert,`
            } else {
                message = `- module ${details.name} : done ${details.total} assert,`
            }
            if (details.passed > 0) {
                message += ` ${details.passed} passed,`
            }
            if (details.failed > 0) {
                message += ` ${details.failed} failed,`
            }
            message += ` used ${details.runtime} milliseconds.`
            console.log(message)
            console.log()
        })
        // 開始 單件測試
        QUnit.testStart((details) => {
            if (!this.details) {
                return
            }
            console.log(`* unit ${details.name}`)
        })
        // 單件測試 完成
        QUnit.testDone((details) => {
            if (!this.details) {
                return
            }
            // 匯總 測試信息 並打印
            let message = `* unit ${details.name} : done ${details.total} assert,`

            if (details.passed > 0) {
                message += ` ${details.passed} passed,`
            }
            if (details.failed > 0) {
                message += ` ${details.failed} failed,`
            }
            message += ` used ${details.runtime} milliseconds.`
            console.log(message)
        })
        const keys = new Set<string>()
        QUnit.log((details) => {
            if (this.details) {
                if (details.result) {
                    console.log(`@ passed ${details.name}`)
                    return
                }
            } else {
                if (details.result) {
                    // 通過 不需要細節直接返回
                    return
                }
                if (!keys.has(details.module)) {
                    if (keys.size != 0) {
                        console.warn()
                    }
                    console.warn(`- module ${details.module}`)
                    keys.add(details.module)
                }
            }
            this._failed(details)
        })
    }
    private _failed(details: QUnit.LogDetails) {
        console.warn(`! failed ${details.name}`)
        if (undefined != details.message) {
            console.warn(`    message  : ${details.message}`)
        }
        console.warn(`    expected : ${details.expected}`)
        console.warn(`    actual   : ${details.actual}`)
        if (details.source != undefined) {
            console.warn(details.source)
        }
    }
    start = QUnit.start
}
