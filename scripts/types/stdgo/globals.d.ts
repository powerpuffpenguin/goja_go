import { Error, Native } from "stdgo/builtin";

class GoError extends Native {
    private readonly __GoError: GoError
    private constructor()
    name: string
    message: string
    stack?: string
    value: Error
}