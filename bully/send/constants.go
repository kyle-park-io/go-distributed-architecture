package send

type Enum int

const (
	// default
	Unknown Enum = iota

	// scenario
	FAILURE

	// task
	BASIC
	ELECTION

	// msg
	DoTask
	Done_Task
	ReportToLeader
	Done_ReportToLeader
	RespondToReport
	Done_RespondToReport
	CompleteRequestToLeader
	Done_CompleteRequestToLeader
	HoldElection
	Done_HoldElection
	RespondToElection
	Done_RespondToElection
	TransferInitiative
	Done_TransferInitiative

	// status
	SUCCESS
	FAIL

	ADMIN_NODE = 9999
)

func (e Enum) String() string {
	switch e {
	case Unknown:
		return "Unknown"
	case FAILURE:
		return "FAILURE"
	case BASIC:
		return "BASIC"
	case ELECTION:
		return "ELECTION"
	case SUCCESS:
		return "SUCCESS"
	case FAIL:
		return "FAIL"
	case DoTask:
		return "DoTask"
	case Done_Task:
		return "Done_Task"
	case ReportToLeader:
		return "ReportToLeader"
	case Done_ReportToLeader:
		return "Done_ReportToLeader"
	case RespondToReport:
		return "RespondToReport"
	case Done_RespondToReport:
		return "Done_RespondToReport"
	case CompleteRequestToLeader:
		return "CompleteRequestToLeader"
	case Done_CompleteRequestToLeader:
		return "Done_CompleteRequestToLeader"
	case HoldElection:
		return "HoldElection"
	case Done_HoldElection:
		return "Done_HoldElection"
	case RespondToElection:
		return "RespondToElection"
	case Done_RespondToElection:
		return "Done_RespondToElection"
	case TransferInitiative:
		return "TransferInitiative"
	case Done_TransferInitiative:
		return "Done_TransferInitiative"
	default:
		return "Unknown"
	}
}
