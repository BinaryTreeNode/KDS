/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by codecgen - DO NOT EDIT.

package componentconfig

import (
	"errors"
	"fmt"
	codec1978 "github.com/ugorji/go/codec"
	pkg1_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"reflect"
	"runtime"
	time "time"
)

const (
	// ----- content types ----
	codecSelferCcUTF81234 = 1
	codecSelferCcRAW1234  = 0
	// ----- value types used ----
	codecSelferValueTypeArray1234  = 10
	codecSelferValueTypeMap1234    = 9
	codecSelferValueTypeString1234 = 6
	codecSelferValueTypeInt1234    = 2
	codecSelferValueTypeUint1234   = 3
	codecSelferValueTypeFloat1234  = 4
)

var (
	codecSelferBitsize1234                         = uint8(reflect.TypeOf(uint(0)).Bits())
	errCodecSelferOnlyMapOrArrayEncodeToStruct1234 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer1234 struct{}

func init() {
	if codec1978.GenVersion != 8 {
		_, file, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate file: %v",
			8, codec1978.GenVersion, file)
		panic(err)
	}
	if false { // reference the types, but skip this branch at build/run time
		var v0 pkg1_v1.TypeMeta
		var v1 time.Duration
		_, _ = v0, v1
	}
}

func (x *DeschedulerConfiguration) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [9]bool
			_ = yyq2
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Kind != ""
			yyq2[1] = x.APIVersion != ""
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(9)
			} else {
				var yynn2 = 7
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.WriteMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[0] {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						r.EncodeString(codecSelferCcUTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferCcUTF81234, "")
				}
			} else {
				if yyq2[0] {
					r.WriteMapElemKey()
					r.EncStructFieldKey(codecSelferValueTypeString1234, `kind`)
					r.WriteMapElemValue()
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						r.EncodeString(codecSelferCcUTF81234, string(x.Kind))
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[1] {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else {
						r.EncodeString(codecSelferCcUTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferCcUTF81234, "")
				}
			} else {
				if yyq2[1] {
					r.WriteMapElemKey()
					r.EncStructFieldKey(codecSelferValueTypeString1234, `apiVersion`)
					r.WriteMapElemValue()
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} else {
						r.EncodeString(codecSelferCcUTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym10 := z.EncBinary()
				_ = yym10
				if false {
				} else if yyxt10 := z.Extension(z.I2Rtid(x.DeschedulingInterval)); yyxt10 != nil {
					z.EncExtension(x.DeschedulingInterval, yyxt10)
				} else {
					r.EncodeInt(int64(x.DeschedulingInterval))
				}
			} else {
				r.WriteMapElemKey()
				r.EncStructFieldKey(codecSelferValueTypeString1234, `DeschedulingInterval`)
				r.WriteMapElemValue()
				yym11 := z.EncBinary()
				_ = yym11
				if false {
				} else if yyxt11 := z.Extension(z.I2Rtid(x.DeschedulingInterval)); yyxt11 != nil {
					z.EncExtension(x.DeschedulingInterval, yyxt11)
				} else {
					r.EncodeInt(int64(x.DeschedulingInterval))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym13 := z.EncBinary()
				_ = yym13
				if false {
				} else {
					r.EncodeString(codecSelferCcUTF81234, string(x.KubeconfigFile))
				}
			} else {
				r.WriteMapElemKey()
				r.EncStructFieldKey(codecSelferValueTypeString1234, `KubeconfigFile`)
				r.WriteMapElemValue()
				yym14 := z.EncBinary()
				_ = yym14
				if false {
				} else {
					r.EncodeString(codecSelferCcUTF81234, string(x.KubeconfigFile))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym16 := z.EncBinary()
				_ = yym16
				if false {
				} else {
					r.EncodeString(codecSelferCcUTF81234, string(x.PolicyConfigFile))
				}
			} else {
				r.WriteMapElemKey()
				r.EncStructFieldKey(codecSelferValueTypeString1234, `PolicyConfigFile`)
				r.WriteMapElemValue()
				yym17 := z.EncBinary()
				_ = yym17
				if false {
				} else {
					r.EncodeString(codecSelferCcUTF81234, string(x.PolicyConfigFile))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym19 := z.EncBinary()
				_ = yym19
				if false {
				} else {
					r.EncodeBool(bool(x.DryRun))
				}
			} else {
				r.WriteMapElemKey()
				r.EncStructFieldKey(codecSelferValueTypeString1234, `DryRun`)
				r.WriteMapElemValue()
				yym20 := z.EncBinary()
				_ = yym20
				if false {
				} else {
					r.EncodeBool(bool(x.DryRun))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym22 := z.EncBinary()
				_ = yym22
				if false {
				} else {
					r.EncodeString(codecSelferCcUTF81234, string(x.NodeSelector))
				}
			} else {
				r.WriteMapElemKey()
				r.EncStructFieldKey(codecSelferValueTypeString1234, `NodeSelector`)
				r.WriteMapElemValue()
				yym23 := z.EncBinary()
				_ = yym23
				if false {
				} else {
					r.EncodeString(codecSelferCcUTF81234, string(x.NodeSelector))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym25 := z.EncBinary()
				_ = yym25
				if false {
				} else {
					r.EncodeInt(int64(x.MaxNoOfPodsToEvictPerNode))
				}
			} else {
				r.WriteMapElemKey()
				r.EncStructFieldKey(codecSelferValueTypeString1234, `MaxNoOfPodsToEvictPerNode`)
				r.WriteMapElemValue()
				yym26 := z.EncBinary()
				_ = yym26
				if false {
				} else {
					r.EncodeInt(int64(x.MaxNoOfPodsToEvictPerNode))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym28 := z.EncBinary()
				_ = yym28
				if false {
				} else {
					r.EncodeBool(bool(x.EvictLocalStoragePods))
				}
			} else {
				r.WriteMapElemKey()
				r.EncStructFieldKey(codecSelferValueTypeString1234, `EvictLocalStoragePods`)
				r.WriteMapElemValue()
				yym29 := z.EncBinary()
				_ = yym29
				if false {
				} else {
					r.EncodeBool(bool(x.EvictLocalStoragePods))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}

func (x *DeschedulerConfiguration) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1234 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray1234 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct1234)
		}
	}
}

func (x *DeschedulerConfiguration) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3 := z.StringView(r.DecStructFieldKey(codecSelferValueTypeString1234, z.DecScratchArrayBuffer()))
		r.ReadMapElemValue()
		switch yys3 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				yyv4 := &x.Kind
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				yyv6 := &x.APIVersion
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*string)(yyv6)) = r.DecodeString()
				}
			}
		case "DeschedulingInterval":
			if r.TryDecodeAsNil() {
				x.DeschedulingInterval = 0
			} else {
				yyv8 := &x.DeschedulingInterval
				yym9 := z.DecBinary()
				_ = yym9
				if false {
				} else if yyxt9 := z.Extension(z.I2Rtid(yyv8)); yyxt9 != nil {
					z.DecExtension(yyv8, yyxt9)
				} else {
					*((*int64)(yyv8)) = int64(r.DecodeInt(64))
				}
			}
		case "KubeconfigFile":
			if r.TryDecodeAsNil() {
				x.KubeconfigFile = ""
			} else {
				yyv10 := &x.KubeconfigFile
				yym11 := z.DecBinary()
				_ = yym11
				if false {
				} else {
					*((*string)(yyv10)) = r.DecodeString()
				}
			}
		case "PolicyConfigFile":
			if r.TryDecodeAsNil() {
				x.PolicyConfigFile = ""
			} else {
				yyv12 := &x.PolicyConfigFile
				yym13 := z.DecBinary()
				_ = yym13
				if false {
				} else {
					*((*string)(yyv12)) = r.DecodeString()
				}
			}
		case "DryRun":
			if r.TryDecodeAsNil() {
				x.DryRun = false
			} else {
				yyv14 := &x.DryRun
				yym15 := z.DecBinary()
				_ = yym15
				if false {
				} else {
					*((*bool)(yyv14)) = r.DecodeBool()
				}
			}
		case "NodeSelector":
			if r.TryDecodeAsNil() {
				x.NodeSelector = ""
			} else {
				yyv16 := &x.NodeSelector
				yym17 := z.DecBinary()
				_ = yym17
				if false {
				} else {
					*((*string)(yyv16)) = r.DecodeString()
				}
			}
		case "MaxNoOfPodsToEvictPerNode":
			if r.TryDecodeAsNil() {
				x.MaxNoOfPodsToEvictPerNode = 0
			} else {
				yyv18 := &x.MaxNoOfPodsToEvictPerNode
				yym19 := z.DecBinary()
				_ = yym19
				if false {
				} else {
					*((*int)(yyv18)) = int(r.DecodeInt(codecSelferBitsize1234))
				}
			}
		case "EvictLocalStoragePods":
			if r.TryDecodeAsNil() {
				x.EvictLocalStoragePods = false
			} else {
				yyv20 := &x.EvictLocalStoragePods
				yym21 := z.DecBinary()
				_ = yym21
				if false {
				} else {
					*((*bool)(yyv20)) = r.DecodeBool()
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	r.ReadMapEnd()
}

func (x *DeschedulerConfiguration) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj22 int
	var yyb22 bool
	var yyhl22 bool = l >= 0
	yyj22++
	if yyhl22 {
		yyb22 = yyj22 > l
	} else {
		yyb22 = r.CheckBreak()
	}
	if yyb22 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		yyv23 := &x.Kind
		yym24 := z.DecBinary()
		_ = yym24
		if false {
		} else {
			*((*string)(yyv23)) = r.DecodeString()
		}
	}
	yyj22++
	if yyhl22 {
		yyb22 = yyj22 > l
	} else {
		yyb22 = r.CheckBreak()
	}
	if yyb22 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		yyv25 := &x.APIVersion
		yym26 := z.DecBinary()
		_ = yym26
		if false {
		} else {
			*((*string)(yyv25)) = r.DecodeString()
		}
	}
	yyj22++
	if yyhl22 {
		yyb22 = yyj22 > l
	} else {
		yyb22 = r.CheckBreak()
	}
	if yyb22 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.DeschedulingInterval = 0
	} else {
		yyv27 := &x.DeschedulingInterval
		yym28 := z.DecBinary()
		_ = yym28
		if false {
		} else if yyxt28 := z.Extension(z.I2Rtid(yyv27)); yyxt28 != nil {
			z.DecExtension(yyv27, yyxt28)
		} else {
			*((*int64)(yyv27)) = int64(r.DecodeInt(64))
		}
	}
	yyj22++
	if yyhl22 {
		yyb22 = yyj22 > l
	} else {
		yyb22 = r.CheckBreak()
	}
	if yyb22 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.KubeconfigFile = ""
	} else {
		yyv29 := &x.KubeconfigFile
		yym30 := z.DecBinary()
		_ = yym30
		if false {
		} else {
			*((*string)(yyv29)) = r.DecodeString()
		}
	}
	yyj22++
	if yyhl22 {
		yyb22 = yyj22 > l
	} else {
		yyb22 = r.CheckBreak()
	}
	if yyb22 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.PolicyConfigFile = ""
	} else {
		yyv31 := &x.PolicyConfigFile
		yym32 := z.DecBinary()
		_ = yym32
		if false {
		} else {
			*((*string)(yyv31)) = r.DecodeString()
		}
	}
	yyj22++
	if yyhl22 {
		yyb22 = yyj22 > l
	} else {
		yyb22 = r.CheckBreak()
	}
	if yyb22 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.DryRun = false
	} else {
		yyv33 := &x.DryRun
		yym34 := z.DecBinary()
		_ = yym34
		if false {
		} else {
			*((*bool)(yyv33)) = r.DecodeBool()
		}
	}
	yyj22++
	if yyhl22 {
		yyb22 = yyj22 > l
	} else {
		yyb22 = r.CheckBreak()
	}
	if yyb22 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.NodeSelector = ""
	} else {
		yyv35 := &x.NodeSelector
		yym36 := z.DecBinary()
		_ = yym36
		if false {
		} else {
			*((*string)(yyv35)) = r.DecodeString()
		}
	}
	yyj22++
	if yyhl22 {
		yyb22 = yyj22 > l
	} else {
		yyb22 = r.CheckBreak()
	}
	if yyb22 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.MaxNoOfPodsToEvictPerNode = 0
	} else {
		yyv37 := &x.MaxNoOfPodsToEvictPerNode
		yym38 := z.DecBinary()
		_ = yym38
		if false {
		} else {
			*((*int)(yyv37)) = int(r.DecodeInt(codecSelferBitsize1234))
		}
	}
	yyj22++
	if yyhl22 {
		yyb22 = yyj22 > l
	} else {
		yyb22 = r.CheckBreak()
	}
	if yyb22 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.EvictLocalStoragePods = false
	} else {
		yyv39 := &x.EvictLocalStoragePods
		yym40 := z.DecBinary()
		_ = yym40
		if false {
		} else {
			*((*bool)(yyv39)) = r.DecodeBool()
		}
	}
	for {
		yyj22++
		if yyhl22 {
			yyb22 = yyj22 > l
		} else {
			yyb22 = r.CheckBreak()
		}
		if yyb22 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj22-1, "")
	}
	r.ReadArrayEnd()
}