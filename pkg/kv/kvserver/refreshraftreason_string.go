// Code generated by "stringer"; DO NOT EDIT.

package kvserver

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[noReason-0]
	_ = x[reasonNewLeader-1]
	_ = x[reasonNewLeaderOrConfigChange-2]
	_ = x[reasonSnapshotApplied-3]
	_ = x[reasonTicks-4]
}

const _refreshRaftReason_name = "noReasonreasonNewLeaderreasonNewLeaderOrConfigChangereasonSnapshotAppliedreasonTicks"

var _refreshRaftReason_index = [...]uint8{0, 8, 23, 52, 73, 84}

func (i refreshRaftReason) String() string {
	if i < 0 || i >= refreshRaftReason(len(_refreshRaftReason_index)-1) {
		return "refreshRaftReason(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _refreshRaftReason_name[_refreshRaftReason_index[i]:_refreshRaftReason_index[i+1]]
}
