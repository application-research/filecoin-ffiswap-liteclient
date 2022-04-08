package ffi

import (
	"io"
	"os"
	"unsafe"

	"github.com/application-research/filecoin-ffiswap-liteclient/cgo"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/network"
	"github.com/filecoin-project/specs-actors/v7/actors/runtime/proof"
	"github.com/ipfs/go-cid"
)

type FVM struct {
	executor unsafe.Pointer
}

const (
	applyExplicit = iota
	applyImplicit
)

type FVMOpts struct {
	FVMVersion uint64
	Externs    cgo.Externs

	Epoch          abi.ChainEpoch
	BaseFee        abi.TokenAmount
	BaseCircSupply abi.TokenAmount
	NetworkVersion network.Version
	StateBase      cid.Cid
	Manifest       cid.Cid
}

type SortedPrivateSectorInfo struct{ abi.SectorNumber }
type PrivateSectorInfo struct {
	abi.SectorNumber
	CacheDirPath     string
	PoStProofType    abi.RegisteredPoStProof
	SealedSectorPath string
	SectorInfo       proof.SectorInfo
}
type FallbackChallenges struct {
	Sectors    []abi.SectorNumber
	Challenges map[abi.SectorNumber][]uint64
}

type FilPartitionProof struct {
	ProofLen       uint
	ProofPtr       []byte
	ref566a2be6    interface{}
	allocs566a2be6 interface{}
}

func GenerateUnsealedCID(abi.RegisteredSealProof, []abi.PieceInfo) (cid.Cid, error) {
	return cid.Undef, nil
}
func GeneratePieceCIDFromFile(abi.RegisteredSealProof, io.Reader, abi.UnpaddedPieceSize) (cid.Cid, error) {
	return cid.Undef, nil
}
func UnsealRange(abi.RegisteredSealProof, string, *os.File, *os.File, abi.SectorNumber, abi.ActorID, abi.SealRandomness, cid.Cid, uint64, uint64) error {
	return nil
}
func SealPreCommitPhase1(abi.RegisteredSealProof, string, string, string, abi.SectorNumber, abi.ActorID, abi.SealRandomness, []abi.PieceInfo) ([]byte, error) {
	return nil, nil
}
func SealPreCommitPhase2([]byte, string, string) (cid.Cid, cid.Cid, error) {
	return cid.Undef, cid.Undef, nil
}
func SealCommitPhase1(abi.RegisteredSealProof, cid.Cid, cid.Cid, string, string, abi.SectorNumber, abi.ActorID, abi.SealRandomness, abi.InteractiveSealRandomness, []abi.PieceInfo) ([]byte, error) {
	return nil, nil
}
func SealCommitPhase2([]byte, abi.SectorNumber, abi.ActorID) ([]byte, error) {
	return nil, nil
}
func ClearCache(uint64, string) error {
	return nil
}
func GenerateWinningPoSt(abi.ActorID, SortedPrivateSectorInfo, abi.PoStRandomness) ([]proof.PoStProof, error) {
	return nil, nil
}
func GenerateWindowPoSt(abi.ActorID, SortedPrivateSectorInfo, abi.PoStRandomness) ([]proof.PoStProof, []abi.SectorNumber, error) {
	return nil, nil, nil
}
func NewSortedPrivateSectorInfo(...PrivateSectorInfo) SortedPrivateSectorInfo {
	return SortedPrivateSectorInfo{}
}
func VerifySeal(proof.SealVerifyInfo) (bool, error) {
	return true, nil
}
func VerifyWinningPoSt(proof.WinningPoStVerifyInfo) (bool, error) {
	return true, nil
}
func VerifyWindowPoSt(proof.WindowPoStVerifyInfo) (bool, error) {
	return true, nil
}
func GenerateWinningPoStSectorChallenge(abi.RegisteredPoStProof, abi.ActorID, abi.PoStRandomness, uint64) ([]uint64, error) {
	return nil, nil
}
func GeneratePoStFallbackSectorChallenges(abi.RegisteredPoStProof, abi.ActorID, abi.PoStRandomness, []abi.SectorNumber) (*FallbackChallenges, error) {
	return nil, nil
}
func GenerateSingleVanillaProof(PrivateSectorInfo, []uint64) ([]byte, error) {
	return nil, nil
}
func GetGPUDevices() ([]string, error) {
	return nil, nil
}
func FauxRep(abi.RegisteredSealProof, string, string) (cid.Cid, error) {
	return cid.Undef, nil
}

// SignatureBytes is the length of a BLS signature
const SignatureBytes = 96

// PrivateKeyBytes is the length of a BLS private key
const PrivateKeyBytes = 32

