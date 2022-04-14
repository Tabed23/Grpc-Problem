package models


type Binder struct {
	BlockA               []string `protobuf:"bytes,1,rep,name=block_a,json=blockA,proto3" json:"block_a,omitempty"`
	BlockB               []string `protobuf:"bytes,2,rep,name=block_b,json=blockB,proto3" json:"block_b,omitempty"`
	BlockC               []string `protobuf:"bytes,3,rep,name=block_c,json=blockC,proto3" json:"block_c,omitempty"`
	BlockD               []string `protobuf:"bytes,4,rep,name=block_d,json=blockD,proto3" json:"block_d,omitempty"`
	BlockE               []string `protobuf:"bytes,5,rep,name=block_e,json=blockE,proto3" json:"block_e,omitempty"`
	BlockF               []string `protobuf:"bytes,6,rep,name=block_f,json=blockF,proto3" json:"block_f,omitempty"`
	BlockG               []string `protobuf:"bytes,7,rep,name=block_g,json=blockG,proto3" json:"block_g,omitempty"`
	BlockH               []string `protobuf:"bytes,8,rep,name=block_h,json=blockH,proto3" json:"block_h,omitempty"`
	BlockI               []string `protobuf:"bytes,9,rep,name=block_i,json=blockI,proto3" json:"block_i,omitempty"`
	BlockJ               []string `protobuf:"bytes,10,rep,name=block_j,json=blockJ,proto3" json:"block_j,omitempty"`
	BlockK               []string `protobuf:"bytes,11,rep,name=block_k,json=blockK,proto3" json:"block_k,omitempty"`
}
type BinderResponse struct {
	Message []string `protobuf:"bytes,1,rep,name=block_a,json=blockA,proto3" json:"block_a,omitempty"`
}