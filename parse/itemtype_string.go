// generated by stringer -type ItemType; DO NOT EDIT

package parse

import "fmt"

const _ItemType_name = "ItemErrorItemEOFItemEndOfLineItemSpaceItemBareItemSingleQuotedItemDoubleQuotedItemRedirLeaderItemStatusRedirLeaderItemPipeItemQuestionLParenItemLParenItemRParenItemLBracketItemRBracketItemLBraceItemRBraceItemDollarItemSemicolonItemAmpersandItemSigil"

var _ItemType_index = [...]uint8{0, 9, 16, 29, 38, 46, 62, 78, 93, 114, 122, 140, 150, 160, 172, 184, 194, 204, 214, 227, 240, 249}

func (i ItemType) String() string {
	if i < 0 || i+1 >= ItemType(len(_ItemType_index)) {
		return fmt.Sprintf("ItemType(%d)", i)
	}
	return _ItemType_name[_ItemType_index[i]:_ItemType_index[i+1]]
}
