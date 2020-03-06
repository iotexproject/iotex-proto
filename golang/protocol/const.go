package protocol

const (
	RewardingProtocolID            = "rewarding"
	ReadAvailableBalanceMethodName = "AvailableBalance"
	ReadTotalBalanceMethodName     = "TotalBalance"
	ReadUnclaimedBalanceMethodName = "UnclaimedBalance"
)

const (
	StakingProtcolID                            = "staking"
	ReadStakingVoteBucketsMethodName            = "getBuckets"
	ReadStakingVoteBucketsByVoterMethodName     = "getBucketsByVoter"
	ReadStakingVoteBucketsByCandidateMethodName = "getBucketsByCandidate"
	ReadStakingCandidatesMethodName             = "getCandidates"
	ReadStakingCandidatesByNameMethodName       = "getCandidatesByName"
)
