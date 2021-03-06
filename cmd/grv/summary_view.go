package main

const (
	// SummaryViewTitle is the title of the Summary View
	SummaryViewTitle = "Summary View"
)

// NewSummaryView creates a new instance
func NewSummaryView(repoData RepoData, repoController RepoController, channels Channels, config Config, variables GRVVariableSetter) *ContainerView {
	gitSummaryView := NewGitSummaryView(repoData, repoController, channels, config, variables)

	summaryView := NewContainerView(channels, config)
	summaryView.SetTitle(SummaryViewTitle)
	summaryView.SetOrientation(CoDynamic)
	summaryView.SetViewID(ViewSummary)
	summaryView.AddChildViews(gitSummaryView)

	return summaryView
}
