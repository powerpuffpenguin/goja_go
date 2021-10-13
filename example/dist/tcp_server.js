"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    Object.defineProperty(o, k2, { enumerable: true, get: function() { return m[k]; } });
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (_) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
Object.defineProperty(exports, "__esModule", { value: true });
// This is an example of an echo server
var builtin = __importStar(require("stdgo/builtin"));
var net = __importStar(require("stdgo/net"));
var io = __importStar(require("stdgo/io"));
var time = __importStar(require("stdgo/time"));
var fmt = __importStar(require("stdgo/fmt"));
function main() {
    return __awaiter(this, void 0, void 0, function () {
        var addr, l, tempDelay, _loop_1, state_1;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    addr = "127.0.0.1:10000";
                    return [4 /*yield*/, builtin.async(net.Listen, 'tcp', addr)];
                case 1:
                    l = _a.sent();
                    console.log("listen on->", addr);
                    tempDelay = 0;
                    _loop_1 = function () {
                        var c_1, remote_1, e_1, ne, duration;
                        return __generator(this, function (_a) {
                            switch (_a.label) {
                                case 0:
                                    _a.trys.push([0, 2, , 5]);
                                    return [4 /*yield*/, builtin.async(l.Accept)];
                                case 1:
                                    c_1 = _a.sent();
                                    remote_1 = c_1.RemoteAddr().String();
                                    console.log("one in:", remote_1);
                                    builtin.async(io.Copy, c_1, c_1).catch(function (e) {
                                        console.log("io.Copy err->", e);
                                    }).finally(function () {
                                        console.log("one out:", remote_1);
                                        c_1.Close();
                                    });
                                    return [3 /*break*/, 5];
                                case 2:
                                    e_1 = _a.sent();
                                    if (!(e_1 instanceof GoError && net.isError(e_1.value))) return [3 /*break*/, 4];
                                    ne = e_1.value;
                                    if (!ne.Temporary()) return [3 /*break*/, 4];
                                    if (tempDelay == 0) {
                                        tempDelay = 5;
                                    }
                                    else {
                                        tempDelay *= 2;
                                        if (tempDelay > 1000) {
                                            tempDelay = 1000;
                                        }
                                    }
                                    duration = time.Duration(builtin.Int64(time.Millisecond).Mul(tempDelay));
                                    fmt.Printf("Accept error: %v; retrying in %v\n", ne, tempDelay);
                                    return [4 /*yield*/, builtin.async(time.Sleep, duration)];
                                case 3:
                                    _a.sent();
                                    return [2 /*return*/, "continue"];
                                case 4:
                                    console.log("accept err->", e_1);
                                    return [2 /*return*/, { value: void 0 }];
                                case 5: return [2 /*return*/];
                            }
                        });
                    };
                    _a.label = 2;
                case 2:
                    if (!true) return [3 /*break*/, 4];
                    return [5 /*yield**/, _loop_1()];
                case 3:
                    state_1 = _a.sent();
                    if (typeof state_1 === "object")
                        return [2 /*return*/, state_1.value];
                    return [3 /*break*/, 2];
                case 4: return [2 /*return*/];
            }
        });
    });
}
main().catch(function (e) {
    console.log("listen err->", e);
});
