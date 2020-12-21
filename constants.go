package vcs

const (
	GitRemoteNameOrigin   = "origin"
	GitRemoteNameUpstream = "upstream"
	GitRemoteNameCrawlab  = "crawlab"
)
const GitDefaultRemoteName = GitRemoteNameOrigin

type GitAuthType int

const (
	GitAuthTypeNone GitAuthType = iota
	GitAuthTypeHTTP
	GitAuthTypeSSH
)

type GitInitType int

const (
	GitInitTypeFs GitInitType = iota
	GitInitTypeMem
)