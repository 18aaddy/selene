package types




type ByteList struct{

}

type Address = [20] byte;
type Bytes32 =[32] byte;
type LogsBloom =[256] byte;
type BLSPubKey =[48] byte;
type SignatureBytes= [96] byte;
type Transaction =ByteList;//1073741824