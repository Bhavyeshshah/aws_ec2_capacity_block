// Code generated by "internal/generate/listpages/main.go -AWSSDKVersion=2 -ListOps=ListActivatedRulesInRuleGroup,ListByteMatchSets,ListIPSets,ListRateBasedRules,ListRegexMatchSets,ListRegexPatternSets,ListRules,ListRuleGroups,ListSubscribedRuleGroups -Paginator=NextMarker"; DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func listActivatedRulesInRuleGroupPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListActivatedRulesInRuleGroupInput, fn func(*wafregional.ListActivatedRulesInRuleGroupOutput, bool) bool) error {
	for {
		output, err := conn.ListActivatedRulesInRuleGroup(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listByteMatchSetsPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListByteMatchSetsInput, fn func(*wafregional.ListByteMatchSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListByteMatchSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listIPSetsPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListIPSetsInput, fn func(*wafregional.ListIPSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListIPSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listRateBasedRulesPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListRateBasedRulesInput, fn func(*wafregional.ListRateBasedRulesOutput, bool) bool) error {
	for {
		output, err := conn.ListRateBasedRules(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listRegexMatchSetsPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListRegexMatchSetsInput, fn func(*wafregional.ListRegexMatchSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListRegexMatchSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listRegexPatternSetsPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListRegexPatternSetsInput, fn func(*wafregional.ListRegexPatternSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListRegexPatternSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listRuleGroupsPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListRuleGroupsInput, fn func(*wafregional.ListRuleGroupsOutput, bool) bool) error {
	for {
		output, err := conn.ListRuleGroups(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listRulesPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListRulesInput, fn func(*wafregional.ListRulesOutput, bool) bool) error {
	for {
		output, err := conn.ListRules(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listSubscribedRuleGroupsPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListSubscribedRuleGroupsInput, fn func(*wafregional.ListSubscribedRuleGroupsOutput, bool) bool) error {
	for {
		output, err := conn.ListSubscribedRuleGroups(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
