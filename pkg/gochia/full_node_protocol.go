package gochia

/*
@dataclass(frozen=True)
@streamable
	class NewPeak(Streamable):
	header_hash: bytes32
	height: uint32
	weight: uint128
	fork_point_with_previous_peak: uint32
	unfinished_reward_block_hash: bytes32
*/
type NewPeak struct {
	header_hash                   bytes32
	height                        uint32
	weight                        uint128.Uint128
	fork_point_with_previous_peak uint32
	unfinished_reward_block_hash  bytes32
}

/*
@dataclass(frozen=True)
@streamable
class NewTransaction(Streamable):
	transaction_id: bytes32
	cost: uint64
	fees: uint64
*/
type NewTransaction struct {
	transaction_id bytes32
	cost           uint64
	fees           uint64
}

/*
@dataclass(frozen=True)
@streamable
class RequestTransaction(Streamable):
transaction_id: bytes32
*/
type  RequestTransaction struct {
	transaction_id bytes32

}

/*
@dataclass(frozen=True)
@streamable
class RespondTransaction(Streamable):
transaction: SpendBundle
*/
type RespondTransaction struct {
	transaction SpendBundle

}

/*
@dataclass(frozen=True)
@streamable
class RequestProofOfWeight(Streamable):
total_number_of_blocks: uint32
tip: bytes32
*/
type RequestProofOfWeight struct {
	total_number_of_blocks uint32
	tip bytes32
}

/*
@dataclass(frozen=True)
@streamable
class RespondProofOfWeight(Streamable):
wp: WeightProof
tip: bytes32
*/
type RespondProofOfWeight struct {
	wp WeightProof
	tip bytes32
}

/*
@dataclass(frozen=True)
@streamable
class RequestBlock(Streamable):
height: uint32
include_transaction_block: bool
*/
type RequestBlock struct {
	height uint32
	include_transaction_block bool
}

/*
@dataclass(frozen=True)
@streamable
class RejectBlock(Streamable):
height: uint32
*/
type RejectBlock struct {
	height uint32
}

/*
@dataclass(frozen=True)
@streamable
class RequestBlocks(Streamable):
start_height: uint32
end_height: uint32
include_transaction_block: bool
*/
type RequestBlocks struct {
	start_height uint32
	end_height uint32
	include_transaction_block bool
}

/*
@dataclass(frozen=True)
@streamable
class RespondBlocks(Streamable):
start_height: uint32
end_height: uint32
blocks: List[FullBlock]
*/
type RespondBlocks struct {
	start_height uint32
	end_height uint32
	blocks []FullBlock
}

/*
@dataclass(frozen=True)
@streamable
class RejectBlocks(Streamable):
start_height: uint32
end_height: uint32
*/
type RejectBlocks struct {
	start_height uint32
	end_height uint32
}

/*
@dataclass(frozen=True)
@streamable
class RespondBlock(Streamable):
block: FullBlock
*/
type  RespondBlock struct {
	block FullBlock
}

/*
@dataclass(frozen=True)
@streamable
class NewUnfinishedBlock(Streamable):
unfinished_reward_hash: bytes32
*/
type NewUnfinishedBlock struct {
	unfinished_reward_hash bytes32
}

/*
@dataclass(frozen=True)
@streamable
class RequestUnfinishedBlock(Streamable):
unfinished_reward_hash: bytes32
*/
type RequestUnfinishedBlock struct {
	unfinished_reward_hash bytes32
}

/*
@dataclass(frozen=True)
@streamable
class RespondUnfinishedBlock(Streamable):
unfinished_block: UnfinishedBlock
*/
type RespondUnfinishedBlock struct {
	unfinished_block UnfinishedBlock
}

/*
@dataclass(frozen=True)
@streamable
class NewSignagePointOrEndOfSubSlot(Streamable):
prev_challenge_hash: Optional[bytes32]
challenge_hash: bytes32
index_from_challenge: uint8
last_rc_infusion: bytes32
*/
type NewSignagePointOrEndOfSubSlot struct {
	prev_challenge_hash Optional[bytes32]
	challenge_hash bytes32
	index_from_challenge uint8
	last_rc_infusion bytes32
}

/*
@dataclass(frozen=True)
@streamable
class RequestSignagePointOrEndOfSubSlot(Streamable):
challenge_hash: bytes32
index_from_challenge: uint8
last_rc_infusion: bytes32
*/
type RequestSignagePointOrEndOfSubSlot struct {
	challenge_hash bytes32
	index_from_challenge uint8
	last_rc_infusion bytes32
}

/*
@dataclass(frozen=True)
@streamable
class RespondSignagePoint(Streamable):
index_from_challenge: uint8
challenge_chain_vdf: VDFInfo
challenge_chain_proof: VDFProof
reward_chain_vdf: VDFInfo
reward_chain_proof: VDFProof
*/
type RespondSignagePoint struct {
	index_from_challenge uint8
	challenge_chain_vdf VDFInfo
	challenge_chain_proof VDFProof
	reward_chain_vdf VDFInfo
	reward_chain_proof VDFProof
}

/*
@dataclass(frozen=True)
@streamable
class RespondEndOfSubSlot(Streamable):
end_of_slot_bundle: EndOfSubSlotBundle
*/
type RespondEndOfSubSlot struct {
	end_of_slot_bundle EndOfSubSlotBundle
}

/*
@dataclass(frozen=True)
@streamable
class RequestMempoolTransactions(Streamable):
filter: bytes
*/
type RequestMempoolTransactions struct {
	filter []byte
}

/*
@dataclass(frozen=True)
@streamable
class NewCompactVDF(Streamable):
height: uint32
header_hash: bytes32
field_vdf: uint8
vdf_info: VDFInfo
*/
type NewCompactVDF struct {
	height uint32
	header_hash bytes32
	field_vdf uint8
	vdf_info VDFInfo
}

/*
@dataclass(frozen=True)
@streamable
class RequestCompactVDF(Streamable):
height: uint32
header_hash: bytes32
field_vdf: uint8
vdf_info: VDFInfo
*/
type RequestCompactVDF struct {
	height uint32
	header_hash bytes32
	field_vdf uint8
	vdf_info VDFInfo
}

/*
@dataclass(frozen=True)
@streamable
class RespondCompactVDF(Streamable):
height: uint32
header_hash: bytes32
field_vdf: uint8
vdf_info: VDFInfo
vdf_proof: VDFProof
*/
type RespondCompactVDF struct {
	height: uint32
	header_hash: bytes32
	field_vdf: uint8
	vdf_info: VDFInfo
	vdf_proof: VDFProof
}

/*
@dataclass(frozen=True)
@streamable
class RequestPeers(Streamable):
"""
Return full list of peers
"""
*/
type RequestPeers struct {
}

/*
@dataclass(frozen=True)
@streamable
class RespondPeers(Streamable):
peer_list: List[TimestampedPeerInfo]
*/
type RespondPeers struct {
	peer_list []TimestampedPeerInfo
}
