// Code generated by "internal/generate/listpages/main.go -AWSSDKVersion=2 -ListOps=ListIPSets,ListRegexPatternSets,ListRuleGroups,ListWebACLs -Paginator=NextMarker"; DO NOT EDIT.

package wafv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
)

func listIPSetsPages(ctx context.Context, conn *wafv2.Client, input *wafv2.ListIPSetsInput, fn func(*wafv2.ListIPSetsOutput, bool) bool) error {
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
func listRegexPatternSetsPages(ctx context.Context, conn *wafv2.Client, input *wafv2.ListRegexPatternSetsInput, fn func(*wafv2.ListRegexPatternSetsOutput, bool) bool) error {
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
func listRuleGroupsPages(ctx context.Context, conn *wafv2.Client, input *wafv2.ListRuleGroupsInput, fn func(*wafv2.ListRuleGroupsOutput, bool) bool) error {
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
func listWebACLsPages(ctx context.Context, conn *wafv2.Client, input *wafv2.ListWebACLsInput, fn func(*wafv2.ListWebACLsOutput, bool) bool) error {
	for {
		output, err := conn.ListWebACLs(ctx, input)
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
