// Copyright (c) 2021, Oleg Romanenko (oleg@romanenko.ro)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lz4

import (
	"unsafe"
)

/*
#include<stdint.h>
#include <string.h>

void set_header(uint32_t size, char* dest) {
    uint32_t *header = (uint32_t*)dest;
    *header = size;
}

uint32_t get_header(char* data) {
    uint32_t orig_size = 0;
    memcpy((char*)&orig_size, data, sizeof(uint32_t));
	return orig_size;
}

extern int LZ4_compressBound(int inputSize);
extern int LZ4_compress_default(const char* source, char* dest, int sourceSize, int maxDestSize);
extern int LZ4_decompress_safe (const char* source, char* dest, int compressedSize, int maxDecompressedSize);
*/
import "C"

const (
	headerSize   = 4
	maxBlockSize = 1024 * 1024
)

func CompressUint32Block(in []byte) []byte {
	if len(in) == 0 {
		return nil
	}

	boundSize := C.LZ4_compressBound(C.int(len(in)))
	if boundSize == 0 {
		return nil
	}

	boundSize += headerSize

	out := make([]byte, boundSize)

	inPtr := (*C.char)(unsafe.Pointer(&in[0]))
	outUnsafePtr := unsafe.Pointer(&out[0])
	outPtr := (*C.char)(outUnsafePtr)

	size := C.LZ4_compress_default(inPtr, (*C.char)(unsafe.Add(outUnsafePtr, headerSize)),
		C.int(len(in)),
		C.int(len(out)-headerSize))

	if size == 0 {
		return nil
	}

	if size+headerSize != boundSize {
		out = out[:size+headerSize]
	}

	C.set_header(C.uint32_t(len(in)), outPtr)

	return out
}

func DecompressUint32Block(in []byte) []byte {
	if len(in) <= headerSize {
		return nil
	}

	inUnsafePtr := unsafe.Pointer(&in[0])
	inPtr := (*C.char)(inUnsafePtr)

	origSize := C.get_header(inPtr)

	if origSize > maxBlockSize {
		return nil
	}

	out := make([]byte, origSize)
	outPtr := (*C.char)(unsafe.Pointer(&out[0]))

	size := C.LZ4_decompress_safe((*C.char)(unsafe.Add(inUnsafePtr, headerSize)), outPtr,
		C.int(len(in)-headerSize), C.int(origSize))

	if size < 0 {
		return nil
	}

	return out
}