// PublicKeyBytes is the length of a BLS public key
const PublicKeyBytes = 48

// DigestBytes is the length of a BLS message hash/digest
const DigestBytes = 96

// Signature is a compressed affine
type Signature [SignatureBytes]byte

// PrivateKey is a compressed affine
type PrivateKey [PrivateKeyBytes]byte

// PublicKey is a compressed affine
type PublicKey [PublicKeyBytes]byte

// Message is a byte slice
type Message []byte

// Digest is a compressed affine
type Digest [DigestBytes]byte

// Used when generating a private key deterministically
type PrivateKeyGenSeed [32]byte

func Hash(message Message) Digest {
	return Digest{}
}

func Verify(signature *Signature, digests []Digest, publicKeys []PublicKey) bool {
	return true
}

func HashVerify(signature *Signature, messages []Message, publicKeys []PublicKey) bool {
	return true
}

func Aggregate(signatures []Signature) *Signature {
	var s Signature
	return &s
}

func PrivateKeyGenerate() PrivateKey {
	return PrivateKey{}
}

func PrivateKeyGenerateWithSeed(seed PrivateKeyGenSeed) PrivateKey {
	return PrivateKey{}
}

func PrivateKeySign(privateKey PrivateKey, message Message) *Signature {
	var s Signature
	return &s
}

func PrivateKeyPublicKey(privateKey PrivateKey) PublicKey {
	return PublicKey{}
}

func CreateZeroSignature() Signature {
	return Signature{}
}

func AggregateSealProofs(aggregateInfo proof.AggregateSealVerifyProofAndInfos, proofs [][]byte) (out []byte, err error) {
	return nil, nil
}

func VerifyAggregateSeals(aggregate proof.AggregateSealVerifyProofAndInfos) (bool, error) {
	return true, nil
}

type FunctionsSectorUpdate struct{}

var SectorUpdate = FunctionsSectorUpdate{}

func (FunctionsSectorUpdate) EncodeInto(
	proofType abi.RegisteredUpdateProof,
	newReplicaPath string,
	newReplicaCachePath string,
	sectorKeyPath string,
	sectorKeyCachePath string,
	stagedDataPath string,
	pieces []abi.PieceInfo,
) (sealedCID cid.Cid, unsealedCID cid.Cid, err error) {
	return cid.Undef, cid.Undef, nil
}

func (FunctionsSectorUpdate) DecodeFrom(
	proofType abi.RegisteredUpdateProof,
	outDataPath string,
	replicaPath string,
	sectorKeyPath string,
	sectorKeyCachePath string,
	unsealedCID cid.Cid,
) error {
	return nil
}

func (FunctionsSectorUpdate) RemoveData(
	proofType abi.RegisteredUpdateProof,
	sectorKeyPath string,
	sectorKeyCachePath string,
	replicaPath string,
	replicaCachePath string,
	dataPath string,
	unsealedCID cid.Cid,
) error {
	return nil
}

func (FunctionsSectorUpdate) GenerateUpdateVanillaProofs(
	proofType abi.RegisteredUpdateProof,
	oldSealedCID cid.Cid,
	newSealedCID cid.Cid,
	unsealedCID cid.Cid,
	newReplicaPath string,
	newReplicaCachePath string,
	sectorKeyPath string,
	sectorKeyCachePath string,
) ([][]byte, error) {
	return nil, nil
}

func (FunctionsSectorUpdate) VerifyVanillaProofs(
	proofType abi.RegisteredUpdateProof,
	oldSealedCID cid.Cid,
	newSealedCID cid.Cid,
	unsealedCID cid.Cid,
	vanillaProofs [][]byte,
) (bool, error) {
	return true, nil
}

func (FunctionsSectorUpdate) GenerateUpdateProofWithVanilla(
	proofType abi.RegisteredUpdateProof,
	oldSealedCID cid.Cid,
	newSealedCID cid.Cid,
	unsealedCID cid.Cid,
	vanillaProofs [][]byte,
) ([]byte, error) {
	return nil, nil
}

func toUpdateVanillaProofs(src [][]byte) ([]FilPartitionProof, func()) {
	return nil, nil
}

func (FunctionsSectorUpdate) GenerateUpdateProof(
	proofType abi.RegisteredUpdateProof,
	oldSealedCID cid.Cid,
	newSealedCID cid.Cid,
	unsealedCID cid.Cid,
	newReplicaPath string,
	newReplicaCachePath string,
	sectorKeyPath string,
	sectorKeyCachePath string,
) ([]byte, error) {
	return nil, nil
}

func (FunctionsSectorUpdate) VerifyUpdateProof(info proof.ReplicaUpdateInfo) (bool, error) {
	return true, nil
}
