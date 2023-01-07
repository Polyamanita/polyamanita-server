package fakes

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBAPI struct {
	BatchExecuteStatementCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			BatchExecuteStatementInput *dynamodb.BatchExecuteStatementInput
		}
		Returns struct {
			BatchExecuteStatementOutput *dynamodb.BatchExecuteStatementOutput
			Error                       error
		}
		Stub func(*dynamodb.BatchExecuteStatementInput) (*dynamodb.BatchExecuteStatementOutput, error)
	}
	BatchExecuteStatementRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			BatchExecuteStatementInput *dynamodb.BatchExecuteStatementInput
		}
		Returns struct {
			Request                     *request.Request
			BatchExecuteStatementOutput *dynamodb.BatchExecuteStatementOutput
		}
		Stub func(*dynamodb.BatchExecuteStatementInput) (*request.Request, *dynamodb.BatchExecuteStatementOutput)
	}
	BatchExecuteStatementWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                    context.Context
			BatchExecuteStatementInput *dynamodb.BatchExecuteStatementInput
			OptionSlice                []request.Option
		}
		Returns struct {
			BatchExecuteStatementOutput *dynamodb.BatchExecuteStatementOutput
			Error                       error
		}
		Stub func(context.Context, *dynamodb.BatchExecuteStatementInput, ...request.Option) (*dynamodb.BatchExecuteStatementOutput, error)
	}
	BatchGetItemCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			BatchGetItemInput *dynamodb.BatchGetItemInput
		}
		Returns struct {
			BatchGetItemOutput *dynamodb.BatchGetItemOutput
			Error              error
		}
		Stub func(*dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error)
	}
	BatchGetItemPagesCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			BatchGetItemInput                      *dynamodb.BatchGetItemInput
			FuncDynamodbBatchGetItemOutputBoolBool func(*dynamodb.BatchGetItemOutput, bool) bool
		}
		Returns struct {
			Error error
		}
		Stub func(*dynamodb.BatchGetItemInput, func(*dynamodb.BatchGetItemOutput, bool) bool) error
	}
	BatchGetItemPagesWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                                context.Context
			BatchGetItemInput                      *dynamodb.BatchGetItemInput
			FuncDynamodbBatchGetItemOutputBoolBool func(*dynamodb.BatchGetItemOutput, bool) bool
			OptionSlice                            []request.Option
		}
		Returns struct {
			Error error
		}
		Stub func(context.Context, *dynamodb.BatchGetItemInput, func(*dynamodb.BatchGetItemOutput, bool) bool, ...request.Option) error
	}
	BatchGetItemRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			BatchGetItemInput *dynamodb.BatchGetItemInput
		}
		Returns struct {
			Request            *request.Request
			BatchGetItemOutput *dynamodb.BatchGetItemOutput
		}
		Stub func(*dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput)
	}
	BatchGetItemWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context           context.Context
			BatchGetItemInput *dynamodb.BatchGetItemInput
			OptionSlice       []request.Option
		}
		Returns struct {
			BatchGetItemOutput *dynamodb.BatchGetItemOutput
			Error              error
		}
		Stub func(context.Context, *dynamodb.BatchGetItemInput, ...request.Option) (*dynamodb.BatchGetItemOutput, error)
	}
	BatchWriteItemCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			BatchWriteItemInput *dynamodb.BatchWriteItemInput
		}
		Returns struct {
			BatchWriteItemOutput *dynamodb.BatchWriteItemOutput
			Error                error
		}
		Stub func(*dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error)
	}
	BatchWriteItemRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			BatchWriteItemInput *dynamodb.BatchWriteItemInput
		}
		Returns struct {
			Request              *request.Request
			BatchWriteItemOutput *dynamodb.BatchWriteItemOutput
		}
		Stub func(*dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput)
	}
	BatchWriteItemWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context             context.Context
			BatchWriteItemInput *dynamodb.BatchWriteItemInput
			OptionSlice         []request.Option
		}
		Returns struct {
			BatchWriteItemOutput *dynamodb.BatchWriteItemOutput
			Error                error
		}
		Stub func(context.Context, *dynamodb.BatchWriteItemInput, ...request.Option) (*dynamodb.BatchWriteItemOutput, error)
	}
	CreateBackupCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			CreateBackupInput *dynamodb.CreateBackupInput
		}
		Returns struct {
			CreateBackupOutput *dynamodb.CreateBackupOutput
			Error              error
		}
		Stub func(*dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error)
	}
	CreateBackupRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			CreateBackupInput *dynamodb.CreateBackupInput
		}
		Returns struct {
			Request            *request.Request
			CreateBackupOutput *dynamodb.CreateBackupOutput
		}
		Stub func(*dynamodb.CreateBackupInput) (*request.Request, *dynamodb.CreateBackupOutput)
	}
	CreateBackupWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context           context.Context
			CreateBackupInput *dynamodb.CreateBackupInput
			OptionSlice       []request.Option
		}
		Returns struct {
			CreateBackupOutput *dynamodb.CreateBackupOutput
			Error              error
		}
		Stub func(context.Context, *dynamodb.CreateBackupInput, ...request.Option) (*dynamodb.CreateBackupOutput, error)
	}
	CreateGlobalTableCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			CreateGlobalTableInput *dynamodb.CreateGlobalTableInput
		}
		Returns struct {
			CreateGlobalTableOutput *dynamodb.CreateGlobalTableOutput
			Error                   error
		}
		Stub func(*dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error)
	}
	CreateGlobalTableRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			CreateGlobalTableInput *dynamodb.CreateGlobalTableInput
		}
		Returns struct {
			Request                 *request.Request
			CreateGlobalTableOutput *dynamodb.CreateGlobalTableOutput
		}
		Stub func(*dynamodb.CreateGlobalTableInput) (*request.Request, *dynamodb.CreateGlobalTableOutput)
	}
	CreateGlobalTableWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                context.Context
			CreateGlobalTableInput *dynamodb.CreateGlobalTableInput
			OptionSlice            []request.Option
		}
		Returns struct {
			CreateGlobalTableOutput *dynamodb.CreateGlobalTableOutput
			Error                   error
		}
		Stub func(context.Context, *dynamodb.CreateGlobalTableInput, ...request.Option) (*dynamodb.CreateGlobalTableOutput, error)
	}
	CreateTableCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			CreateTableInput *dynamodb.CreateTableInput
		}
		Returns struct {
			CreateTableOutput *dynamodb.CreateTableOutput
			Error             error
		}
		Stub func(*dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error)
	}
	CreateTableRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			CreateTableInput *dynamodb.CreateTableInput
		}
		Returns struct {
			Request           *request.Request
			CreateTableOutput *dynamodb.CreateTableOutput
		}
		Stub func(*dynamodb.CreateTableInput) (*request.Request, *dynamodb.CreateTableOutput)
	}
	CreateTableWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context          context.Context
			CreateTableInput *dynamodb.CreateTableInput
			OptionSlice      []request.Option
		}
		Returns struct {
			CreateTableOutput *dynamodb.CreateTableOutput
			Error             error
		}
		Stub func(context.Context, *dynamodb.CreateTableInput, ...request.Option) (*dynamodb.CreateTableOutput, error)
	}
	DeleteBackupCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBackupInput *dynamodb.DeleteBackupInput
		}
		Returns struct {
			DeleteBackupOutput *dynamodb.DeleteBackupOutput
			Error              error
		}
		Stub func(*dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error)
	}
	DeleteBackupRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBackupInput *dynamodb.DeleteBackupInput
		}
		Returns struct {
			Request            *request.Request
			DeleteBackupOutput *dynamodb.DeleteBackupOutput
		}
		Stub func(*dynamodb.DeleteBackupInput) (*request.Request, *dynamodb.DeleteBackupOutput)
	}
	DeleteBackupWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context           context.Context
			DeleteBackupInput *dynamodb.DeleteBackupInput
			OptionSlice       []request.Option
		}
		Returns struct {
			DeleteBackupOutput *dynamodb.DeleteBackupOutput
			Error              error
		}
		Stub func(context.Context, *dynamodb.DeleteBackupInput, ...request.Option) (*dynamodb.DeleteBackupOutput, error)
	}
	DeleteItemCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteItemInput *dynamodb.DeleteItemInput
		}
		Returns struct {
			DeleteItemOutput *dynamodb.DeleteItemOutput
			Error            error
		}
		Stub func(*dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error)
	}
	DeleteItemRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteItemInput *dynamodb.DeleteItemInput
		}
		Returns struct {
			Request          *request.Request
			DeleteItemOutput *dynamodb.DeleteItemOutput
		}
		Stub func(*dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput)
	}
	DeleteItemWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context         context.Context
			DeleteItemInput *dynamodb.DeleteItemInput
			OptionSlice     []request.Option
		}
		Returns struct {
			DeleteItemOutput *dynamodb.DeleteItemOutput
			Error            error
		}
		Stub func(context.Context, *dynamodb.DeleteItemInput, ...request.Option) (*dynamodb.DeleteItemOutput, error)
	}
	DeleteTableCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteTableInput *dynamodb.DeleteTableInput
		}
		Returns struct {
			DeleteTableOutput *dynamodb.DeleteTableOutput
			Error             error
		}
		Stub func(*dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error)
	}
	DeleteTableRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteTableInput *dynamodb.DeleteTableInput
		}
		Returns struct {
			Request           *request.Request
			DeleteTableOutput *dynamodb.DeleteTableOutput
		}
		Stub func(*dynamodb.DeleteTableInput) (*request.Request, *dynamodb.DeleteTableOutput)
	}
	DeleteTableWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context          context.Context
			DeleteTableInput *dynamodb.DeleteTableInput
			OptionSlice      []request.Option
		}
		Returns struct {
			DeleteTableOutput *dynamodb.DeleteTableOutput
			Error             error
		}
		Stub func(context.Context, *dynamodb.DeleteTableInput, ...request.Option) (*dynamodb.DeleteTableOutput, error)
	}
	DescribeBackupCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeBackupInput *dynamodb.DescribeBackupInput
		}
		Returns struct {
			DescribeBackupOutput *dynamodb.DescribeBackupOutput
			Error                error
		}
		Stub func(*dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error)
	}
	DescribeBackupRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeBackupInput *dynamodb.DescribeBackupInput
		}
		Returns struct {
			Request              *request.Request
			DescribeBackupOutput *dynamodb.DescribeBackupOutput
		}
		Stub func(*dynamodb.DescribeBackupInput) (*request.Request, *dynamodb.DescribeBackupOutput)
	}
	DescribeBackupWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context             context.Context
			DescribeBackupInput *dynamodb.DescribeBackupInput
			OptionSlice         []request.Option
		}
		Returns struct {
			DescribeBackupOutput *dynamodb.DescribeBackupOutput
			Error                error
		}
		Stub func(context.Context, *dynamodb.DescribeBackupInput, ...request.Option) (*dynamodb.DescribeBackupOutput, error)
	}
	DescribeContinuousBackupsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeContinuousBackupsInput *dynamodb.DescribeContinuousBackupsInput
		}
		Returns struct {
			DescribeContinuousBackupsOutput *dynamodb.DescribeContinuousBackupsOutput
			Error                           error
		}
		Stub func(*dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error)
	}
	DescribeContinuousBackupsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeContinuousBackupsInput *dynamodb.DescribeContinuousBackupsInput
		}
		Returns struct {
			Request                         *request.Request
			DescribeContinuousBackupsOutput *dynamodb.DescribeContinuousBackupsOutput
		}
		Stub func(*dynamodb.DescribeContinuousBackupsInput) (*request.Request, *dynamodb.DescribeContinuousBackupsOutput)
	}
	DescribeContinuousBackupsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                        context.Context
			DescribeContinuousBackupsInput *dynamodb.DescribeContinuousBackupsInput
			OptionSlice                    []request.Option
		}
		Returns struct {
			DescribeContinuousBackupsOutput *dynamodb.DescribeContinuousBackupsOutput
			Error                           error
		}
		Stub func(context.Context, *dynamodb.DescribeContinuousBackupsInput, ...request.Option) (*dynamodb.DescribeContinuousBackupsOutput, error)
	}
	DescribeContributorInsightsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeContributorInsightsInput *dynamodb.DescribeContributorInsightsInput
		}
		Returns struct {
			DescribeContributorInsightsOutput *dynamodb.DescribeContributorInsightsOutput
			Error                             error
		}
		Stub func(*dynamodb.DescribeContributorInsightsInput) (*dynamodb.DescribeContributorInsightsOutput, error)
	}
	DescribeContributorInsightsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeContributorInsightsInput *dynamodb.DescribeContributorInsightsInput
		}
		Returns struct {
			Request                           *request.Request
			DescribeContributorInsightsOutput *dynamodb.DescribeContributorInsightsOutput
		}
		Stub func(*dynamodb.DescribeContributorInsightsInput) (*request.Request, *dynamodb.DescribeContributorInsightsOutput)
	}
	DescribeContributorInsightsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                          context.Context
			DescribeContributorInsightsInput *dynamodb.DescribeContributorInsightsInput
			OptionSlice                      []request.Option
		}
		Returns struct {
			DescribeContributorInsightsOutput *dynamodb.DescribeContributorInsightsOutput
			Error                             error
		}
		Stub func(context.Context, *dynamodb.DescribeContributorInsightsInput, ...request.Option) (*dynamodb.DescribeContributorInsightsOutput, error)
	}
	DescribeEndpointsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeEndpointsInput *dynamodb.DescribeEndpointsInput
		}
		Returns struct {
			DescribeEndpointsOutput *dynamodb.DescribeEndpointsOutput
			Error                   error
		}
		Stub func(*dynamodb.DescribeEndpointsInput) (*dynamodb.DescribeEndpointsOutput, error)
	}
	DescribeEndpointsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeEndpointsInput *dynamodb.DescribeEndpointsInput
		}
		Returns struct {
			Request                 *request.Request
			DescribeEndpointsOutput *dynamodb.DescribeEndpointsOutput
		}
		Stub func(*dynamodb.DescribeEndpointsInput) (*request.Request, *dynamodb.DescribeEndpointsOutput)
	}
	DescribeEndpointsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                context.Context
			DescribeEndpointsInput *dynamodb.DescribeEndpointsInput
			OptionSlice            []request.Option
		}
		Returns struct {
			DescribeEndpointsOutput *dynamodb.DescribeEndpointsOutput
			Error                   error
		}
		Stub func(context.Context, *dynamodb.DescribeEndpointsInput, ...request.Option) (*dynamodb.DescribeEndpointsOutput, error)
	}
	DescribeExportCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeExportInput *dynamodb.DescribeExportInput
		}
		Returns struct {
			DescribeExportOutput *dynamodb.DescribeExportOutput
			Error                error
		}
		Stub func(*dynamodb.DescribeExportInput) (*dynamodb.DescribeExportOutput, error)
	}
	DescribeExportRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeExportInput *dynamodb.DescribeExportInput
		}
		Returns struct {
			Request              *request.Request
			DescribeExportOutput *dynamodb.DescribeExportOutput
		}
		Stub func(*dynamodb.DescribeExportInput) (*request.Request, *dynamodb.DescribeExportOutput)
	}
	DescribeExportWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context             context.Context
			DescribeExportInput *dynamodb.DescribeExportInput
			OptionSlice         []request.Option
		}
		Returns struct {
			DescribeExportOutput *dynamodb.DescribeExportOutput
			Error                error
		}
		Stub func(context.Context, *dynamodb.DescribeExportInput, ...request.Option) (*dynamodb.DescribeExportOutput, error)
	}
	DescribeGlobalTableCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeGlobalTableInput *dynamodb.DescribeGlobalTableInput
		}
		Returns struct {
			DescribeGlobalTableOutput *dynamodb.DescribeGlobalTableOutput
			Error                     error
		}
		Stub func(*dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error)
	}
	DescribeGlobalTableRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeGlobalTableInput *dynamodb.DescribeGlobalTableInput
		}
		Returns struct {
			Request                   *request.Request
			DescribeGlobalTableOutput *dynamodb.DescribeGlobalTableOutput
		}
		Stub func(*dynamodb.DescribeGlobalTableInput) (*request.Request, *dynamodb.DescribeGlobalTableOutput)
	}
	DescribeGlobalTableSettingsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeGlobalTableSettingsInput *dynamodb.DescribeGlobalTableSettingsInput
		}
		Returns struct {
			DescribeGlobalTableSettingsOutput *dynamodb.DescribeGlobalTableSettingsOutput
			Error                             error
		}
		Stub func(*dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error)
	}
	DescribeGlobalTableSettingsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeGlobalTableSettingsInput *dynamodb.DescribeGlobalTableSettingsInput
		}
		Returns struct {
			Request                           *request.Request
			DescribeGlobalTableSettingsOutput *dynamodb.DescribeGlobalTableSettingsOutput
		}
		Stub func(*dynamodb.DescribeGlobalTableSettingsInput) (*request.Request, *dynamodb.DescribeGlobalTableSettingsOutput)
	}
	DescribeGlobalTableSettingsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                          context.Context
			DescribeGlobalTableSettingsInput *dynamodb.DescribeGlobalTableSettingsInput
			OptionSlice                      []request.Option
		}
		Returns struct {
			DescribeGlobalTableSettingsOutput *dynamodb.DescribeGlobalTableSettingsOutput
			Error                             error
		}
		Stub func(context.Context, *dynamodb.DescribeGlobalTableSettingsInput, ...request.Option) (*dynamodb.DescribeGlobalTableSettingsOutput, error)
	}
	DescribeGlobalTableWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                  context.Context
			DescribeGlobalTableInput *dynamodb.DescribeGlobalTableInput
			OptionSlice              []request.Option
		}
		Returns struct {
			DescribeGlobalTableOutput *dynamodb.DescribeGlobalTableOutput
			Error                     error
		}
		Stub func(context.Context, *dynamodb.DescribeGlobalTableInput, ...request.Option) (*dynamodb.DescribeGlobalTableOutput, error)
	}
	DescribeImportCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeImportInput *dynamodb.DescribeImportInput
		}
		Returns struct {
			DescribeImportOutput *dynamodb.DescribeImportOutput
			Error                error
		}
		Stub func(*dynamodb.DescribeImportInput) (*dynamodb.DescribeImportOutput, error)
	}
	DescribeImportRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeImportInput *dynamodb.DescribeImportInput
		}
		Returns struct {
			Request              *request.Request
			DescribeImportOutput *dynamodb.DescribeImportOutput
		}
		Stub func(*dynamodb.DescribeImportInput) (*request.Request, *dynamodb.DescribeImportOutput)
	}
	DescribeImportWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context             context.Context
			DescribeImportInput *dynamodb.DescribeImportInput
			OptionSlice         []request.Option
		}
		Returns struct {
			DescribeImportOutput *dynamodb.DescribeImportOutput
			Error                error
		}
		Stub func(context.Context, *dynamodb.DescribeImportInput, ...request.Option) (*dynamodb.DescribeImportOutput, error)
	}
	DescribeKinesisStreamingDestinationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeKinesisStreamingDestinationInput *dynamodb.DescribeKinesisStreamingDestinationInput
		}
		Returns struct {
			DescribeKinesisStreamingDestinationOutput *dynamodb.DescribeKinesisStreamingDestinationOutput
			Error                                     error
		}
		Stub func(*dynamodb.DescribeKinesisStreamingDestinationInput) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error)
	}
	DescribeKinesisStreamingDestinationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeKinesisStreamingDestinationInput *dynamodb.DescribeKinesisStreamingDestinationInput
		}
		Returns struct {
			Request                                   *request.Request
			DescribeKinesisStreamingDestinationOutput *dynamodb.DescribeKinesisStreamingDestinationOutput
		}
		Stub func(*dynamodb.DescribeKinesisStreamingDestinationInput) (*request.Request, *dynamodb.DescribeKinesisStreamingDestinationOutput)
	}
	DescribeKinesisStreamingDestinationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                                  context.Context
			DescribeKinesisStreamingDestinationInput *dynamodb.DescribeKinesisStreamingDestinationInput
			OptionSlice                              []request.Option
		}
		Returns struct {
			DescribeKinesisStreamingDestinationOutput *dynamodb.DescribeKinesisStreamingDestinationOutput
			Error                                     error
		}
		Stub func(context.Context, *dynamodb.DescribeKinesisStreamingDestinationInput, ...request.Option) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error)
	}
	DescribeLimitsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeLimitsInput *dynamodb.DescribeLimitsInput
		}
		Returns struct {
			DescribeLimitsOutput *dynamodb.DescribeLimitsOutput
			Error                error
		}
		Stub func(*dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error)
	}
	DescribeLimitsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeLimitsInput *dynamodb.DescribeLimitsInput
		}
		Returns struct {
			Request              *request.Request
			DescribeLimitsOutput *dynamodb.DescribeLimitsOutput
		}
		Stub func(*dynamodb.DescribeLimitsInput) (*request.Request, *dynamodb.DescribeLimitsOutput)
	}
	DescribeLimitsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context             context.Context
			DescribeLimitsInput *dynamodb.DescribeLimitsInput
			OptionSlice         []request.Option
		}
		Returns struct {
			DescribeLimitsOutput *dynamodb.DescribeLimitsOutput
			Error                error
		}
		Stub func(context.Context, *dynamodb.DescribeLimitsInput, ...request.Option) (*dynamodb.DescribeLimitsOutput, error)
	}
	DescribeTableCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeTableInput *dynamodb.DescribeTableInput
		}
		Returns struct {
			DescribeTableOutput *dynamodb.DescribeTableOutput
			Error               error
		}
		Stub func(*dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error)
	}
	DescribeTableReplicaAutoScalingCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeTableReplicaAutoScalingInput *dynamodb.DescribeTableReplicaAutoScalingInput
		}
		Returns struct {
			DescribeTableReplicaAutoScalingOutput *dynamodb.DescribeTableReplicaAutoScalingOutput
			Error                                 error
		}
		Stub func(*dynamodb.DescribeTableReplicaAutoScalingInput) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error)
	}
	DescribeTableReplicaAutoScalingRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeTableReplicaAutoScalingInput *dynamodb.DescribeTableReplicaAutoScalingInput
		}
		Returns struct {
			Request                               *request.Request
			DescribeTableReplicaAutoScalingOutput *dynamodb.DescribeTableReplicaAutoScalingOutput
		}
		Stub func(*dynamodb.DescribeTableReplicaAutoScalingInput) (*request.Request, *dynamodb.DescribeTableReplicaAutoScalingOutput)
	}
	DescribeTableReplicaAutoScalingWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                              context.Context
			DescribeTableReplicaAutoScalingInput *dynamodb.DescribeTableReplicaAutoScalingInput
			OptionSlice                          []request.Option
		}
		Returns struct {
			DescribeTableReplicaAutoScalingOutput *dynamodb.DescribeTableReplicaAutoScalingOutput
			Error                                 error
		}
		Stub func(context.Context, *dynamodb.DescribeTableReplicaAutoScalingInput, ...request.Option) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error)
	}
	DescribeTableRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeTableInput *dynamodb.DescribeTableInput
		}
		Returns struct {
			Request             *request.Request
			DescribeTableOutput *dynamodb.DescribeTableOutput
		}
		Stub func(*dynamodb.DescribeTableInput) (*request.Request, *dynamodb.DescribeTableOutput)
	}
	DescribeTableWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context            context.Context
			DescribeTableInput *dynamodb.DescribeTableInput
			OptionSlice        []request.Option
		}
		Returns struct {
			DescribeTableOutput *dynamodb.DescribeTableOutput
			Error               error
		}
		Stub func(context.Context, *dynamodb.DescribeTableInput, ...request.Option) (*dynamodb.DescribeTableOutput, error)
	}
	DescribeTimeToLiveCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeTimeToLiveInput *dynamodb.DescribeTimeToLiveInput
		}
		Returns struct {
			DescribeTimeToLiveOutput *dynamodb.DescribeTimeToLiveOutput
			Error                    error
		}
		Stub func(*dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error)
	}
	DescribeTimeToLiveRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeTimeToLiveInput *dynamodb.DescribeTimeToLiveInput
		}
		Returns struct {
			Request                  *request.Request
			DescribeTimeToLiveOutput *dynamodb.DescribeTimeToLiveOutput
		}
		Stub func(*dynamodb.DescribeTimeToLiveInput) (*request.Request, *dynamodb.DescribeTimeToLiveOutput)
	}
	DescribeTimeToLiveWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                 context.Context
			DescribeTimeToLiveInput *dynamodb.DescribeTimeToLiveInput
			OptionSlice             []request.Option
		}
		Returns struct {
			DescribeTimeToLiveOutput *dynamodb.DescribeTimeToLiveOutput
			Error                    error
		}
		Stub func(context.Context, *dynamodb.DescribeTimeToLiveInput, ...request.Option) (*dynamodb.DescribeTimeToLiveOutput, error)
	}
	DisableKinesisStreamingDestinationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DisableKinesisStreamingDestinationInput *dynamodb.DisableKinesisStreamingDestinationInput
		}
		Returns struct {
			DisableKinesisStreamingDestinationOutput *dynamodb.DisableKinesisStreamingDestinationOutput
			Error                                    error
		}
		Stub func(*dynamodb.DisableKinesisStreamingDestinationInput) (*dynamodb.DisableKinesisStreamingDestinationOutput, error)
	}
	DisableKinesisStreamingDestinationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DisableKinesisStreamingDestinationInput *dynamodb.DisableKinesisStreamingDestinationInput
		}
		Returns struct {
			Request                                  *request.Request
			DisableKinesisStreamingDestinationOutput *dynamodb.DisableKinesisStreamingDestinationOutput
		}
		Stub func(*dynamodb.DisableKinesisStreamingDestinationInput) (*request.Request, *dynamodb.DisableKinesisStreamingDestinationOutput)
	}
	DisableKinesisStreamingDestinationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                                 context.Context
			DisableKinesisStreamingDestinationInput *dynamodb.DisableKinesisStreamingDestinationInput
			OptionSlice                             []request.Option
		}
		Returns struct {
			DisableKinesisStreamingDestinationOutput *dynamodb.DisableKinesisStreamingDestinationOutput
			Error                                    error
		}
		Stub func(context.Context, *dynamodb.DisableKinesisStreamingDestinationInput, ...request.Option) (*dynamodb.DisableKinesisStreamingDestinationOutput, error)
	}
	EnableKinesisStreamingDestinationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			EnableKinesisStreamingDestinationInput *dynamodb.EnableKinesisStreamingDestinationInput
		}
		Returns struct {
			EnableKinesisStreamingDestinationOutput *dynamodb.EnableKinesisStreamingDestinationOutput
			Error                                   error
		}
		Stub func(*dynamodb.EnableKinesisStreamingDestinationInput) (*dynamodb.EnableKinesisStreamingDestinationOutput, error)
	}
	EnableKinesisStreamingDestinationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			EnableKinesisStreamingDestinationInput *dynamodb.EnableKinesisStreamingDestinationInput
		}
		Returns struct {
			Request                                 *request.Request
			EnableKinesisStreamingDestinationOutput *dynamodb.EnableKinesisStreamingDestinationOutput
		}
		Stub func(*dynamodb.EnableKinesisStreamingDestinationInput) (*request.Request, *dynamodb.EnableKinesisStreamingDestinationOutput)
	}
	EnableKinesisStreamingDestinationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                                context.Context
			EnableKinesisStreamingDestinationInput *dynamodb.EnableKinesisStreamingDestinationInput
			OptionSlice                            []request.Option
		}
		Returns struct {
			EnableKinesisStreamingDestinationOutput *dynamodb.EnableKinesisStreamingDestinationOutput
			Error                                   error
		}
		Stub func(context.Context, *dynamodb.EnableKinesisStreamingDestinationInput, ...request.Option) (*dynamodb.EnableKinesisStreamingDestinationOutput, error)
	}
	ExecuteStatementCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ExecuteStatementInput *dynamodb.ExecuteStatementInput
		}
		Returns struct {
			ExecuteStatementOutput *dynamodb.ExecuteStatementOutput
			Error                  error
		}
		Stub func(*dynamodb.ExecuteStatementInput) (*dynamodb.ExecuteStatementOutput, error)
	}
	ExecuteStatementRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ExecuteStatementInput *dynamodb.ExecuteStatementInput
		}
		Returns struct {
			Request                *request.Request
			ExecuteStatementOutput *dynamodb.ExecuteStatementOutput
		}
		Stub func(*dynamodb.ExecuteStatementInput) (*request.Request, *dynamodb.ExecuteStatementOutput)
	}
	ExecuteStatementWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context               context.Context
			ExecuteStatementInput *dynamodb.ExecuteStatementInput
			OptionSlice           []request.Option
		}
		Returns struct {
			ExecuteStatementOutput *dynamodb.ExecuteStatementOutput
			Error                  error
		}
		Stub func(context.Context, *dynamodb.ExecuteStatementInput, ...request.Option) (*dynamodb.ExecuteStatementOutput, error)
	}
	ExecuteTransactionCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ExecuteTransactionInput *dynamodb.ExecuteTransactionInput
		}
		Returns struct {
			ExecuteTransactionOutput *dynamodb.ExecuteTransactionOutput
			Error                    error
		}
		Stub func(*dynamodb.ExecuteTransactionInput) (*dynamodb.ExecuteTransactionOutput, error)
	}
	ExecuteTransactionRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ExecuteTransactionInput *dynamodb.ExecuteTransactionInput
		}
		Returns struct {
			Request                  *request.Request
			ExecuteTransactionOutput *dynamodb.ExecuteTransactionOutput
		}
		Stub func(*dynamodb.ExecuteTransactionInput) (*request.Request, *dynamodb.ExecuteTransactionOutput)
	}
	ExecuteTransactionWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                 context.Context
			ExecuteTransactionInput *dynamodb.ExecuteTransactionInput
			OptionSlice             []request.Option
		}
		Returns struct {
			ExecuteTransactionOutput *dynamodb.ExecuteTransactionOutput
			Error                    error
		}
		Stub func(context.Context, *dynamodb.ExecuteTransactionInput, ...request.Option) (*dynamodb.ExecuteTransactionOutput, error)
	}
	ExportTableToPointInTimeCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ExportTableToPointInTimeInput *dynamodb.ExportTableToPointInTimeInput
		}
		Returns struct {
			ExportTableToPointInTimeOutput *dynamodb.ExportTableToPointInTimeOutput
			Error                          error
		}
		Stub func(*dynamodb.ExportTableToPointInTimeInput) (*dynamodb.ExportTableToPointInTimeOutput, error)
	}
	ExportTableToPointInTimeRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ExportTableToPointInTimeInput *dynamodb.ExportTableToPointInTimeInput
		}
		Returns struct {
			Request                        *request.Request
			ExportTableToPointInTimeOutput *dynamodb.ExportTableToPointInTimeOutput
		}
		Stub func(*dynamodb.ExportTableToPointInTimeInput) (*request.Request, *dynamodb.ExportTableToPointInTimeOutput)
	}
	ExportTableToPointInTimeWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                       context.Context
			ExportTableToPointInTimeInput *dynamodb.ExportTableToPointInTimeInput
			OptionSlice                   []request.Option
		}
		Returns struct {
			ExportTableToPointInTimeOutput *dynamodb.ExportTableToPointInTimeOutput
			Error                          error
		}
		Stub func(context.Context, *dynamodb.ExportTableToPointInTimeInput, ...request.Option) (*dynamodb.ExportTableToPointInTimeOutput, error)
	}
	GetItemCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetItemInput *dynamodb.GetItemInput
		}
		Returns struct {
			GetItemOutput *dynamodb.GetItemOutput
			Error         error
		}
		Stub func(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
	}
	GetItemRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetItemInput *dynamodb.GetItemInput
		}
		Returns struct {
			Request       *request.Request
			GetItemOutput *dynamodb.GetItemOutput
		}
		Stub func(*dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput)
	}
	GetItemWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context      context.Context
			GetItemInput *dynamodb.GetItemInput
			OptionSlice  []request.Option
		}
		Returns struct {
			GetItemOutput *dynamodb.GetItemOutput
			Error         error
		}
		Stub func(context.Context, *dynamodb.GetItemInput, ...request.Option) (*dynamodb.GetItemOutput, error)
	}
	ImportTableCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ImportTableInput *dynamodb.ImportTableInput
		}
		Returns struct {
			ImportTableOutput *dynamodb.ImportTableOutput
			Error             error
		}
		Stub func(*dynamodb.ImportTableInput) (*dynamodb.ImportTableOutput, error)
	}
	ImportTableRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ImportTableInput *dynamodb.ImportTableInput
		}
		Returns struct {
			Request           *request.Request
			ImportTableOutput *dynamodb.ImportTableOutput
		}
		Stub func(*dynamodb.ImportTableInput) (*request.Request, *dynamodb.ImportTableOutput)
	}
	ImportTableWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context          context.Context
			ImportTableInput *dynamodb.ImportTableInput
			OptionSlice      []request.Option
		}
		Returns struct {
			ImportTableOutput *dynamodb.ImportTableOutput
			Error             error
		}
		Stub func(context.Context, *dynamodb.ImportTableInput, ...request.Option) (*dynamodb.ImportTableOutput, error)
	}
	ListBackupsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListBackupsInput *dynamodb.ListBackupsInput
		}
		Returns struct {
			ListBackupsOutput *dynamodb.ListBackupsOutput
			Error             error
		}
		Stub func(*dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error)
	}
	ListBackupsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListBackupsInput *dynamodb.ListBackupsInput
		}
		Returns struct {
			Request           *request.Request
			ListBackupsOutput *dynamodb.ListBackupsOutput
		}
		Stub func(*dynamodb.ListBackupsInput) (*request.Request, *dynamodb.ListBackupsOutput)
	}
	ListBackupsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context          context.Context
			ListBackupsInput *dynamodb.ListBackupsInput
			OptionSlice      []request.Option
		}
		Returns struct {
			ListBackupsOutput *dynamodb.ListBackupsOutput
			Error             error
		}
		Stub func(context.Context, *dynamodb.ListBackupsInput, ...request.Option) (*dynamodb.ListBackupsOutput, error)
	}
	ListContributorInsightsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListContributorInsightsInput *dynamodb.ListContributorInsightsInput
		}
		Returns struct {
			ListContributorInsightsOutput *dynamodb.ListContributorInsightsOutput
			Error                         error
		}
		Stub func(*dynamodb.ListContributorInsightsInput) (*dynamodb.ListContributorInsightsOutput, error)
	}
	ListContributorInsightsPagesCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListContributorInsightsInput                      *dynamodb.ListContributorInsightsInput
			FuncDynamodbListContributorInsightsOutputBoolBool func(*dynamodb.ListContributorInsightsOutput, bool) bool
		}
		Returns struct {
			Error error
		}
		Stub func(*dynamodb.ListContributorInsightsInput, func(*dynamodb.ListContributorInsightsOutput, bool) bool) error
	}
	ListContributorInsightsPagesWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                                           context.Context
			ListContributorInsightsInput                      *dynamodb.ListContributorInsightsInput
			FuncDynamodbListContributorInsightsOutputBoolBool func(*dynamodb.ListContributorInsightsOutput, bool) bool
			OptionSlice                                       []request.Option
		}
		Returns struct {
			Error error
		}
		Stub func(context.Context, *dynamodb.ListContributorInsightsInput, func(*dynamodb.ListContributorInsightsOutput, bool) bool, ...request.Option) error
	}
	ListContributorInsightsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListContributorInsightsInput *dynamodb.ListContributorInsightsInput
		}
		Returns struct {
			Request                       *request.Request
			ListContributorInsightsOutput *dynamodb.ListContributorInsightsOutput
		}
		Stub func(*dynamodb.ListContributorInsightsInput) (*request.Request, *dynamodb.ListContributorInsightsOutput)
	}
	ListContributorInsightsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                      context.Context
			ListContributorInsightsInput *dynamodb.ListContributorInsightsInput
			OptionSlice                  []request.Option
		}
		Returns struct {
			ListContributorInsightsOutput *dynamodb.ListContributorInsightsOutput
			Error                         error
		}
		Stub func(context.Context, *dynamodb.ListContributorInsightsInput, ...request.Option) (*dynamodb.ListContributorInsightsOutput, error)
	}
	ListExportsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListExportsInput *dynamodb.ListExportsInput
		}
		Returns struct {
			ListExportsOutput *dynamodb.ListExportsOutput
			Error             error
		}
		Stub func(*dynamodb.ListExportsInput) (*dynamodb.ListExportsOutput, error)
	}
	ListExportsPagesCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListExportsInput                      *dynamodb.ListExportsInput
			FuncDynamodbListExportsOutputBoolBool func(*dynamodb.ListExportsOutput, bool) bool
		}
		Returns struct {
			Error error
		}
		Stub func(*dynamodb.ListExportsInput, func(*dynamodb.ListExportsOutput, bool) bool) error
	}
	ListExportsPagesWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                               context.Context
			ListExportsInput                      *dynamodb.ListExportsInput
			FuncDynamodbListExportsOutputBoolBool func(*dynamodb.ListExportsOutput, bool) bool
			OptionSlice                           []request.Option
		}
		Returns struct {
			Error error
		}
		Stub func(context.Context, *dynamodb.ListExportsInput, func(*dynamodb.ListExportsOutput, bool) bool, ...request.Option) error
	}
	ListExportsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListExportsInput *dynamodb.ListExportsInput
		}
		Returns struct {
			Request           *request.Request
			ListExportsOutput *dynamodb.ListExportsOutput
		}
		Stub func(*dynamodb.ListExportsInput) (*request.Request, *dynamodb.ListExportsOutput)
	}
	ListExportsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context          context.Context
			ListExportsInput *dynamodb.ListExportsInput
			OptionSlice      []request.Option
		}
		Returns struct {
			ListExportsOutput *dynamodb.ListExportsOutput
			Error             error
		}
		Stub func(context.Context, *dynamodb.ListExportsInput, ...request.Option) (*dynamodb.ListExportsOutput, error)
	}
	ListGlobalTablesCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListGlobalTablesInput *dynamodb.ListGlobalTablesInput
		}
		Returns struct {
			ListGlobalTablesOutput *dynamodb.ListGlobalTablesOutput
			Error                  error
		}
		Stub func(*dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error)
	}
	ListGlobalTablesRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListGlobalTablesInput *dynamodb.ListGlobalTablesInput
		}
		Returns struct {
			Request                *request.Request
			ListGlobalTablesOutput *dynamodb.ListGlobalTablesOutput
		}
		Stub func(*dynamodb.ListGlobalTablesInput) (*request.Request, *dynamodb.ListGlobalTablesOutput)
	}
	ListGlobalTablesWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context               context.Context
			ListGlobalTablesInput *dynamodb.ListGlobalTablesInput
			OptionSlice           []request.Option
		}
		Returns struct {
			ListGlobalTablesOutput *dynamodb.ListGlobalTablesOutput
			Error                  error
		}
		Stub func(context.Context, *dynamodb.ListGlobalTablesInput, ...request.Option) (*dynamodb.ListGlobalTablesOutput, error)
	}
	ListImportsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListImportsInput *dynamodb.ListImportsInput
		}
		Returns struct {
			ListImportsOutput *dynamodb.ListImportsOutput
			Error             error
		}
		Stub func(*dynamodb.ListImportsInput) (*dynamodb.ListImportsOutput, error)
	}
	ListImportsPagesCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListImportsInput                      *dynamodb.ListImportsInput
			FuncDynamodbListImportsOutputBoolBool func(*dynamodb.ListImportsOutput, bool) bool
		}
		Returns struct {
			Error error
		}
		Stub func(*dynamodb.ListImportsInput, func(*dynamodb.ListImportsOutput, bool) bool) error
	}
	ListImportsPagesWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                               context.Context
			ListImportsInput                      *dynamodb.ListImportsInput
			FuncDynamodbListImportsOutputBoolBool func(*dynamodb.ListImportsOutput, bool) bool
			OptionSlice                           []request.Option
		}
		Returns struct {
			Error error
		}
		Stub func(context.Context, *dynamodb.ListImportsInput, func(*dynamodb.ListImportsOutput, bool) bool, ...request.Option) error
	}
	ListImportsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListImportsInput *dynamodb.ListImportsInput
		}
		Returns struct {
			Request           *request.Request
			ListImportsOutput *dynamodb.ListImportsOutput
		}
		Stub func(*dynamodb.ListImportsInput) (*request.Request, *dynamodb.ListImportsOutput)
	}
	ListImportsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context          context.Context
			ListImportsInput *dynamodb.ListImportsInput
			OptionSlice      []request.Option
		}
		Returns struct {
			ListImportsOutput *dynamodb.ListImportsOutput
			Error             error
		}
		Stub func(context.Context, *dynamodb.ListImportsInput, ...request.Option) (*dynamodb.ListImportsOutput, error)
	}
	ListTablesCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListTablesInput *dynamodb.ListTablesInput
		}
		Returns struct {
			ListTablesOutput *dynamodb.ListTablesOutput
			Error            error
		}
		Stub func(*dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error)
	}
	ListTablesPagesCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListTablesInput                      *dynamodb.ListTablesInput
			FuncDynamodbListTablesOutputBoolBool func(*dynamodb.ListTablesOutput, bool) bool
		}
		Returns struct {
			Error error
		}
		Stub func(*dynamodb.ListTablesInput, func(*dynamodb.ListTablesOutput, bool) bool) error
	}
	ListTablesPagesWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                              context.Context
			ListTablesInput                      *dynamodb.ListTablesInput
			FuncDynamodbListTablesOutputBoolBool func(*dynamodb.ListTablesOutput, bool) bool
			OptionSlice                          []request.Option
		}
		Returns struct {
			Error error
		}
		Stub func(context.Context, *dynamodb.ListTablesInput, func(*dynamodb.ListTablesOutput, bool) bool, ...request.Option) error
	}
	ListTablesRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListTablesInput *dynamodb.ListTablesInput
		}
		Returns struct {
			Request          *request.Request
			ListTablesOutput *dynamodb.ListTablesOutput
		}
		Stub func(*dynamodb.ListTablesInput) (*request.Request, *dynamodb.ListTablesOutput)
	}
	ListTablesWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context         context.Context
			ListTablesInput *dynamodb.ListTablesInput
			OptionSlice     []request.Option
		}
		Returns struct {
			ListTablesOutput *dynamodb.ListTablesOutput
			Error            error
		}
		Stub func(context.Context, *dynamodb.ListTablesInput, ...request.Option) (*dynamodb.ListTablesOutput, error)
	}
	ListTagsOfResourceCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListTagsOfResourceInput *dynamodb.ListTagsOfResourceInput
		}
		Returns struct {
			ListTagsOfResourceOutput *dynamodb.ListTagsOfResourceOutput
			Error                    error
		}
		Stub func(*dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error)
	}
	ListTagsOfResourceRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListTagsOfResourceInput *dynamodb.ListTagsOfResourceInput
		}
		Returns struct {
			Request                  *request.Request
			ListTagsOfResourceOutput *dynamodb.ListTagsOfResourceOutput
		}
		Stub func(*dynamodb.ListTagsOfResourceInput) (*request.Request, *dynamodb.ListTagsOfResourceOutput)
	}
	ListTagsOfResourceWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                 context.Context
			ListTagsOfResourceInput *dynamodb.ListTagsOfResourceInput
			OptionSlice             []request.Option
		}
		Returns struct {
			ListTagsOfResourceOutput *dynamodb.ListTagsOfResourceOutput
			Error                    error
		}
		Stub func(context.Context, *dynamodb.ListTagsOfResourceInput, ...request.Option) (*dynamodb.ListTagsOfResourceOutput, error)
	}
	PutItemCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutItemInput *dynamodb.PutItemInput
		}
		Returns struct {
			PutItemOutput *dynamodb.PutItemOutput
			Error         error
		}
		Stub func(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	}
	PutItemRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutItemInput *dynamodb.PutItemInput
		}
		Returns struct {
			Request       *request.Request
			PutItemOutput *dynamodb.PutItemOutput
		}
		Stub func(*dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput)
	}
	PutItemWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context      context.Context
			PutItemInput *dynamodb.PutItemInput
			OptionSlice  []request.Option
		}
		Returns struct {
			PutItemOutput *dynamodb.PutItemOutput
			Error         error
		}
		Stub func(context.Context, *dynamodb.PutItemInput, ...request.Option) (*dynamodb.PutItemOutput, error)
	}
	QueryCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			QueryInput *dynamodb.QueryInput
		}
		Returns struct {
			QueryOutput *dynamodb.QueryOutput
			Error       error
		}
		Stub func(*dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
	}
	QueryPagesCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			QueryInput                      *dynamodb.QueryInput
			FuncDynamodbQueryOutputBoolBool func(*dynamodb.QueryOutput, bool) bool
		}
		Returns struct {
			Error error
		}
		Stub func(*dynamodb.QueryInput, func(*dynamodb.QueryOutput, bool) bool) error
	}
	QueryPagesWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                         context.Context
			QueryInput                      *dynamodb.QueryInput
			FuncDynamodbQueryOutputBoolBool func(*dynamodb.QueryOutput, bool) bool
			OptionSlice                     []request.Option
		}
		Returns struct {
			Error error
		}
		Stub func(context.Context, *dynamodb.QueryInput, func(*dynamodb.QueryOutput, bool) bool, ...request.Option) error
	}
	QueryRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			QueryInput *dynamodb.QueryInput
		}
		Returns struct {
			Request     *request.Request
			QueryOutput *dynamodb.QueryOutput
		}
		Stub func(*dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput)
	}
	QueryWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context     context.Context
			QueryInput  *dynamodb.QueryInput
			OptionSlice []request.Option
		}
		Returns struct {
			QueryOutput *dynamodb.QueryOutput
			Error       error
		}
		Stub func(context.Context, *dynamodb.QueryInput, ...request.Option) (*dynamodb.QueryOutput, error)
	}
	RestoreTableFromBackupCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			RestoreTableFromBackupInput *dynamodb.RestoreTableFromBackupInput
		}
		Returns struct {
			RestoreTableFromBackupOutput *dynamodb.RestoreTableFromBackupOutput
			Error                        error
		}
		Stub func(*dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error)
	}
	RestoreTableFromBackupRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			RestoreTableFromBackupInput *dynamodb.RestoreTableFromBackupInput
		}
		Returns struct {
			Request                      *request.Request
			RestoreTableFromBackupOutput *dynamodb.RestoreTableFromBackupOutput
		}
		Stub func(*dynamodb.RestoreTableFromBackupInput) (*request.Request, *dynamodb.RestoreTableFromBackupOutput)
	}
	RestoreTableFromBackupWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                     context.Context
			RestoreTableFromBackupInput *dynamodb.RestoreTableFromBackupInput
			OptionSlice                 []request.Option
		}
		Returns struct {
			RestoreTableFromBackupOutput *dynamodb.RestoreTableFromBackupOutput
			Error                        error
		}
		Stub func(context.Context, *dynamodb.RestoreTableFromBackupInput, ...request.Option) (*dynamodb.RestoreTableFromBackupOutput, error)
	}
	RestoreTableToPointInTimeCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			RestoreTableToPointInTimeInput *dynamodb.RestoreTableToPointInTimeInput
		}
		Returns struct {
			RestoreTableToPointInTimeOutput *dynamodb.RestoreTableToPointInTimeOutput
			Error                           error
		}
		Stub func(*dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error)
	}
	RestoreTableToPointInTimeRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			RestoreTableToPointInTimeInput *dynamodb.RestoreTableToPointInTimeInput
		}
		Returns struct {
			Request                         *request.Request
			RestoreTableToPointInTimeOutput *dynamodb.RestoreTableToPointInTimeOutput
		}
		Stub func(*dynamodb.RestoreTableToPointInTimeInput) (*request.Request, *dynamodb.RestoreTableToPointInTimeOutput)
	}
	RestoreTableToPointInTimeWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                        context.Context
			RestoreTableToPointInTimeInput *dynamodb.RestoreTableToPointInTimeInput
			OptionSlice                    []request.Option
		}
		Returns struct {
			RestoreTableToPointInTimeOutput *dynamodb.RestoreTableToPointInTimeOutput
			Error                           error
		}
		Stub func(context.Context, *dynamodb.RestoreTableToPointInTimeInput, ...request.Option) (*dynamodb.RestoreTableToPointInTimeOutput, error)
	}
	ScanCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ScanInput *dynamodb.ScanInput
		}
		Returns struct {
			ScanOutput *dynamodb.ScanOutput
			Error      error
		}
		Stub func(*dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
	}
	ScanPagesCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ScanInput                      *dynamodb.ScanInput
			FuncDynamodbScanOutputBoolBool func(*dynamodb.ScanOutput, bool) bool
		}
		Returns struct {
			Error error
		}
		Stub func(*dynamodb.ScanInput, func(*dynamodb.ScanOutput, bool) bool) error
	}
	ScanPagesWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                        context.Context
			ScanInput                      *dynamodb.ScanInput
			FuncDynamodbScanOutputBoolBool func(*dynamodb.ScanOutput, bool) bool
			OptionSlice                    []request.Option
		}
		Returns struct {
			Error error
		}
		Stub func(context.Context, *dynamodb.ScanInput, func(*dynamodb.ScanOutput, bool) bool, ...request.Option) error
	}
	ScanRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ScanInput *dynamodb.ScanInput
		}
		Returns struct {
			Request    *request.Request
			ScanOutput *dynamodb.ScanOutput
		}
		Stub func(*dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput)
	}
	ScanWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context     context.Context
			ScanInput   *dynamodb.ScanInput
			OptionSlice []request.Option
		}
		Returns struct {
			ScanOutput *dynamodb.ScanOutput
			Error      error
		}
		Stub func(context.Context, *dynamodb.ScanInput, ...request.Option) (*dynamodb.ScanOutput, error)
	}
	TagResourceCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			TagResourceInput *dynamodb.TagResourceInput
		}
		Returns struct {
			TagResourceOutput *dynamodb.TagResourceOutput
			Error             error
		}
		Stub func(*dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error)
	}
	TagResourceRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			TagResourceInput *dynamodb.TagResourceInput
		}
		Returns struct {
			Request           *request.Request
			TagResourceOutput *dynamodb.TagResourceOutput
		}
		Stub func(*dynamodb.TagResourceInput) (*request.Request, *dynamodb.TagResourceOutput)
	}
	TagResourceWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context          context.Context
			TagResourceInput *dynamodb.TagResourceInput
			OptionSlice      []request.Option
		}
		Returns struct {
			TagResourceOutput *dynamodb.TagResourceOutput
			Error             error
		}
		Stub func(context.Context, *dynamodb.TagResourceInput, ...request.Option) (*dynamodb.TagResourceOutput, error)
	}
	TransactGetItemsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			TransactGetItemsInput *dynamodb.TransactGetItemsInput
		}
		Returns struct {
			TransactGetItemsOutput *dynamodb.TransactGetItemsOutput
			Error                  error
		}
		Stub func(*dynamodb.TransactGetItemsInput) (*dynamodb.TransactGetItemsOutput, error)
	}
	TransactGetItemsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			TransactGetItemsInput *dynamodb.TransactGetItemsInput
		}
		Returns struct {
			Request                *request.Request
			TransactGetItemsOutput *dynamodb.TransactGetItemsOutput
		}
		Stub func(*dynamodb.TransactGetItemsInput) (*request.Request, *dynamodb.TransactGetItemsOutput)
	}
	TransactGetItemsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context               context.Context
			TransactGetItemsInput *dynamodb.TransactGetItemsInput
			OptionSlice           []request.Option
		}
		Returns struct {
			TransactGetItemsOutput *dynamodb.TransactGetItemsOutput
			Error                  error
		}
		Stub func(context.Context, *dynamodb.TransactGetItemsInput, ...request.Option) (*dynamodb.TransactGetItemsOutput, error)
	}
	TransactWriteItemsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			TransactWriteItemsInput *dynamodb.TransactWriteItemsInput
		}
		Returns struct {
			TransactWriteItemsOutput *dynamodb.TransactWriteItemsOutput
			Error                    error
		}
		Stub func(*dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error)
	}
	TransactWriteItemsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			TransactWriteItemsInput *dynamodb.TransactWriteItemsInput
		}
		Returns struct {
			Request                  *request.Request
			TransactWriteItemsOutput *dynamodb.TransactWriteItemsOutput
		}
		Stub func(*dynamodb.TransactWriteItemsInput) (*request.Request, *dynamodb.TransactWriteItemsOutput)
	}
	TransactWriteItemsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                 context.Context
			TransactWriteItemsInput *dynamodb.TransactWriteItemsInput
			OptionSlice             []request.Option
		}
		Returns struct {
			TransactWriteItemsOutput *dynamodb.TransactWriteItemsOutput
			Error                    error
		}
		Stub func(context.Context, *dynamodb.TransactWriteItemsInput, ...request.Option) (*dynamodb.TransactWriteItemsOutput, error)
	}
	UntagResourceCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UntagResourceInput *dynamodb.UntagResourceInput
		}
		Returns struct {
			UntagResourceOutput *dynamodb.UntagResourceOutput
			Error               error
		}
		Stub func(*dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error)
	}
	UntagResourceRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UntagResourceInput *dynamodb.UntagResourceInput
		}
		Returns struct {
			Request             *request.Request
			UntagResourceOutput *dynamodb.UntagResourceOutput
		}
		Stub func(*dynamodb.UntagResourceInput) (*request.Request, *dynamodb.UntagResourceOutput)
	}
	UntagResourceWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context            context.Context
			UntagResourceInput *dynamodb.UntagResourceInput
			OptionSlice        []request.Option
		}
		Returns struct {
			UntagResourceOutput *dynamodb.UntagResourceOutput
			Error               error
		}
		Stub func(context.Context, *dynamodb.UntagResourceInput, ...request.Option) (*dynamodb.UntagResourceOutput, error)
	}
	UpdateContinuousBackupsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UpdateContinuousBackupsInput *dynamodb.UpdateContinuousBackupsInput
		}
		Returns struct {
			UpdateContinuousBackupsOutput *dynamodb.UpdateContinuousBackupsOutput
			Error                         error
		}
		Stub func(*dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error)
	}
	UpdateContinuousBackupsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UpdateContinuousBackupsInput *dynamodb.UpdateContinuousBackupsInput
		}
		Returns struct {
			Request                       *request.Request
			UpdateContinuousBackupsOutput *dynamodb.UpdateContinuousBackupsOutput
		}
		Stub func(*dynamodb.UpdateContinuousBackupsInput) (*request.Request, *dynamodb.UpdateContinuousBackupsOutput)
	}
	UpdateContinuousBackupsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                      context.Context
			UpdateContinuousBackupsInput *dynamodb.UpdateContinuousBackupsInput
			OptionSlice                  []request.Option
		}
		Returns struct {
			UpdateContinuousBackupsOutput *dynamodb.UpdateContinuousBackupsOutput
			Error                         error
		}
		Stub func(context.Context, *dynamodb.UpdateContinuousBackupsInput, ...request.Option) (*dynamodb.UpdateContinuousBackupsOutput, error)
	}
	UpdateContributorInsightsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UpdateContributorInsightsInput *dynamodb.UpdateContributorInsightsInput
		}
		Returns struct {
			UpdateContributorInsightsOutput *dynamodb.UpdateContributorInsightsOutput
			Error                           error
		}
		Stub func(*dynamodb.UpdateContributorInsightsInput) (*dynamodb.UpdateContributorInsightsOutput, error)
	}
	UpdateContributorInsightsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UpdateContributorInsightsInput *dynamodb.UpdateContributorInsightsInput
		}
		Returns struct {
			Request                         *request.Request
			UpdateContributorInsightsOutput *dynamodb.UpdateContributorInsightsOutput
		}
		Stub func(*dynamodb.UpdateContributorInsightsInput) (*request.Request, *dynamodb.UpdateContributorInsightsOutput)
	}
	UpdateContributorInsightsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                        context.Context
			UpdateContributorInsightsInput *dynamodb.UpdateContributorInsightsInput
			OptionSlice                    []request.Option
		}
		Returns struct {
			UpdateContributorInsightsOutput *dynamodb.UpdateContributorInsightsOutput
			Error                           error
		}
		Stub func(context.Context, *dynamodb.UpdateContributorInsightsInput, ...request.Option) (*dynamodb.UpdateContributorInsightsOutput, error)
	}
	UpdateGlobalTableCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UpdateGlobalTableInput *dynamodb.UpdateGlobalTableInput
		}
		Returns struct {
			UpdateGlobalTableOutput *dynamodb.UpdateGlobalTableOutput
			Error                   error
		}
		Stub func(*dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error)
	}
	UpdateGlobalTableRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UpdateGlobalTableInput *dynamodb.UpdateGlobalTableInput
		}
		Returns struct {
			Request                 *request.Request
			UpdateGlobalTableOutput *dynamodb.UpdateGlobalTableOutput
		}
		Stub func(*dynamodb.UpdateGlobalTableInput) (*request.Request, *dynamodb.UpdateGlobalTableOutput)
	}
	UpdateGlobalTableSettingsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UpdateGlobalTableSettingsInput *dynamodb.UpdateGlobalTableSettingsInput
		}
		Returns struct {
			UpdateGlobalTableSettingsOutput *dynamodb.UpdateGlobalTableSettingsOutput
			Error                           error
		}
		Stub func(*dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error)
	}
	UpdateGlobalTableSettingsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UpdateGlobalTableSettingsInput *dynamodb.UpdateGlobalTableSettingsInput
		}
		Returns struct {
			Request                         *request.Request
			UpdateGlobalTableSettingsOutput *dynamodb.UpdateGlobalTableSettingsOutput
		}
		Stub func(*dynamodb.UpdateGlobalTableSettingsInput) (*request.Request, *dynamodb.UpdateGlobalTableSettingsOutput)
	}
	UpdateGlobalTableSettingsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                        context.Context
			UpdateGlobalTableSettingsInput *dynamodb.UpdateGlobalTableSettingsInput
			OptionSlice                    []request.Option
		}
		Returns struct {
			UpdateGlobalTableSettingsOutput *dynamodb.UpdateGlobalTableSettingsOutput
			Error                           error
		}
		Stub func(context.Context, *dynamodb.UpdateGlobalTableSettingsInput, ...request.Option) (*dynamodb.UpdateGlobalTableSettingsOutput, error)
	}
	UpdateGlobalTableWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                context.Context
			UpdateGlobalTableInput *dynamodb.UpdateGlobalTableInput
			OptionSlice            []request.Option
		}
		Returns struct {
			UpdateGlobalTableOutput *dynamodb.UpdateGlobalTableOutput
			Error                   error
		}
		Stub func(context.Context, *dynamodb.UpdateGlobalTableInput, ...request.Option) (*dynamodb.UpdateGlobalTableOutput, error)
	}
	UpdateItemCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UpdateItemInput *dynamodb.UpdateItemInput
		}
		Returns struct {
			UpdateItemOutput *dynamodb.UpdateItemOutput
			Error            error
		}
		Stub func(*dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)
	}
	UpdateItemRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UpdateItemInput *dynamodb.UpdateItemInput
		}
		Returns struct {
			Request          *request.Request
			UpdateItemOutput *dynamodb.UpdateItemOutput
		}
		Stub func(*dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput)
	}
	UpdateItemWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context         context.Context
			UpdateItemInput *dynamodb.UpdateItemInput
			OptionSlice     []request.Option
		}
		Returns struct {
			UpdateItemOutput *dynamodb.UpdateItemOutput
			Error            error
		}
		Stub func(context.Context, *dynamodb.UpdateItemInput, ...request.Option) (*dynamodb.UpdateItemOutput, error)
	}
	UpdateTableCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UpdateTableInput *dynamodb.UpdateTableInput
		}
		Returns struct {
			UpdateTableOutput *dynamodb.UpdateTableOutput
			Error             error
		}
		Stub func(*dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error)
	}
	UpdateTableReplicaAutoScalingCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UpdateTableReplicaAutoScalingInput *dynamodb.UpdateTableReplicaAutoScalingInput
		}
		Returns struct {
			UpdateTableReplicaAutoScalingOutput *dynamodb.UpdateTableReplicaAutoScalingOutput
			Error                               error
		}
		Stub func(*dynamodb.UpdateTableReplicaAutoScalingInput) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error)
	}
	UpdateTableReplicaAutoScalingRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UpdateTableReplicaAutoScalingInput *dynamodb.UpdateTableReplicaAutoScalingInput
		}
		Returns struct {
			Request                             *request.Request
			UpdateTableReplicaAutoScalingOutput *dynamodb.UpdateTableReplicaAutoScalingOutput
		}
		Stub func(*dynamodb.UpdateTableReplicaAutoScalingInput) (*request.Request, *dynamodb.UpdateTableReplicaAutoScalingOutput)
	}
	UpdateTableReplicaAutoScalingWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                            context.Context
			UpdateTableReplicaAutoScalingInput *dynamodb.UpdateTableReplicaAutoScalingInput
			OptionSlice                        []request.Option
		}
		Returns struct {
			UpdateTableReplicaAutoScalingOutput *dynamodb.UpdateTableReplicaAutoScalingOutput
			Error                               error
		}
		Stub func(context.Context, *dynamodb.UpdateTableReplicaAutoScalingInput, ...request.Option) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error)
	}
	UpdateTableRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UpdateTableInput *dynamodb.UpdateTableInput
		}
		Returns struct {
			Request           *request.Request
			UpdateTableOutput *dynamodb.UpdateTableOutput
		}
		Stub func(*dynamodb.UpdateTableInput) (*request.Request, *dynamodb.UpdateTableOutput)
	}
	UpdateTableWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context          context.Context
			UpdateTableInput *dynamodb.UpdateTableInput
			OptionSlice      []request.Option
		}
		Returns struct {
			UpdateTableOutput *dynamodb.UpdateTableOutput
			Error             error
		}
		Stub func(context.Context, *dynamodb.UpdateTableInput, ...request.Option) (*dynamodb.UpdateTableOutput, error)
	}
	UpdateTimeToLiveCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UpdateTimeToLiveInput *dynamodb.UpdateTimeToLiveInput
		}
		Returns struct {
			UpdateTimeToLiveOutput *dynamodb.UpdateTimeToLiveOutput
			Error                  error
		}
		Stub func(*dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error)
	}
	UpdateTimeToLiveRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UpdateTimeToLiveInput *dynamodb.UpdateTimeToLiveInput
		}
		Returns struct {
			Request                *request.Request
			UpdateTimeToLiveOutput *dynamodb.UpdateTimeToLiveOutput
		}
		Stub func(*dynamodb.UpdateTimeToLiveInput) (*request.Request, *dynamodb.UpdateTimeToLiveOutput)
	}
	UpdateTimeToLiveWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context               context.Context
			UpdateTimeToLiveInput *dynamodb.UpdateTimeToLiveInput
			OptionSlice           []request.Option
		}
		Returns struct {
			UpdateTimeToLiveOutput *dynamodb.UpdateTimeToLiveOutput
			Error                  error
		}
		Stub func(context.Context, *dynamodb.UpdateTimeToLiveInput, ...request.Option) (*dynamodb.UpdateTimeToLiveOutput, error)
	}
	WaitUntilTableExistsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeTableInput *dynamodb.DescribeTableInput
		}
		Returns struct {
			Error error
		}
		Stub func(*dynamodb.DescribeTableInput) error
	}
	WaitUntilTableExistsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context            context.Context
			DescribeTableInput *dynamodb.DescribeTableInput
			WaiterOptionSlice  []request.WaiterOption
		}
		Returns struct {
			Error error
		}
		Stub func(context.Context, *dynamodb.DescribeTableInput, ...request.WaiterOption) error
	}
	WaitUntilTableNotExistsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DescribeTableInput *dynamodb.DescribeTableInput
		}
		Returns struct {
			Error error
		}
		Stub func(*dynamodb.DescribeTableInput) error
	}
	WaitUntilTableNotExistsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context            context.Context
			DescribeTableInput *dynamodb.DescribeTableInput
			WaiterOptionSlice  []request.WaiterOption
		}
		Returns struct {
			Error error
		}
		Stub func(context.Context, *dynamodb.DescribeTableInput, ...request.WaiterOption) error
	}
}

func (f *DynamoDBAPI) BatchExecuteStatement(param1 *dynamodb.BatchExecuteStatementInput) (*dynamodb.BatchExecuteStatementOutput, error) {
	f.BatchExecuteStatementCall.mutex.Lock()
	defer f.BatchExecuteStatementCall.mutex.Unlock()
	f.BatchExecuteStatementCall.CallCount++
	f.BatchExecuteStatementCall.Receives.BatchExecuteStatementInput = param1
	if f.BatchExecuteStatementCall.Stub != nil {
		return f.BatchExecuteStatementCall.Stub(param1)
	}
	return f.BatchExecuteStatementCall.Returns.BatchExecuteStatementOutput, f.BatchExecuteStatementCall.Returns.Error
}
func (f *DynamoDBAPI) BatchExecuteStatementRequest(param1 *dynamodb.BatchExecuteStatementInput) (*request.Request, *dynamodb.BatchExecuteStatementOutput) {
	f.BatchExecuteStatementRequestCall.mutex.Lock()
	defer f.BatchExecuteStatementRequestCall.mutex.Unlock()
	f.BatchExecuteStatementRequestCall.CallCount++
	f.BatchExecuteStatementRequestCall.Receives.BatchExecuteStatementInput = param1
	if f.BatchExecuteStatementRequestCall.Stub != nil {
		return f.BatchExecuteStatementRequestCall.Stub(param1)
	}
	return f.BatchExecuteStatementRequestCall.Returns.Request, f.BatchExecuteStatementRequestCall.Returns.BatchExecuteStatementOutput
}
func (f *DynamoDBAPI) BatchExecuteStatementWithContext(param1 context.Context, param2 *dynamodb.BatchExecuteStatementInput, param3 ...request.Option) (*dynamodb.BatchExecuteStatementOutput, error) {
	f.BatchExecuteStatementWithContextCall.mutex.Lock()
	defer f.BatchExecuteStatementWithContextCall.mutex.Unlock()
	f.BatchExecuteStatementWithContextCall.CallCount++
	f.BatchExecuteStatementWithContextCall.Receives.Context = param1
	f.BatchExecuteStatementWithContextCall.Receives.BatchExecuteStatementInput = param2
	f.BatchExecuteStatementWithContextCall.Receives.OptionSlice = param3
	if f.BatchExecuteStatementWithContextCall.Stub != nil {
		return f.BatchExecuteStatementWithContextCall.Stub(param1, param2, param3...)
	}
	return f.BatchExecuteStatementWithContextCall.Returns.BatchExecuteStatementOutput, f.BatchExecuteStatementWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) BatchGetItem(param1 *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	f.BatchGetItemCall.mutex.Lock()
	defer f.BatchGetItemCall.mutex.Unlock()
	f.BatchGetItemCall.CallCount++
	f.BatchGetItemCall.Receives.BatchGetItemInput = param1
	if f.BatchGetItemCall.Stub != nil {
		return f.BatchGetItemCall.Stub(param1)
	}
	return f.BatchGetItemCall.Returns.BatchGetItemOutput, f.BatchGetItemCall.Returns.Error
}
func (f *DynamoDBAPI) BatchGetItemPages(param1 *dynamodb.BatchGetItemInput, param2 func(*dynamodb.BatchGetItemOutput, bool) bool) error {
	f.BatchGetItemPagesCall.mutex.Lock()
	defer f.BatchGetItemPagesCall.mutex.Unlock()
	f.BatchGetItemPagesCall.CallCount++
	f.BatchGetItemPagesCall.Receives.BatchGetItemInput = param1
	f.BatchGetItemPagesCall.Receives.FuncDynamodbBatchGetItemOutputBoolBool = param2
	if f.BatchGetItemPagesCall.Stub != nil {
		return f.BatchGetItemPagesCall.Stub(param1, param2)
	}
	return f.BatchGetItemPagesCall.Returns.Error
}
func (f *DynamoDBAPI) BatchGetItemPagesWithContext(param1 context.Context, param2 *dynamodb.BatchGetItemInput, param3 func(*dynamodb.BatchGetItemOutput, bool) bool, param4 ...request.Option) error {
	f.BatchGetItemPagesWithContextCall.mutex.Lock()
	defer f.BatchGetItemPagesWithContextCall.mutex.Unlock()
	f.BatchGetItemPagesWithContextCall.CallCount++
	f.BatchGetItemPagesWithContextCall.Receives.Context = param1
	f.BatchGetItemPagesWithContextCall.Receives.BatchGetItemInput = param2
	f.BatchGetItemPagesWithContextCall.Receives.FuncDynamodbBatchGetItemOutputBoolBool = param3
	f.BatchGetItemPagesWithContextCall.Receives.OptionSlice = param4
	if f.BatchGetItemPagesWithContextCall.Stub != nil {
		return f.BatchGetItemPagesWithContextCall.Stub(param1, param2, param3, param4...)
	}
	return f.BatchGetItemPagesWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) BatchGetItemRequest(param1 *dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput) {
	f.BatchGetItemRequestCall.mutex.Lock()
	defer f.BatchGetItemRequestCall.mutex.Unlock()
	f.BatchGetItemRequestCall.CallCount++
	f.BatchGetItemRequestCall.Receives.BatchGetItemInput = param1
	if f.BatchGetItemRequestCall.Stub != nil {
		return f.BatchGetItemRequestCall.Stub(param1)
	}
	return f.BatchGetItemRequestCall.Returns.Request, f.BatchGetItemRequestCall.Returns.BatchGetItemOutput
}
func (f *DynamoDBAPI) BatchGetItemWithContext(param1 context.Context, param2 *dynamodb.BatchGetItemInput, param3 ...request.Option) (*dynamodb.BatchGetItemOutput, error) {
	f.BatchGetItemWithContextCall.mutex.Lock()
	defer f.BatchGetItemWithContextCall.mutex.Unlock()
	f.BatchGetItemWithContextCall.CallCount++
	f.BatchGetItemWithContextCall.Receives.Context = param1
	f.BatchGetItemWithContextCall.Receives.BatchGetItemInput = param2
	f.BatchGetItemWithContextCall.Receives.OptionSlice = param3
	if f.BatchGetItemWithContextCall.Stub != nil {
		return f.BatchGetItemWithContextCall.Stub(param1, param2, param3...)
	}
	return f.BatchGetItemWithContextCall.Returns.BatchGetItemOutput, f.BatchGetItemWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) BatchWriteItem(param1 *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	f.BatchWriteItemCall.mutex.Lock()
	defer f.BatchWriteItemCall.mutex.Unlock()
	f.BatchWriteItemCall.CallCount++
	f.BatchWriteItemCall.Receives.BatchWriteItemInput = param1
	if f.BatchWriteItemCall.Stub != nil {
		return f.BatchWriteItemCall.Stub(param1)
	}
	return f.BatchWriteItemCall.Returns.BatchWriteItemOutput, f.BatchWriteItemCall.Returns.Error
}
func (f *DynamoDBAPI) BatchWriteItemRequest(param1 *dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput) {
	f.BatchWriteItemRequestCall.mutex.Lock()
	defer f.BatchWriteItemRequestCall.mutex.Unlock()
	f.BatchWriteItemRequestCall.CallCount++
	f.BatchWriteItemRequestCall.Receives.BatchWriteItemInput = param1
	if f.BatchWriteItemRequestCall.Stub != nil {
		return f.BatchWriteItemRequestCall.Stub(param1)
	}
	return f.BatchWriteItemRequestCall.Returns.Request, f.BatchWriteItemRequestCall.Returns.BatchWriteItemOutput
}
func (f *DynamoDBAPI) BatchWriteItemWithContext(param1 context.Context, param2 *dynamodb.BatchWriteItemInput, param3 ...request.Option) (*dynamodb.BatchWriteItemOutput, error) {
	f.BatchWriteItemWithContextCall.mutex.Lock()
	defer f.BatchWriteItemWithContextCall.mutex.Unlock()
	f.BatchWriteItemWithContextCall.CallCount++
	f.BatchWriteItemWithContextCall.Receives.Context = param1
	f.BatchWriteItemWithContextCall.Receives.BatchWriteItemInput = param2
	f.BatchWriteItemWithContextCall.Receives.OptionSlice = param3
	if f.BatchWriteItemWithContextCall.Stub != nil {
		return f.BatchWriteItemWithContextCall.Stub(param1, param2, param3...)
	}
	return f.BatchWriteItemWithContextCall.Returns.BatchWriteItemOutput, f.BatchWriteItemWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) CreateBackup(param1 *dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error) {
	f.CreateBackupCall.mutex.Lock()
	defer f.CreateBackupCall.mutex.Unlock()
	f.CreateBackupCall.CallCount++
	f.CreateBackupCall.Receives.CreateBackupInput = param1
	if f.CreateBackupCall.Stub != nil {
		return f.CreateBackupCall.Stub(param1)
	}
	return f.CreateBackupCall.Returns.CreateBackupOutput, f.CreateBackupCall.Returns.Error
}
func (f *DynamoDBAPI) CreateBackupRequest(param1 *dynamodb.CreateBackupInput) (*request.Request, *dynamodb.CreateBackupOutput) {
	f.CreateBackupRequestCall.mutex.Lock()
	defer f.CreateBackupRequestCall.mutex.Unlock()
	f.CreateBackupRequestCall.CallCount++
	f.CreateBackupRequestCall.Receives.CreateBackupInput = param1
	if f.CreateBackupRequestCall.Stub != nil {
		return f.CreateBackupRequestCall.Stub(param1)
	}
	return f.CreateBackupRequestCall.Returns.Request, f.CreateBackupRequestCall.Returns.CreateBackupOutput
}
func (f *DynamoDBAPI) CreateBackupWithContext(param1 context.Context, param2 *dynamodb.CreateBackupInput, param3 ...request.Option) (*dynamodb.CreateBackupOutput, error) {
	f.CreateBackupWithContextCall.mutex.Lock()
	defer f.CreateBackupWithContextCall.mutex.Unlock()
	f.CreateBackupWithContextCall.CallCount++
	f.CreateBackupWithContextCall.Receives.Context = param1
	f.CreateBackupWithContextCall.Receives.CreateBackupInput = param2
	f.CreateBackupWithContextCall.Receives.OptionSlice = param3
	if f.CreateBackupWithContextCall.Stub != nil {
		return f.CreateBackupWithContextCall.Stub(param1, param2, param3...)
	}
	return f.CreateBackupWithContextCall.Returns.CreateBackupOutput, f.CreateBackupWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) CreateGlobalTable(param1 *dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error) {
	f.CreateGlobalTableCall.mutex.Lock()
	defer f.CreateGlobalTableCall.mutex.Unlock()
	f.CreateGlobalTableCall.CallCount++
	f.CreateGlobalTableCall.Receives.CreateGlobalTableInput = param1
	if f.CreateGlobalTableCall.Stub != nil {
		return f.CreateGlobalTableCall.Stub(param1)
	}
	return f.CreateGlobalTableCall.Returns.CreateGlobalTableOutput, f.CreateGlobalTableCall.Returns.Error
}
func (f *DynamoDBAPI) CreateGlobalTableRequest(param1 *dynamodb.CreateGlobalTableInput) (*request.Request, *dynamodb.CreateGlobalTableOutput) {
	f.CreateGlobalTableRequestCall.mutex.Lock()
	defer f.CreateGlobalTableRequestCall.mutex.Unlock()
	f.CreateGlobalTableRequestCall.CallCount++
	f.CreateGlobalTableRequestCall.Receives.CreateGlobalTableInput = param1
	if f.CreateGlobalTableRequestCall.Stub != nil {
		return f.CreateGlobalTableRequestCall.Stub(param1)
	}
	return f.CreateGlobalTableRequestCall.Returns.Request, f.CreateGlobalTableRequestCall.Returns.CreateGlobalTableOutput
}
func (f *DynamoDBAPI) CreateGlobalTableWithContext(param1 context.Context, param2 *dynamodb.CreateGlobalTableInput, param3 ...request.Option) (*dynamodb.CreateGlobalTableOutput, error) {
	f.CreateGlobalTableWithContextCall.mutex.Lock()
	defer f.CreateGlobalTableWithContextCall.mutex.Unlock()
	f.CreateGlobalTableWithContextCall.CallCount++
	f.CreateGlobalTableWithContextCall.Receives.Context = param1
	f.CreateGlobalTableWithContextCall.Receives.CreateGlobalTableInput = param2
	f.CreateGlobalTableWithContextCall.Receives.OptionSlice = param3
	if f.CreateGlobalTableWithContextCall.Stub != nil {
		return f.CreateGlobalTableWithContextCall.Stub(param1, param2, param3...)
	}
	return f.CreateGlobalTableWithContextCall.Returns.CreateGlobalTableOutput, f.CreateGlobalTableWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) CreateTable(param1 *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	f.CreateTableCall.mutex.Lock()
	defer f.CreateTableCall.mutex.Unlock()
	f.CreateTableCall.CallCount++
	f.CreateTableCall.Receives.CreateTableInput = param1
	if f.CreateTableCall.Stub != nil {
		return f.CreateTableCall.Stub(param1)
	}
	return f.CreateTableCall.Returns.CreateTableOutput, f.CreateTableCall.Returns.Error
}
func (f *DynamoDBAPI) CreateTableRequest(param1 *dynamodb.CreateTableInput) (*request.Request, *dynamodb.CreateTableOutput) {
	f.CreateTableRequestCall.mutex.Lock()
	defer f.CreateTableRequestCall.mutex.Unlock()
	f.CreateTableRequestCall.CallCount++
	f.CreateTableRequestCall.Receives.CreateTableInput = param1
	if f.CreateTableRequestCall.Stub != nil {
		return f.CreateTableRequestCall.Stub(param1)
	}
	return f.CreateTableRequestCall.Returns.Request, f.CreateTableRequestCall.Returns.CreateTableOutput
}
func (f *DynamoDBAPI) CreateTableWithContext(param1 context.Context, param2 *dynamodb.CreateTableInput, param3 ...request.Option) (*dynamodb.CreateTableOutput, error) {
	f.CreateTableWithContextCall.mutex.Lock()
	defer f.CreateTableWithContextCall.mutex.Unlock()
	f.CreateTableWithContextCall.CallCount++
	f.CreateTableWithContextCall.Receives.Context = param1
	f.CreateTableWithContextCall.Receives.CreateTableInput = param2
	f.CreateTableWithContextCall.Receives.OptionSlice = param3
	if f.CreateTableWithContextCall.Stub != nil {
		return f.CreateTableWithContextCall.Stub(param1, param2, param3...)
	}
	return f.CreateTableWithContextCall.Returns.CreateTableOutput, f.CreateTableWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) DeleteBackup(param1 *dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error) {
	f.DeleteBackupCall.mutex.Lock()
	defer f.DeleteBackupCall.mutex.Unlock()
	f.DeleteBackupCall.CallCount++
	f.DeleteBackupCall.Receives.DeleteBackupInput = param1
	if f.DeleteBackupCall.Stub != nil {
		return f.DeleteBackupCall.Stub(param1)
	}
	return f.DeleteBackupCall.Returns.DeleteBackupOutput, f.DeleteBackupCall.Returns.Error
}
func (f *DynamoDBAPI) DeleteBackupRequest(param1 *dynamodb.DeleteBackupInput) (*request.Request, *dynamodb.DeleteBackupOutput) {
	f.DeleteBackupRequestCall.mutex.Lock()
	defer f.DeleteBackupRequestCall.mutex.Unlock()
	f.DeleteBackupRequestCall.CallCount++
	f.DeleteBackupRequestCall.Receives.DeleteBackupInput = param1
	if f.DeleteBackupRequestCall.Stub != nil {
		return f.DeleteBackupRequestCall.Stub(param1)
	}
	return f.DeleteBackupRequestCall.Returns.Request, f.DeleteBackupRequestCall.Returns.DeleteBackupOutput
}
func (f *DynamoDBAPI) DeleteBackupWithContext(param1 context.Context, param2 *dynamodb.DeleteBackupInput, param3 ...request.Option) (*dynamodb.DeleteBackupOutput, error) {
	f.DeleteBackupWithContextCall.mutex.Lock()
	defer f.DeleteBackupWithContextCall.mutex.Unlock()
	f.DeleteBackupWithContextCall.CallCount++
	f.DeleteBackupWithContextCall.Receives.Context = param1
	f.DeleteBackupWithContextCall.Receives.DeleteBackupInput = param2
	f.DeleteBackupWithContextCall.Receives.OptionSlice = param3
	if f.DeleteBackupWithContextCall.Stub != nil {
		return f.DeleteBackupWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DeleteBackupWithContextCall.Returns.DeleteBackupOutput, f.DeleteBackupWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) DeleteItem(param1 *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	f.DeleteItemCall.mutex.Lock()
	defer f.DeleteItemCall.mutex.Unlock()
	f.DeleteItemCall.CallCount++
	f.DeleteItemCall.Receives.DeleteItemInput = param1
	if f.DeleteItemCall.Stub != nil {
		return f.DeleteItemCall.Stub(param1)
	}
	return f.DeleteItemCall.Returns.DeleteItemOutput, f.DeleteItemCall.Returns.Error
}
func (f *DynamoDBAPI) DeleteItemRequest(param1 *dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput) {
	f.DeleteItemRequestCall.mutex.Lock()
	defer f.DeleteItemRequestCall.mutex.Unlock()
	f.DeleteItemRequestCall.CallCount++
	f.DeleteItemRequestCall.Receives.DeleteItemInput = param1
	if f.DeleteItemRequestCall.Stub != nil {
		return f.DeleteItemRequestCall.Stub(param1)
	}
	return f.DeleteItemRequestCall.Returns.Request, f.DeleteItemRequestCall.Returns.DeleteItemOutput
}
func (f *DynamoDBAPI) DeleteItemWithContext(param1 context.Context, param2 *dynamodb.DeleteItemInput, param3 ...request.Option) (*dynamodb.DeleteItemOutput, error) {
	f.DeleteItemWithContextCall.mutex.Lock()
	defer f.DeleteItemWithContextCall.mutex.Unlock()
	f.DeleteItemWithContextCall.CallCount++
	f.DeleteItemWithContextCall.Receives.Context = param1
	f.DeleteItemWithContextCall.Receives.DeleteItemInput = param2
	f.DeleteItemWithContextCall.Receives.OptionSlice = param3
	if f.DeleteItemWithContextCall.Stub != nil {
		return f.DeleteItemWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DeleteItemWithContextCall.Returns.DeleteItemOutput, f.DeleteItemWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) DeleteTable(param1 *dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error) {
	f.DeleteTableCall.mutex.Lock()
	defer f.DeleteTableCall.mutex.Unlock()
	f.DeleteTableCall.CallCount++
	f.DeleteTableCall.Receives.DeleteTableInput = param1
	if f.DeleteTableCall.Stub != nil {
		return f.DeleteTableCall.Stub(param1)
	}
	return f.DeleteTableCall.Returns.DeleteTableOutput, f.DeleteTableCall.Returns.Error
}
func (f *DynamoDBAPI) DeleteTableRequest(param1 *dynamodb.DeleteTableInput) (*request.Request, *dynamodb.DeleteTableOutput) {
	f.DeleteTableRequestCall.mutex.Lock()
	defer f.DeleteTableRequestCall.mutex.Unlock()
	f.DeleteTableRequestCall.CallCount++
	f.DeleteTableRequestCall.Receives.DeleteTableInput = param1
	if f.DeleteTableRequestCall.Stub != nil {
		return f.DeleteTableRequestCall.Stub(param1)
	}
	return f.DeleteTableRequestCall.Returns.Request, f.DeleteTableRequestCall.Returns.DeleteTableOutput
}
func (f *DynamoDBAPI) DeleteTableWithContext(param1 context.Context, param2 *dynamodb.DeleteTableInput, param3 ...request.Option) (*dynamodb.DeleteTableOutput, error) {
	f.DeleteTableWithContextCall.mutex.Lock()
	defer f.DeleteTableWithContextCall.mutex.Unlock()
	f.DeleteTableWithContextCall.CallCount++
	f.DeleteTableWithContextCall.Receives.Context = param1
	f.DeleteTableWithContextCall.Receives.DeleteTableInput = param2
	f.DeleteTableWithContextCall.Receives.OptionSlice = param3
	if f.DeleteTableWithContextCall.Stub != nil {
		return f.DeleteTableWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DeleteTableWithContextCall.Returns.DeleteTableOutput, f.DeleteTableWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeBackup(param1 *dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error) {
	f.DescribeBackupCall.mutex.Lock()
	defer f.DescribeBackupCall.mutex.Unlock()
	f.DescribeBackupCall.CallCount++
	f.DescribeBackupCall.Receives.DescribeBackupInput = param1
	if f.DescribeBackupCall.Stub != nil {
		return f.DescribeBackupCall.Stub(param1)
	}
	return f.DescribeBackupCall.Returns.DescribeBackupOutput, f.DescribeBackupCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeBackupRequest(param1 *dynamodb.DescribeBackupInput) (*request.Request, *dynamodb.DescribeBackupOutput) {
	f.DescribeBackupRequestCall.mutex.Lock()
	defer f.DescribeBackupRequestCall.mutex.Unlock()
	f.DescribeBackupRequestCall.CallCount++
	f.DescribeBackupRequestCall.Receives.DescribeBackupInput = param1
	if f.DescribeBackupRequestCall.Stub != nil {
		return f.DescribeBackupRequestCall.Stub(param1)
	}
	return f.DescribeBackupRequestCall.Returns.Request, f.DescribeBackupRequestCall.Returns.DescribeBackupOutput
}
func (f *DynamoDBAPI) DescribeBackupWithContext(param1 context.Context, param2 *dynamodb.DescribeBackupInput, param3 ...request.Option) (*dynamodb.DescribeBackupOutput, error) {
	f.DescribeBackupWithContextCall.mutex.Lock()
	defer f.DescribeBackupWithContextCall.mutex.Unlock()
	f.DescribeBackupWithContextCall.CallCount++
	f.DescribeBackupWithContextCall.Receives.Context = param1
	f.DescribeBackupWithContextCall.Receives.DescribeBackupInput = param2
	f.DescribeBackupWithContextCall.Receives.OptionSlice = param3
	if f.DescribeBackupWithContextCall.Stub != nil {
		return f.DescribeBackupWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DescribeBackupWithContextCall.Returns.DescribeBackupOutput, f.DescribeBackupWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeContinuousBackups(param1 *dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	f.DescribeContinuousBackupsCall.mutex.Lock()
	defer f.DescribeContinuousBackupsCall.mutex.Unlock()
	f.DescribeContinuousBackupsCall.CallCount++
	f.DescribeContinuousBackupsCall.Receives.DescribeContinuousBackupsInput = param1
	if f.DescribeContinuousBackupsCall.Stub != nil {
		return f.DescribeContinuousBackupsCall.Stub(param1)
	}
	return f.DescribeContinuousBackupsCall.Returns.DescribeContinuousBackupsOutput, f.DescribeContinuousBackupsCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeContinuousBackupsRequest(param1 *dynamodb.DescribeContinuousBackupsInput) (*request.Request, *dynamodb.DescribeContinuousBackupsOutput) {
	f.DescribeContinuousBackupsRequestCall.mutex.Lock()
	defer f.DescribeContinuousBackupsRequestCall.mutex.Unlock()
	f.DescribeContinuousBackupsRequestCall.CallCount++
	f.DescribeContinuousBackupsRequestCall.Receives.DescribeContinuousBackupsInput = param1
	if f.DescribeContinuousBackupsRequestCall.Stub != nil {
		return f.DescribeContinuousBackupsRequestCall.Stub(param1)
	}
	return f.DescribeContinuousBackupsRequestCall.Returns.Request, f.DescribeContinuousBackupsRequestCall.Returns.DescribeContinuousBackupsOutput
}
func (f *DynamoDBAPI) DescribeContinuousBackupsWithContext(param1 context.Context, param2 *dynamodb.DescribeContinuousBackupsInput, param3 ...request.Option) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	f.DescribeContinuousBackupsWithContextCall.mutex.Lock()
	defer f.DescribeContinuousBackupsWithContextCall.mutex.Unlock()
	f.DescribeContinuousBackupsWithContextCall.CallCount++
	f.DescribeContinuousBackupsWithContextCall.Receives.Context = param1
	f.DescribeContinuousBackupsWithContextCall.Receives.DescribeContinuousBackupsInput = param2
	f.DescribeContinuousBackupsWithContextCall.Receives.OptionSlice = param3
	if f.DescribeContinuousBackupsWithContextCall.Stub != nil {
		return f.DescribeContinuousBackupsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DescribeContinuousBackupsWithContextCall.Returns.DescribeContinuousBackupsOutput, f.DescribeContinuousBackupsWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeContributorInsights(param1 *dynamodb.DescribeContributorInsightsInput) (*dynamodb.DescribeContributorInsightsOutput, error) {
	f.DescribeContributorInsightsCall.mutex.Lock()
	defer f.DescribeContributorInsightsCall.mutex.Unlock()
	f.DescribeContributorInsightsCall.CallCount++
	f.DescribeContributorInsightsCall.Receives.DescribeContributorInsightsInput = param1
	if f.DescribeContributorInsightsCall.Stub != nil {
		return f.DescribeContributorInsightsCall.Stub(param1)
	}
	return f.DescribeContributorInsightsCall.Returns.DescribeContributorInsightsOutput, f.DescribeContributorInsightsCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeContributorInsightsRequest(param1 *dynamodb.DescribeContributorInsightsInput) (*request.Request, *dynamodb.DescribeContributorInsightsOutput) {
	f.DescribeContributorInsightsRequestCall.mutex.Lock()
	defer f.DescribeContributorInsightsRequestCall.mutex.Unlock()
	f.DescribeContributorInsightsRequestCall.CallCount++
	f.DescribeContributorInsightsRequestCall.Receives.DescribeContributorInsightsInput = param1
	if f.DescribeContributorInsightsRequestCall.Stub != nil {
		return f.DescribeContributorInsightsRequestCall.Stub(param1)
	}
	return f.DescribeContributorInsightsRequestCall.Returns.Request, f.DescribeContributorInsightsRequestCall.Returns.DescribeContributorInsightsOutput
}
func (f *DynamoDBAPI) DescribeContributorInsightsWithContext(param1 context.Context, param2 *dynamodb.DescribeContributorInsightsInput, param3 ...request.Option) (*dynamodb.DescribeContributorInsightsOutput, error) {
	f.DescribeContributorInsightsWithContextCall.mutex.Lock()
	defer f.DescribeContributorInsightsWithContextCall.mutex.Unlock()
	f.DescribeContributorInsightsWithContextCall.CallCount++
	f.DescribeContributorInsightsWithContextCall.Receives.Context = param1
	f.DescribeContributorInsightsWithContextCall.Receives.DescribeContributorInsightsInput = param2
	f.DescribeContributorInsightsWithContextCall.Receives.OptionSlice = param3
	if f.DescribeContributorInsightsWithContextCall.Stub != nil {
		return f.DescribeContributorInsightsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DescribeContributorInsightsWithContextCall.Returns.DescribeContributorInsightsOutput, f.DescribeContributorInsightsWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeEndpoints(param1 *dynamodb.DescribeEndpointsInput) (*dynamodb.DescribeEndpointsOutput, error) {
	f.DescribeEndpointsCall.mutex.Lock()
	defer f.DescribeEndpointsCall.mutex.Unlock()
	f.DescribeEndpointsCall.CallCount++
	f.DescribeEndpointsCall.Receives.DescribeEndpointsInput = param1
	if f.DescribeEndpointsCall.Stub != nil {
		return f.DescribeEndpointsCall.Stub(param1)
	}
	return f.DescribeEndpointsCall.Returns.DescribeEndpointsOutput, f.DescribeEndpointsCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeEndpointsRequest(param1 *dynamodb.DescribeEndpointsInput) (*request.Request, *dynamodb.DescribeEndpointsOutput) {
	f.DescribeEndpointsRequestCall.mutex.Lock()
	defer f.DescribeEndpointsRequestCall.mutex.Unlock()
	f.DescribeEndpointsRequestCall.CallCount++
	f.DescribeEndpointsRequestCall.Receives.DescribeEndpointsInput = param1
	if f.DescribeEndpointsRequestCall.Stub != nil {
		return f.DescribeEndpointsRequestCall.Stub(param1)
	}
	return f.DescribeEndpointsRequestCall.Returns.Request, f.DescribeEndpointsRequestCall.Returns.DescribeEndpointsOutput
}
func (f *DynamoDBAPI) DescribeEndpointsWithContext(param1 context.Context, param2 *dynamodb.DescribeEndpointsInput, param3 ...request.Option) (*dynamodb.DescribeEndpointsOutput, error) {
	f.DescribeEndpointsWithContextCall.mutex.Lock()
	defer f.DescribeEndpointsWithContextCall.mutex.Unlock()
	f.DescribeEndpointsWithContextCall.CallCount++
	f.DescribeEndpointsWithContextCall.Receives.Context = param1
	f.DescribeEndpointsWithContextCall.Receives.DescribeEndpointsInput = param2
	f.DescribeEndpointsWithContextCall.Receives.OptionSlice = param3
	if f.DescribeEndpointsWithContextCall.Stub != nil {
		return f.DescribeEndpointsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DescribeEndpointsWithContextCall.Returns.DescribeEndpointsOutput, f.DescribeEndpointsWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeExport(param1 *dynamodb.DescribeExportInput) (*dynamodb.DescribeExportOutput, error) {
	f.DescribeExportCall.mutex.Lock()
	defer f.DescribeExportCall.mutex.Unlock()
	f.DescribeExportCall.CallCount++
	f.DescribeExportCall.Receives.DescribeExportInput = param1
	if f.DescribeExportCall.Stub != nil {
		return f.DescribeExportCall.Stub(param1)
	}
	return f.DescribeExportCall.Returns.DescribeExportOutput, f.DescribeExportCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeExportRequest(param1 *dynamodb.DescribeExportInput) (*request.Request, *dynamodb.DescribeExportOutput) {
	f.DescribeExportRequestCall.mutex.Lock()
	defer f.DescribeExportRequestCall.mutex.Unlock()
	f.DescribeExportRequestCall.CallCount++
	f.DescribeExportRequestCall.Receives.DescribeExportInput = param1
	if f.DescribeExportRequestCall.Stub != nil {
		return f.DescribeExportRequestCall.Stub(param1)
	}
	return f.DescribeExportRequestCall.Returns.Request, f.DescribeExportRequestCall.Returns.DescribeExportOutput
}
func (f *DynamoDBAPI) DescribeExportWithContext(param1 context.Context, param2 *dynamodb.DescribeExportInput, param3 ...request.Option) (*dynamodb.DescribeExportOutput, error) {
	f.DescribeExportWithContextCall.mutex.Lock()
	defer f.DescribeExportWithContextCall.mutex.Unlock()
	f.DescribeExportWithContextCall.CallCount++
	f.DescribeExportWithContextCall.Receives.Context = param1
	f.DescribeExportWithContextCall.Receives.DescribeExportInput = param2
	f.DescribeExportWithContextCall.Receives.OptionSlice = param3
	if f.DescribeExportWithContextCall.Stub != nil {
		return f.DescribeExportWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DescribeExportWithContextCall.Returns.DescribeExportOutput, f.DescribeExportWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeGlobalTable(param1 *dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error) {
	f.DescribeGlobalTableCall.mutex.Lock()
	defer f.DescribeGlobalTableCall.mutex.Unlock()
	f.DescribeGlobalTableCall.CallCount++
	f.DescribeGlobalTableCall.Receives.DescribeGlobalTableInput = param1
	if f.DescribeGlobalTableCall.Stub != nil {
		return f.DescribeGlobalTableCall.Stub(param1)
	}
	return f.DescribeGlobalTableCall.Returns.DescribeGlobalTableOutput, f.DescribeGlobalTableCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeGlobalTableRequest(param1 *dynamodb.DescribeGlobalTableInput) (*request.Request, *dynamodb.DescribeGlobalTableOutput) {
	f.DescribeGlobalTableRequestCall.mutex.Lock()
	defer f.DescribeGlobalTableRequestCall.mutex.Unlock()
	f.DescribeGlobalTableRequestCall.CallCount++
	f.DescribeGlobalTableRequestCall.Receives.DescribeGlobalTableInput = param1
	if f.DescribeGlobalTableRequestCall.Stub != nil {
		return f.DescribeGlobalTableRequestCall.Stub(param1)
	}
	return f.DescribeGlobalTableRequestCall.Returns.Request, f.DescribeGlobalTableRequestCall.Returns.DescribeGlobalTableOutput
}
func (f *DynamoDBAPI) DescribeGlobalTableSettings(param1 *dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	f.DescribeGlobalTableSettingsCall.mutex.Lock()
	defer f.DescribeGlobalTableSettingsCall.mutex.Unlock()
	f.DescribeGlobalTableSettingsCall.CallCount++
	f.DescribeGlobalTableSettingsCall.Receives.DescribeGlobalTableSettingsInput = param1
	if f.DescribeGlobalTableSettingsCall.Stub != nil {
		return f.DescribeGlobalTableSettingsCall.Stub(param1)
	}
	return f.DescribeGlobalTableSettingsCall.Returns.DescribeGlobalTableSettingsOutput, f.DescribeGlobalTableSettingsCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeGlobalTableSettingsRequest(param1 *dynamodb.DescribeGlobalTableSettingsInput) (*request.Request, *dynamodb.DescribeGlobalTableSettingsOutput) {
	f.DescribeGlobalTableSettingsRequestCall.mutex.Lock()
	defer f.DescribeGlobalTableSettingsRequestCall.mutex.Unlock()
	f.DescribeGlobalTableSettingsRequestCall.CallCount++
	f.DescribeGlobalTableSettingsRequestCall.Receives.DescribeGlobalTableSettingsInput = param1
	if f.DescribeGlobalTableSettingsRequestCall.Stub != nil {
		return f.DescribeGlobalTableSettingsRequestCall.Stub(param1)
	}
	return f.DescribeGlobalTableSettingsRequestCall.Returns.Request, f.DescribeGlobalTableSettingsRequestCall.Returns.DescribeGlobalTableSettingsOutput
}
func (f *DynamoDBAPI) DescribeGlobalTableSettingsWithContext(param1 context.Context, param2 *dynamodb.DescribeGlobalTableSettingsInput, param3 ...request.Option) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	f.DescribeGlobalTableSettingsWithContextCall.mutex.Lock()
	defer f.DescribeGlobalTableSettingsWithContextCall.mutex.Unlock()
	f.DescribeGlobalTableSettingsWithContextCall.CallCount++
	f.DescribeGlobalTableSettingsWithContextCall.Receives.Context = param1
	f.DescribeGlobalTableSettingsWithContextCall.Receives.DescribeGlobalTableSettingsInput = param2
	f.DescribeGlobalTableSettingsWithContextCall.Receives.OptionSlice = param3
	if f.DescribeGlobalTableSettingsWithContextCall.Stub != nil {
		return f.DescribeGlobalTableSettingsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DescribeGlobalTableSettingsWithContextCall.Returns.DescribeGlobalTableSettingsOutput, f.DescribeGlobalTableSettingsWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeGlobalTableWithContext(param1 context.Context, param2 *dynamodb.DescribeGlobalTableInput, param3 ...request.Option) (*dynamodb.DescribeGlobalTableOutput, error) {
	f.DescribeGlobalTableWithContextCall.mutex.Lock()
	defer f.DescribeGlobalTableWithContextCall.mutex.Unlock()
	f.DescribeGlobalTableWithContextCall.CallCount++
	f.DescribeGlobalTableWithContextCall.Receives.Context = param1
	f.DescribeGlobalTableWithContextCall.Receives.DescribeGlobalTableInput = param2
	f.DescribeGlobalTableWithContextCall.Receives.OptionSlice = param3
	if f.DescribeGlobalTableWithContextCall.Stub != nil {
		return f.DescribeGlobalTableWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DescribeGlobalTableWithContextCall.Returns.DescribeGlobalTableOutput, f.DescribeGlobalTableWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeImport(param1 *dynamodb.DescribeImportInput) (*dynamodb.DescribeImportOutput, error) {
	f.DescribeImportCall.mutex.Lock()
	defer f.DescribeImportCall.mutex.Unlock()
	f.DescribeImportCall.CallCount++
	f.DescribeImportCall.Receives.DescribeImportInput = param1
	if f.DescribeImportCall.Stub != nil {
		return f.DescribeImportCall.Stub(param1)
	}
	return f.DescribeImportCall.Returns.DescribeImportOutput, f.DescribeImportCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeImportRequest(param1 *dynamodb.DescribeImportInput) (*request.Request, *dynamodb.DescribeImportOutput) {
	f.DescribeImportRequestCall.mutex.Lock()
	defer f.DescribeImportRequestCall.mutex.Unlock()
	f.DescribeImportRequestCall.CallCount++
	f.DescribeImportRequestCall.Receives.DescribeImportInput = param1
	if f.DescribeImportRequestCall.Stub != nil {
		return f.DescribeImportRequestCall.Stub(param1)
	}
	return f.DescribeImportRequestCall.Returns.Request, f.DescribeImportRequestCall.Returns.DescribeImportOutput
}
func (f *DynamoDBAPI) DescribeImportWithContext(param1 context.Context, param2 *dynamodb.DescribeImportInput, param3 ...request.Option) (*dynamodb.DescribeImportOutput, error) {
	f.DescribeImportWithContextCall.mutex.Lock()
	defer f.DescribeImportWithContextCall.mutex.Unlock()
	f.DescribeImportWithContextCall.CallCount++
	f.DescribeImportWithContextCall.Receives.Context = param1
	f.DescribeImportWithContextCall.Receives.DescribeImportInput = param2
	f.DescribeImportWithContextCall.Receives.OptionSlice = param3
	if f.DescribeImportWithContextCall.Stub != nil {
		return f.DescribeImportWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DescribeImportWithContextCall.Returns.DescribeImportOutput, f.DescribeImportWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeKinesisStreamingDestination(param1 *dynamodb.DescribeKinesisStreamingDestinationInput) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error) {
	f.DescribeKinesisStreamingDestinationCall.mutex.Lock()
	defer f.DescribeKinesisStreamingDestinationCall.mutex.Unlock()
	f.DescribeKinesisStreamingDestinationCall.CallCount++
	f.DescribeKinesisStreamingDestinationCall.Receives.DescribeKinesisStreamingDestinationInput = param1
	if f.DescribeKinesisStreamingDestinationCall.Stub != nil {
		return f.DescribeKinesisStreamingDestinationCall.Stub(param1)
	}
	return f.DescribeKinesisStreamingDestinationCall.Returns.DescribeKinesisStreamingDestinationOutput, f.DescribeKinesisStreamingDestinationCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeKinesisStreamingDestinationRequest(param1 *dynamodb.DescribeKinesisStreamingDestinationInput) (*request.Request, *dynamodb.DescribeKinesisStreamingDestinationOutput) {
	f.DescribeKinesisStreamingDestinationRequestCall.mutex.Lock()
	defer f.DescribeKinesisStreamingDestinationRequestCall.mutex.Unlock()
	f.DescribeKinesisStreamingDestinationRequestCall.CallCount++
	f.DescribeKinesisStreamingDestinationRequestCall.Receives.DescribeKinesisStreamingDestinationInput = param1
	if f.DescribeKinesisStreamingDestinationRequestCall.Stub != nil {
		return f.DescribeKinesisStreamingDestinationRequestCall.Stub(param1)
	}
	return f.DescribeKinesisStreamingDestinationRequestCall.Returns.Request, f.DescribeKinesisStreamingDestinationRequestCall.Returns.DescribeKinesisStreamingDestinationOutput
}
func (f *DynamoDBAPI) DescribeKinesisStreamingDestinationWithContext(param1 context.Context, param2 *dynamodb.DescribeKinesisStreamingDestinationInput, param3 ...request.Option) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error) {
	f.DescribeKinesisStreamingDestinationWithContextCall.mutex.Lock()
	defer f.DescribeKinesisStreamingDestinationWithContextCall.mutex.Unlock()
	f.DescribeKinesisStreamingDestinationWithContextCall.CallCount++
	f.DescribeKinesisStreamingDestinationWithContextCall.Receives.Context = param1
	f.DescribeKinesisStreamingDestinationWithContextCall.Receives.DescribeKinesisStreamingDestinationInput = param2
	f.DescribeKinesisStreamingDestinationWithContextCall.Receives.OptionSlice = param3
	if f.DescribeKinesisStreamingDestinationWithContextCall.Stub != nil {
		return f.DescribeKinesisStreamingDestinationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DescribeKinesisStreamingDestinationWithContextCall.Returns.DescribeKinesisStreamingDestinationOutput, f.DescribeKinesisStreamingDestinationWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeLimits(param1 *dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error) {
	f.DescribeLimitsCall.mutex.Lock()
	defer f.DescribeLimitsCall.mutex.Unlock()
	f.DescribeLimitsCall.CallCount++
	f.DescribeLimitsCall.Receives.DescribeLimitsInput = param1
	if f.DescribeLimitsCall.Stub != nil {
		return f.DescribeLimitsCall.Stub(param1)
	}
	return f.DescribeLimitsCall.Returns.DescribeLimitsOutput, f.DescribeLimitsCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeLimitsRequest(param1 *dynamodb.DescribeLimitsInput) (*request.Request, *dynamodb.DescribeLimitsOutput) {
	f.DescribeLimitsRequestCall.mutex.Lock()
	defer f.DescribeLimitsRequestCall.mutex.Unlock()
	f.DescribeLimitsRequestCall.CallCount++
	f.DescribeLimitsRequestCall.Receives.DescribeLimitsInput = param1
	if f.DescribeLimitsRequestCall.Stub != nil {
		return f.DescribeLimitsRequestCall.Stub(param1)
	}
	return f.DescribeLimitsRequestCall.Returns.Request, f.DescribeLimitsRequestCall.Returns.DescribeLimitsOutput
}
func (f *DynamoDBAPI) DescribeLimitsWithContext(param1 context.Context, param2 *dynamodb.DescribeLimitsInput, param3 ...request.Option) (*dynamodb.DescribeLimitsOutput, error) {
	f.DescribeLimitsWithContextCall.mutex.Lock()
	defer f.DescribeLimitsWithContextCall.mutex.Unlock()
	f.DescribeLimitsWithContextCall.CallCount++
	f.DescribeLimitsWithContextCall.Receives.Context = param1
	f.DescribeLimitsWithContextCall.Receives.DescribeLimitsInput = param2
	f.DescribeLimitsWithContextCall.Receives.OptionSlice = param3
	if f.DescribeLimitsWithContextCall.Stub != nil {
		return f.DescribeLimitsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DescribeLimitsWithContextCall.Returns.DescribeLimitsOutput, f.DescribeLimitsWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeTable(param1 *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
	f.DescribeTableCall.mutex.Lock()
	defer f.DescribeTableCall.mutex.Unlock()
	f.DescribeTableCall.CallCount++
	f.DescribeTableCall.Receives.DescribeTableInput = param1
	if f.DescribeTableCall.Stub != nil {
		return f.DescribeTableCall.Stub(param1)
	}
	return f.DescribeTableCall.Returns.DescribeTableOutput, f.DescribeTableCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeTableReplicaAutoScaling(param1 *dynamodb.DescribeTableReplicaAutoScalingInput) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	f.DescribeTableReplicaAutoScalingCall.mutex.Lock()
	defer f.DescribeTableReplicaAutoScalingCall.mutex.Unlock()
	f.DescribeTableReplicaAutoScalingCall.CallCount++
	f.DescribeTableReplicaAutoScalingCall.Receives.DescribeTableReplicaAutoScalingInput = param1
	if f.DescribeTableReplicaAutoScalingCall.Stub != nil {
		return f.DescribeTableReplicaAutoScalingCall.Stub(param1)
	}
	return f.DescribeTableReplicaAutoScalingCall.Returns.DescribeTableReplicaAutoScalingOutput, f.DescribeTableReplicaAutoScalingCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeTableReplicaAutoScalingRequest(param1 *dynamodb.DescribeTableReplicaAutoScalingInput) (*request.Request, *dynamodb.DescribeTableReplicaAutoScalingOutput) {
	f.DescribeTableReplicaAutoScalingRequestCall.mutex.Lock()
	defer f.DescribeTableReplicaAutoScalingRequestCall.mutex.Unlock()
	f.DescribeTableReplicaAutoScalingRequestCall.CallCount++
	f.DescribeTableReplicaAutoScalingRequestCall.Receives.DescribeTableReplicaAutoScalingInput = param1
	if f.DescribeTableReplicaAutoScalingRequestCall.Stub != nil {
		return f.DescribeTableReplicaAutoScalingRequestCall.Stub(param1)
	}
	return f.DescribeTableReplicaAutoScalingRequestCall.Returns.Request, f.DescribeTableReplicaAutoScalingRequestCall.Returns.DescribeTableReplicaAutoScalingOutput
}
func (f *DynamoDBAPI) DescribeTableReplicaAutoScalingWithContext(param1 context.Context, param2 *dynamodb.DescribeTableReplicaAutoScalingInput, param3 ...request.Option) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	f.DescribeTableReplicaAutoScalingWithContextCall.mutex.Lock()
	defer f.DescribeTableReplicaAutoScalingWithContextCall.mutex.Unlock()
	f.DescribeTableReplicaAutoScalingWithContextCall.CallCount++
	f.DescribeTableReplicaAutoScalingWithContextCall.Receives.Context = param1
	f.DescribeTableReplicaAutoScalingWithContextCall.Receives.DescribeTableReplicaAutoScalingInput = param2
	f.DescribeTableReplicaAutoScalingWithContextCall.Receives.OptionSlice = param3
	if f.DescribeTableReplicaAutoScalingWithContextCall.Stub != nil {
		return f.DescribeTableReplicaAutoScalingWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DescribeTableReplicaAutoScalingWithContextCall.Returns.DescribeTableReplicaAutoScalingOutput, f.DescribeTableReplicaAutoScalingWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeTableRequest(param1 *dynamodb.DescribeTableInput) (*request.Request, *dynamodb.DescribeTableOutput) {
	f.DescribeTableRequestCall.mutex.Lock()
	defer f.DescribeTableRequestCall.mutex.Unlock()
	f.DescribeTableRequestCall.CallCount++
	f.DescribeTableRequestCall.Receives.DescribeTableInput = param1
	if f.DescribeTableRequestCall.Stub != nil {
		return f.DescribeTableRequestCall.Stub(param1)
	}
	return f.DescribeTableRequestCall.Returns.Request, f.DescribeTableRequestCall.Returns.DescribeTableOutput
}
func (f *DynamoDBAPI) DescribeTableWithContext(param1 context.Context, param2 *dynamodb.DescribeTableInput, param3 ...request.Option) (*dynamodb.DescribeTableOutput, error) {
	f.DescribeTableWithContextCall.mutex.Lock()
	defer f.DescribeTableWithContextCall.mutex.Unlock()
	f.DescribeTableWithContextCall.CallCount++
	f.DescribeTableWithContextCall.Receives.Context = param1
	f.DescribeTableWithContextCall.Receives.DescribeTableInput = param2
	f.DescribeTableWithContextCall.Receives.OptionSlice = param3
	if f.DescribeTableWithContextCall.Stub != nil {
		return f.DescribeTableWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DescribeTableWithContextCall.Returns.DescribeTableOutput, f.DescribeTableWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeTimeToLive(param1 *dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error) {
	f.DescribeTimeToLiveCall.mutex.Lock()
	defer f.DescribeTimeToLiveCall.mutex.Unlock()
	f.DescribeTimeToLiveCall.CallCount++
	f.DescribeTimeToLiveCall.Receives.DescribeTimeToLiveInput = param1
	if f.DescribeTimeToLiveCall.Stub != nil {
		return f.DescribeTimeToLiveCall.Stub(param1)
	}
	return f.DescribeTimeToLiveCall.Returns.DescribeTimeToLiveOutput, f.DescribeTimeToLiveCall.Returns.Error
}
func (f *DynamoDBAPI) DescribeTimeToLiveRequest(param1 *dynamodb.DescribeTimeToLiveInput) (*request.Request, *dynamodb.DescribeTimeToLiveOutput) {
	f.DescribeTimeToLiveRequestCall.mutex.Lock()
	defer f.DescribeTimeToLiveRequestCall.mutex.Unlock()
	f.DescribeTimeToLiveRequestCall.CallCount++
	f.DescribeTimeToLiveRequestCall.Receives.DescribeTimeToLiveInput = param1
	if f.DescribeTimeToLiveRequestCall.Stub != nil {
		return f.DescribeTimeToLiveRequestCall.Stub(param1)
	}
	return f.DescribeTimeToLiveRequestCall.Returns.Request, f.DescribeTimeToLiveRequestCall.Returns.DescribeTimeToLiveOutput
}
func (f *DynamoDBAPI) DescribeTimeToLiveWithContext(param1 context.Context, param2 *dynamodb.DescribeTimeToLiveInput, param3 ...request.Option) (*dynamodb.DescribeTimeToLiveOutput, error) {
	f.DescribeTimeToLiveWithContextCall.mutex.Lock()
	defer f.DescribeTimeToLiveWithContextCall.mutex.Unlock()
	f.DescribeTimeToLiveWithContextCall.CallCount++
	f.DescribeTimeToLiveWithContextCall.Receives.Context = param1
	f.DescribeTimeToLiveWithContextCall.Receives.DescribeTimeToLiveInput = param2
	f.DescribeTimeToLiveWithContextCall.Receives.OptionSlice = param3
	if f.DescribeTimeToLiveWithContextCall.Stub != nil {
		return f.DescribeTimeToLiveWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DescribeTimeToLiveWithContextCall.Returns.DescribeTimeToLiveOutput, f.DescribeTimeToLiveWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) DisableKinesisStreamingDestination(param1 *dynamodb.DisableKinesisStreamingDestinationInput) (*dynamodb.DisableKinesisStreamingDestinationOutput, error) {
	f.DisableKinesisStreamingDestinationCall.mutex.Lock()
	defer f.DisableKinesisStreamingDestinationCall.mutex.Unlock()
	f.DisableKinesisStreamingDestinationCall.CallCount++
	f.DisableKinesisStreamingDestinationCall.Receives.DisableKinesisStreamingDestinationInput = param1
	if f.DisableKinesisStreamingDestinationCall.Stub != nil {
		return f.DisableKinesisStreamingDestinationCall.Stub(param1)
	}
	return f.DisableKinesisStreamingDestinationCall.Returns.DisableKinesisStreamingDestinationOutput, f.DisableKinesisStreamingDestinationCall.Returns.Error
}
func (f *DynamoDBAPI) DisableKinesisStreamingDestinationRequest(param1 *dynamodb.DisableKinesisStreamingDestinationInput) (*request.Request, *dynamodb.DisableKinesisStreamingDestinationOutput) {
	f.DisableKinesisStreamingDestinationRequestCall.mutex.Lock()
	defer f.DisableKinesisStreamingDestinationRequestCall.mutex.Unlock()
	f.DisableKinesisStreamingDestinationRequestCall.CallCount++
	f.DisableKinesisStreamingDestinationRequestCall.Receives.DisableKinesisStreamingDestinationInput = param1
	if f.DisableKinesisStreamingDestinationRequestCall.Stub != nil {
		return f.DisableKinesisStreamingDestinationRequestCall.Stub(param1)
	}
	return f.DisableKinesisStreamingDestinationRequestCall.Returns.Request, f.DisableKinesisStreamingDestinationRequestCall.Returns.DisableKinesisStreamingDestinationOutput
}
func (f *DynamoDBAPI) DisableKinesisStreamingDestinationWithContext(param1 context.Context, param2 *dynamodb.DisableKinesisStreamingDestinationInput, param3 ...request.Option) (*dynamodb.DisableKinesisStreamingDestinationOutput, error) {
	f.DisableKinesisStreamingDestinationWithContextCall.mutex.Lock()
	defer f.DisableKinesisStreamingDestinationWithContextCall.mutex.Unlock()
	f.DisableKinesisStreamingDestinationWithContextCall.CallCount++
	f.DisableKinesisStreamingDestinationWithContextCall.Receives.Context = param1
	f.DisableKinesisStreamingDestinationWithContextCall.Receives.DisableKinesisStreamingDestinationInput = param2
	f.DisableKinesisStreamingDestinationWithContextCall.Receives.OptionSlice = param3
	if f.DisableKinesisStreamingDestinationWithContextCall.Stub != nil {
		return f.DisableKinesisStreamingDestinationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DisableKinesisStreamingDestinationWithContextCall.Returns.DisableKinesisStreamingDestinationOutput, f.DisableKinesisStreamingDestinationWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) EnableKinesisStreamingDestination(param1 *dynamodb.EnableKinesisStreamingDestinationInput) (*dynamodb.EnableKinesisStreamingDestinationOutput, error) {
	f.EnableKinesisStreamingDestinationCall.mutex.Lock()
	defer f.EnableKinesisStreamingDestinationCall.mutex.Unlock()
	f.EnableKinesisStreamingDestinationCall.CallCount++
	f.EnableKinesisStreamingDestinationCall.Receives.EnableKinesisStreamingDestinationInput = param1
	if f.EnableKinesisStreamingDestinationCall.Stub != nil {
		return f.EnableKinesisStreamingDestinationCall.Stub(param1)
	}
	return f.EnableKinesisStreamingDestinationCall.Returns.EnableKinesisStreamingDestinationOutput, f.EnableKinesisStreamingDestinationCall.Returns.Error
}
func (f *DynamoDBAPI) EnableKinesisStreamingDestinationRequest(param1 *dynamodb.EnableKinesisStreamingDestinationInput) (*request.Request, *dynamodb.EnableKinesisStreamingDestinationOutput) {
	f.EnableKinesisStreamingDestinationRequestCall.mutex.Lock()
	defer f.EnableKinesisStreamingDestinationRequestCall.mutex.Unlock()
	f.EnableKinesisStreamingDestinationRequestCall.CallCount++
	f.EnableKinesisStreamingDestinationRequestCall.Receives.EnableKinesisStreamingDestinationInput = param1
	if f.EnableKinesisStreamingDestinationRequestCall.Stub != nil {
		return f.EnableKinesisStreamingDestinationRequestCall.Stub(param1)
	}
	return f.EnableKinesisStreamingDestinationRequestCall.Returns.Request, f.EnableKinesisStreamingDestinationRequestCall.Returns.EnableKinesisStreamingDestinationOutput
}
func (f *DynamoDBAPI) EnableKinesisStreamingDestinationWithContext(param1 context.Context, param2 *dynamodb.EnableKinesisStreamingDestinationInput, param3 ...request.Option) (*dynamodb.EnableKinesisStreamingDestinationOutput, error) {
	f.EnableKinesisStreamingDestinationWithContextCall.mutex.Lock()
	defer f.EnableKinesisStreamingDestinationWithContextCall.mutex.Unlock()
	f.EnableKinesisStreamingDestinationWithContextCall.CallCount++
	f.EnableKinesisStreamingDestinationWithContextCall.Receives.Context = param1
	f.EnableKinesisStreamingDestinationWithContextCall.Receives.EnableKinesisStreamingDestinationInput = param2
	f.EnableKinesisStreamingDestinationWithContextCall.Receives.OptionSlice = param3
	if f.EnableKinesisStreamingDestinationWithContextCall.Stub != nil {
		return f.EnableKinesisStreamingDestinationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.EnableKinesisStreamingDestinationWithContextCall.Returns.EnableKinesisStreamingDestinationOutput, f.EnableKinesisStreamingDestinationWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) ExecuteStatement(param1 *dynamodb.ExecuteStatementInput) (*dynamodb.ExecuteStatementOutput, error) {
	f.ExecuteStatementCall.mutex.Lock()
	defer f.ExecuteStatementCall.mutex.Unlock()
	f.ExecuteStatementCall.CallCount++
	f.ExecuteStatementCall.Receives.ExecuteStatementInput = param1
	if f.ExecuteStatementCall.Stub != nil {
		return f.ExecuteStatementCall.Stub(param1)
	}
	return f.ExecuteStatementCall.Returns.ExecuteStatementOutput, f.ExecuteStatementCall.Returns.Error
}
func (f *DynamoDBAPI) ExecuteStatementRequest(param1 *dynamodb.ExecuteStatementInput) (*request.Request, *dynamodb.ExecuteStatementOutput) {
	f.ExecuteStatementRequestCall.mutex.Lock()
	defer f.ExecuteStatementRequestCall.mutex.Unlock()
	f.ExecuteStatementRequestCall.CallCount++
	f.ExecuteStatementRequestCall.Receives.ExecuteStatementInput = param1
	if f.ExecuteStatementRequestCall.Stub != nil {
		return f.ExecuteStatementRequestCall.Stub(param1)
	}
	return f.ExecuteStatementRequestCall.Returns.Request, f.ExecuteStatementRequestCall.Returns.ExecuteStatementOutput
}
func (f *DynamoDBAPI) ExecuteStatementWithContext(param1 context.Context, param2 *dynamodb.ExecuteStatementInput, param3 ...request.Option) (*dynamodb.ExecuteStatementOutput, error) {
	f.ExecuteStatementWithContextCall.mutex.Lock()
	defer f.ExecuteStatementWithContextCall.mutex.Unlock()
	f.ExecuteStatementWithContextCall.CallCount++
	f.ExecuteStatementWithContextCall.Receives.Context = param1
	f.ExecuteStatementWithContextCall.Receives.ExecuteStatementInput = param2
	f.ExecuteStatementWithContextCall.Receives.OptionSlice = param3
	if f.ExecuteStatementWithContextCall.Stub != nil {
		return f.ExecuteStatementWithContextCall.Stub(param1, param2, param3...)
	}
	return f.ExecuteStatementWithContextCall.Returns.ExecuteStatementOutput, f.ExecuteStatementWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) ExecuteTransaction(param1 *dynamodb.ExecuteTransactionInput) (*dynamodb.ExecuteTransactionOutput, error) {
	f.ExecuteTransactionCall.mutex.Lock()
	defer f.ExecuteTransactionCall.mutex.Unlock()
	f.ExecuteTransactionCall.CallCount++
	f.ExecuteTransactionCall.Receives.ExecuteTransactionInput = param1
	if f.ExecuteTransactionCall.Stub != nil {
		return f.ExecuteTransactionCall.Stub(param1)
	}
	return f.ExecuteTransactionCall.Returns.ExecuteTransactionOutput, f.ExecuteTransactionCall.Returns.Error
}
func (f *DynamoDBAPI) ExecuteTransactionRequest(param1 *dynamodb.ExecuteTransactionInput) (*request.Request, *dynamodb.ExecuteTransactionOutput) {
	f.ExecuteTransactionRequestCall.mutex.Lock()
	defer f.ExecuteTransactionRequestCall.mutex.Unlock()
	f.ExecuteTransactionRequestCall.CallCount++
	f.ExecuteTransactionRequestCall.Receives.ExecuteTransactionInput = param1
	if f.ExecuteTransactionRequestCall.Stub != nil {
		return f.ExecuteTransactionRequestCall.Stub(param1)
	}
	return f.ExecuteTransactionRequestCall.Returns.Request, f.ExecuteTransactionRequestCall.Returns.ExecuteTransactionOutput
}
func (f *DynamoDBAPI) ExecuteTransactionWithContext(param1 context.Context, param2 *dynamodb.ExecuteTransactionInput, param3 ...request.Option) (*dynamodb.ExecuteTransactionOutput, error) {
	f.ExecuteTransactionWithContextCall.mutex.Lock()
	defer f.ExecuteTransactionWithContextCall.mutex.Unlock()
	f.ExecuteTransactionWithContextCall.CallCount++
	f.ExecuteTransactionWithContextCall.Receives.Context = param1
	f.ExecuteTransactionWithContextCall.Receives.ExecuteTransactionInput = param2
	f.ExecuteTransactionWithContextCall.Receives.OptionSlice = param3
	if f.ExecuteTransactionWithContextCall.Stub != nil {
		return f.ExecuteTransactionWithContextCall.Stub(param1, param2, param3...)
	}
	return f.ExecuteTransactionWithContextCall.Returns.ExecuteTransactionOutput, f.ExecuteTransactionWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) ExportTableToPointInTime(param1 *dynamodb.ExportTableToPointInTimeInput) (*dynamodb.ExportTableToPointInTimeOutput, error) {
	f.ExportTableToPointInTimeCall.mutex.Lock()
	defer f.ExportTableToPointInTimeCall.mutex.Unlock()
	f.ExportTableToPointInTimeCall.CallCount++
	f.ExportTableToPointInTimeCall.Receives.ExportTableToPointInTimeInput = param1
	if f.ExportTableToPointInTimeCall.Stub != nil {
		return f.ExportTableToPointInTimeCall.Stub(param1)
	}
	return f.ExportTableToPointInTimeCall.Returns.ExportTableToPointInTimeOutput, f.ExportTableToPointInTimeCall.Returns.Error
}
func (f *DynamoDBAPI) ExportTableToPointInTimeRequest(param1 *dynamodb.ExportTableToPointInTimeInput) (*request.Request, *dynamodb.ExportTableToPointInTimeOutput) {
	f.ExportTableToPointInTimeRequestCall.mutex.Lock()
	defer f.ExportTableToPointInTimeRequestCall.mutex.Unlock()
	f.ExportTableToPointInTimeRequestCall.CallCount++
	f.ExportTableToPointInTimeRequestCall.Receives.ExportTableToPointInTimeInput = param1
	if f.ExportTableToPointInTimeRequestCall.Stub != nil {
		return f.ExportTableToPointInTimeRequestCall.Stub(param1)
	}
	return f.ExportTableToPointInTimeRequestCall.Returns.Request, f.ExportTableToPointInTimeRequestCall.Returns.ExportTableToPointInTimeOutput
}
func (f *DynamoDBAPI) ExportTableToPointInTimeWithContext(param1 context.Context, param2 *dynamodb.ExportTableToPointInTimeInput, param3 ...request.Option) (*dynamodb.ExportTableToPointInTimeOutput, error) {
	f.ExportTableToPointInTimeWithContextCall.mutex.Lock()
	defer f.ExportTableToPointInTimeWithContextCall.mutex.Unlock()
	f.ExportTableToPointInTimeWithContextCall.CallCount++
	f.ExportTableToPointInTimeWithContextCall.Receives.Context = param1
	f.ExportTableToPointInTimeWithContextCall.Receives.ExportTableToPointInTimeInput = param2
	f.ExportTableToPointInTimeWithContextCall.Receives.OptionSlice = param3
	if f.ExportTableToPointInTimeWithContextCall.Stub != nil {
		return f.ExportTableToPointInTimeWithContextCall.Stub(param1, param2, param3...)
	}
	return f.ExportTableToPointInTimeWithContextCall.Returns.ExportTableToPointInTimeOutput, f.ExportTableToPointInTimeWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) GetItem(param1 *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	f.GetItemCall.mutex.Lock()
	defer f.GetItemCall.mutex.Unlock()
	f.GetItemCall.CallCount++
	f.GetItemCall.Receives.GetItemInput = param1
	if f.GetItemCall.Stub != nil {
		return f.GetItemCall.Stub(param1)
	}
	return f.GetItemCall.Returns.GetItemOutput, f.GetItemCall.Returns.Error
}
func (f *DynamoDBAPI) GetItemRequest(param1 *dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput) {
	f.GetItemRequestCall.mutex.Lock()
	defer f.GetItemRequestCall.mutex.Unlock()
	f.GetItemRequestCall.CallCount++
	f.GetItemRequestCall.Receives.GetItemInput = param1
	if f.GetItemRequestCall.Stub != nil {
		return f.GetItemRequestCall.Stub(param1)
	}
	return f.GetItemRequestCall.Returns.Request, f.GetItemRequestCall.Returns.GetItemOutput
}
func (f *DynamoDBAPI) GetItemWithContext(param1 context.Context, param2 *dynamodb.GetItemInput, param3 ...request.Option) (*dynamodb.GetItemOutput, error) {
	f.GetItemWithContextCall.mutex.Lock()
	defer f.GetItemWithContextCall.mutex.Unlock()
	f.GetItemWithContextCall.CallCount++
	f.GetItemWithContextCall.Receives.Context = param1
	f.GetItemWithContextCall.Receives.GetItemInput = param2
	f.GetItemWithContextCall.Receives.OptionSlice = param3
	if f.GetItemWithContextCall.Stub != nil {
		return f.GetItemWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetItemWithContextCall.Returns.GetItemOutput, f.GetItemWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) ImportTable(param1 *dynamodb.ImportTableInput) (*dynamodb.ImportTableOutput, error) {
	f.ImportTableCall.mutex.Lock()
	defer f.ImportTableCall.mutex.Unlock()
	f.ImportTableCall.CallCount++
	f.ImportTableCall.Receives.ImportTableInput = param1
	if f.ImportTableCall.Stub != nil {
		return f.ImportTableCall.Stub(param1)
	}
	return f.ImportTableCall.Returns.ImportTableOutput, f.ImportTableCall.Returns.Error
}
func (f *DynamoDBAPI) ImportTableRequest(param1 *dynamodb.ImportTableInput) (*request.Request, *dynamodb.ImportTableOutput) {
	f.ImportTableRequestCall.mutex.Lock()
	defer f.ImportTableRequestCall.mutex.Unlock()
	f.ImportTableRequestCall.CallCount++
	f.ImportTableRequestCall.Receives.ImportTableInput = param1
	if f.ImportTableRequestCall.Stub != nil {
		return f.ImportTableRequestCall.Stub(param1)
	}
	return f.ImportTableRequestCall.Returns.Request, f.ImportTableRequestCall.Returns.ImportTableOutput
}
func (f *DynamoDBAPI) ImportTableWithContext(param1 context.Context, param2 *dynamodb.ImportTableInput, param3 ...request.Option) (*dynamodb.ImportTableOutput, error) {
	f.ImportTableWithContextCall.mutex.Lock()
	defer f.ImportTableWithContextCall.mutex.Unlock()
	f.ImportTableWithContextCall.CallCount++
	f.ImportTableWithContextCall.Receives.Context = param1
	f.ImportTableWithContextCall.Receives.ImportTableInput = param2
	f.ImportTableWithContextCall.Receives.OptionSlice = param3
	if f.ImportTableWithContextCall.Stub != nil {
		return f.ImportTableWithContextCall.Stub(param1, param2, param3...)
	}
	return f.ImportTableWithContextCall.Returns.ImportTableOutput, f.ImportTableWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) ListBackups(param1 *dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error) {
	f.ListBackupsCall.mutex.Lock()
	defer f.ListBackupsCall.mutex.Unlock()
	f.ListBackupsCall.CallCount++
	f.ListBackupsCall.Receives.ListBackupsInput = param1
	if f.ListBackupsCall.Stub != nil {
		return f.ListBackupsCall.Stub(param1)
	}
	return f.ListBackupsCall.Returns.ListBackupsOutput, f.ListBackupsCall.Returns.Error
}
func (f *DynamoDBAPI) ListBackupsRequest(param1 *dynamodb.ListBackupsInput) (*request.Request, *dynamodb.ListBackupsOutput) {
	f.ListBackupsRequestCall.mutex.Lock()
	defer f.ListBackupsRequestCall.mutex.Unlock()
	f.ListBackupsRequestCall.CallCount++
	f.ListBackupsRequestCall.Receives.ListBackupsInput = param1
	if f.ListBackupsRequestCall.Stub != nil {
		return f.ListBackupsRequestCall.Stub(param1)
	}
	return f.ListBackupsRequestCall.Returns.Request, f.ListBackupsRequestCall.Returns.ListBackupsOutput
}
func (f *DynamoDBAPI) ListBackupsWithContext(param1 context.Context, param2 *dynamodb.ListBackupsInput, param3 ...request.Option) (*dynamodb.ListBackupsOutput, error) {
	f.ListBackupsWithContextCall.mutex.Lock()
	defer f.ListBackupsWithContextCall.mutex.Unlock()
	f.ListBackupsWithContextCall.CallCount++
	f.ListBackupsWithContextCall.Receives.Context = param1
	f.ListBackupsWithContextCall.Receives.ListBackupsInput = param2
	f.ListBackupsWithContextCall.Receives.OptionSlice = param3
	if f.ListBackupsWithContextCall.Stub != nil {
		return f.ListBackupsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.ListBackupsWithContextCall.Returns.ListBackupsOutput, f.ListBackupsWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) ListContributorInsights(param1 *dynamodb.ListContributorInsightsInput) (*dynamodb.ListContributorInsightsOutput, error) {
	f.ListContributorInsightsCall.mutex.Lock()
	defer f.ListContributorInsightsCall.mutex.Unlock()
	f.ListContributorInsightsCall.CallCount++
	f.ListContributorInsightsCall.Receives.ListContributorInsightsInput = param1
	if f.ListContributorInsightsCall.Stub != nil {
		return f.ListContributorInsightsCall.Stub(param1)
	}
	return f.ListContributorInsightsCall.Returns.ListContributorInsightsOutput, f.ListContributorInsightsCall.Returns.Error
}
func (f *DynamoDBAPI) ListContributorInsightsPages(param1 *dynamodb.ListContributorInsightsInput, param2 func(*dynamodb.ListContributorInsightsOutput, bool) bool) error {
	f.ListContributorInsightsPagesCall.mutex.Lock()
	defer f.ListContributorInsightsPagesCall.mutex.Unlock()
	f.ListContributorInsightsPagesCall.CallCount++
	f.ListContributorInsightsPagesCall.Receives.ListContributorInsightsInput = param1
	f.ListContributorInsightsPagesCall.Receives.FuncDynamodbListContributorInsightsOutputBoolBool = param2
	if f.ListContributorInsightsPagesCall.Stub != nil {
		return f.ListContributorInsightsPagesCall.Stub(param1, param2)
	}
	return f.ListContributorInsightsPagesCall.Returns.Error
}
func (f *DynamoDBAPI) ListContributorInsightsPagesWithContext(param1 context.Context, param2 *dynamodb.ListContributorInsightsInput, param3 func(*dynamodb.ListContributorInsightsOutput, bool) bool, param4 ...request.Option) error {
	f.ListContributorInsightsPagesWithContextCall.mutex.Lock()
	defer f.ListContributorInsightsPagesWithContextCall.mutex.Unlock()
	f.ListContributorInsightsPagesWithContextCall.CallCount++
	f.ListContributorInsightsPagesWithContextCall.Receives.Context = param1
	f.ListContributorInsightsPagesWithContextCall.Receives.ListContributorInsightsInput = param2
	f.ListContributorInsightsPagesWithContextCall.Receives.FuncDynamodbListContributorInsightsOutputBoolBool = param3
	f.ListContributorInsightsPagesWithContextCall.Receives.OptionSlice = param4
	if f.ListContributorInsightsPagesWithContextCall.Stub != nil {
		return f.ListContributorInsightsPagesWithContextCall.Stub(param1, param2, param3, param4...)
	}
	return f.ListContributorInsightsPagesWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) ListContributorInsightsRequest(param1 *dynamodb.ListContributorInsightsInput) (*request.Request, *dynamodb.ListContributorInsightsOutput) {
	f.ListContributorInsightsRequestCall.mutex.Lock()
	defer f.ListContributorInsightsRequestCall.mutex.Unlock()
	f.ListContributorInsightsRequestCall.CallCount++
	f.ListContributorInsightsRequestCall.Receives.ListContributorInsightsInput = param1
	if f.ListContributorInsightsRequestCall.Stub != nil {
		return f.ListContributorInsightsRequestCall.Stub(param1)
	}
	return f.ListContributorInsightsRequestCall.Returns.Request, f.ListContributorInsightsRequestCall.Returns.ListContributorInsightsOutput
}
func (f *DynamoDBAPI) ListContributorInsightsWithContext(param1 context.Context, param2 *dynamodb.ListContributorInsightsInput, param3 ...request.Option) (*dynamodb.ListContributorInsightsOutput, error) {
	f.ListContributorInsightsWithContextCall.mutex.Lock()
	defer f.ListContributorInsightsWithContextCall.mutex.Unlock()
	f.ListContributorInsightsWithContextCall.CallCount++
	f.ListContributorInsightsWithContextCall.Receives.Context = param1
	f.ListContributorInsightsWithContextCall.Receives.ListContributorInsightsInput = param2
	f.ListContributorInsightsWithContextCall.Receives.OptionSlice = param3
	if f.ListContributorInsightsWithContextCall.Stub != nil {
		return f.ListContributorInsightsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.ListContributorInsightsWithContextCall.Returns.ListContributorInsightsOutput, f.ListContributorInsightsWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) ListExports(param1 *dynamodb.ListExportsInput) (*dynamodb.ListExportsOutput, error) {
	f.ListExportsCall.mutex.Lock()
	defer f.ListExportsCall.mutex.Unlock()
	f.ListExportsCall.CallCount++
	f.ListExportsCall.Receives.ListExportsInput = param1
	if f.ListExportsCall.Stub != nil {
		return f.ListExportsCall.Stub(param1)
	}
	return f.ListExportsCall.Returns.ListExportsOutput, f.ListExportsCall.Returns.Error
}
func (f *DynamoDBAPI) ListExportsPages(param1 *dynamodb.ListExportsInput, param2 func(*dynamodb.ListExportsOutput, bool) bool) error {
	f.ListExportsPagesCall.mutex.Lock()
	defer f.ListExportsPagesCall.mutex.Unlock()
	f.ListExportsPagesCall.CallCount++
	f.ListExportsPagesCall.Receives.ListExportsInput = param1
	f.ListExportsPagesCall.Receives.FuncDynamodbListExportsOutputBoolBool = param2
	if f.ListExportsPagesCall.Stub != nil {
		return f.ListExportsPagesCall.Stub(param1, param2)
	}
	return f.ListExportsPagesCall.Returns.Error
}
func (f *DynamoDBAPI) ListExportsPagesWithContext(param1 context.Context, param2 *dynamodb.ListExportsInput, param3 func(*dynamodb.ListExportsOutput, bool) bool, param4 ...request.Option) error {
	f.ListExportsPagesWithContextCall.mutex.Lock()
	defer f.ListExportsPagesWithContextCall.mutex.Unlock()
	f.ListExportsPagesWithContextCall.CallCount++
	f.ListExportsPagesWithContextCall.Receives.Context = param1
	f.ListExportsPagesWithContextCall.Receives.ListExportsInput = param2
	f.ListExportsPagesWithContextCall.Receives.FuncDynamodbListExportsOutputBoolBool = param3
	f.ListExportsPagesWithContextCall.Receives.OptionSlice = param4
	if f.ListExportsPagesWithContextCall.Stub != nil {
		return f.ListExportsPagesWithContextCall.Stub(param1, param2, param3, param4...)
	}
	return f.ListExportsPagesWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) ListExportsRequest(param1 *dynamodb.ListExportsInput) (*request.Request, *dynamodb.ListExportsOutput) {
	f.ListExportsRequestCall.mutex.Lock()
	defer f.ListExportsRequestCall.mutex.Unlock()
	f.ListExportsRequestCall.CallCount++
	f.ListExportsRequestCall.Receives.ListExportsInput = param1
	if f.ListExportsRequestCall.Stub != nil {
		return f.ListExportsRequestCall.Stub(param1)
	}
	return f.ListExportsRequestCall.Returns.Request, f.ListExportsRequestCall.Returns.ListExportsOutput
}
func (f *DynamoDBAPI) ListExportsWithContext(param1 context.Context, param2 *dynamodb.ListExportsInput, param3 ...request.Option) (*dynamodb.ListExportsOutput, error) {
	f.ListExportsWithContextCall.mutex.Lock()
	defer f.ListExportsWithContextCall.mutex.Unlock()
	f.ListExportsWithContextCall.CallCount++
	f.ListExportsWithContextCall.Receives.Context = param1
	f.ListExportsWithContextCall.Receives.ListExportsInput = param2
	f.ListExportsWithContextCall.Receives.OptionSlice = param3
	if f.ListExportsWithContextCall.Stub != nil {
		return f.ListExportsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.ListExportsWithContextCall.Returns.ListExportsOutput, f.ListExportsWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) ListGlobalTables(param1 *dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error) {
	f.ListGlobalTablesCall.mutex.Lock()
	defer f.ListGlobalTablesCall.mutex.Unlock()
	f.ListGlobalTablesCall.CallCount++
	f.ListGlobalTablesCall.Receives.ListGlobalTablesInput = param1
	if f.ListGlobalTablesCall.Stub != nil {
		return f.ListGlobalTablesCall.Stub(param1)
	}
	return f.ListGlobalTablesCall.Returns.ListGlobalTablesOutput, f.ListGlobalTablesCall.Returns.Error
}
func (f *DynamoDBAPI) ListGlobalTablesRequest(param1 *dynamodb.ListGlobalTablesInput) (*request.Request, *dynamodb.ListGlobalTablesOutput) {
	f.ListGlobalTablesRequestCall.mutex.Lock()
	defer f.ListGlobalTablesRequestCall.mutex.Unlock()
	f.ListGlobalTablesRequestCall.CallCount++
	f.ListGlobalTablesRequestCall.Receives.ListGlobalTablesInput = param1
	if f.ListGlobalTablesRequestCall.Stub != nil {
		return f.ListGlobalTablesRequestCall.Stub(param1)
	}
	return f.ListGlobalTablesRequestCall.Returns.Request, f.ListGlobalTablesRequestCall.Returns.ListGlobalTablesOutput
}
func (f *DynamoDBAPI) ListGlobalTablesWithContext(param1 context.Context, param2 *dynamodb.ListGlobalTablesInput, param3 ...request.Option) (*dynamodb.ListGlobalTablesOutput, error) {
	f.ListGlobalTablesWithContextCall.mutex.Lock()
	defer f.ListGlobalTablesWithContextCall.mutex.Unlock()
	f.ListGlobalTablesWithContextCall.CallCount++
	f.ListGlobalTablesWithContextCall.Receives.Context = param1
	f.ListGlobalTablesWithContextCall.Receives.ListGlobalTablesInput = param2
	f.ListGlobalTablesWithContextCall.Receives.OptionSlice = param3
	if f.ListGlobalTablesWithContextCall.Stub != nil {
		return f.ListGlobalTablesWithContextCall.Stub(param1, param2, param3...)
	}
	return f.ListGlobalTablesWithContextCall.Returns.ListGlobalTablesOutput, f.ListGlobalTablesWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) ListImports(param1 *dynamodb.ListImportsInput) (*dynamodb.ListImportsOutput, error) {
	f.ListImportsCall.mutex.Lock()
	defer f.ListImportsCall.mutex.Unlock()
	f.ListImportsCall.CallCount++
	f.ListImportsCall.Receives.ListImportsInput = param1
	if f.ListImportsCall.Stub != nil {
		return f.ListImportsCall.Stub(param1)
	}
	return f.ListImportsCall.Returns.ListImportsOutput, f.ListImportsCall.Returns.Error
}
func (f *DynamoDBAPI) ListImportsPages(param1 *dynamodb.ListImportsInput, param2 func(*dynamodb.ListImportsOutput, bool) bool) error {
	f.ListImportsPagesCall.mutex.Lock()
	defer f.ListImportsPagesCall.mutex.Unlock()
	f.ListImportsPagesCall.CallCount++
	f.ListImportsPagesCall.Receives.ListImportsInput = param1
	f.ListImportsPagesCall.Receives.FuncDynamodbListImportsOutputBoolBool = param2
	if f.ListImportsPagesCall.Stub != nil {
		return f.ListImportsPagesCall.Stub(param1, param2)
	}
	return f.ListImportsPagesCall.Returns.Error
}
func (f *DynamoDBAPI) ListImportsPagesWithContext(param1 context.Context, param2 *dynamodb.ListImportsInput, param3 func(*dynamodb.ListImportsOutput, bool) bool, param4 ...request.Option) error {
	f.ListImportsPagesWithContextCall.mutex.Lock()
	defer f.ListImportsPagesWithContextCall.mutex.Unlock()
	f.ListImportsPagesWithContextCall.CallCount++
	f.ListImportsPagesWithContextCall.Receives.Context = param1
	f.ListImportsPagesWithContextCall.Receives.ListImportsInput = param2
	f.ListImportsPagesWithContextCall.Receives.FuncDynamodbListImportsOutputBoolBool = param3
	f.ListImportsPagesWithContextCall.Receives.OptionSlice = param4
	if f.ListImportsPagesWithContextCall.Stub != nil {
		return f.ListImportsPagesWithContextCall.Stub(param1, param2, param3, param4...)
	}
	return f.ListImportsPagesWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) ListImportsRequest(param1 *dynamodb.ListImportsInput) (*request.Request, *dynamodb.ListImportsOutput) {
	f.ListImportsRequestCall.mutex.Lock()
	defer f.ListImportsRequestCall.mutex.Unlock()
	f.ListImportsRequestCall.CallCount++
	f.ListImportsRequestCall.Receives.ListImportsInput = param1
	if f.ListImportsRequestCall.Stub != nil {
		return f.ListImportsRequestCall.Stub(param1)
	}
	return f.ListImportsRequestCall.Returns.Request, f.ListImportsRequestCall.Returns.ListImportsOutput
}
func (f *DynamoDBAPI) ListImportsWithContext(param1 context.Context, param2 *dynamodb.ListImportsInput, param3 ...request.Option) (*dynamodb.ListImportsOutput, error) {
	f.ListImportsWithContextCall.mutex.Lock()
	defer f.ListImportsWithContextCall.mutex.Unlock()
	f.ListImportsWithContextCall.CallCount++
	f.ListImportsWithContextCall.Receives.Context = param1
	f.ListImportsWithContextCall.Receives.ListImportsInput = param2
	f.ListImportsWithContextCall.Receives.OptionSlice = param3
	if f.ListImportsWithContextCall.Stub != nil {
		return f.ListImportsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.ListImportsWithContextCall.Returns.ListImportsOutput, f.ListImportsWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) ListTables(param1 *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
	f.ListTablesCall.mutex.Lock()
	defer f.ListTablesCall.mutex.Unlock()
	f.ListTablesCall.CallCount++
	f.ListTablesCall.Receives.ListTablesInput = param1
	if f.ListTablesCall.Stub != nil {
		return f.ListTablesCall.Stub(param1)
	}
	return f.ListTablesCall.Returns.ListTablesOutput, f.ListTablesCall.Returns.Error
}
func (f *DynamoDBAPI) ListTablesPages(param1 *dynamodb.ListTablesInput, param2 func(*dynamodb.ListTablesOutput, bool) bool) error {
	f.ListTablesPagesCall.mutex.Lock()
	defer f.ListTablesPagesCall.mutex.Unlock()
	f.ListTablesPagesCall.CallCount++
	f.ListTablesPagesCall.Receives.ListTablesInput = param1
	f.ListTablesPagesCall.Receives.FuncDynamodbListTablesOutputBoolBool = param2
	if f.ListTablesPagesCall.Stub != nil {
		return f.ListTablesPagesCall.Stub(param1, param2)
	}
	return f.ListTablesPagesCall.Returns.Error
}
func (f *DynamoDBAPI) ListTablesPagesWithContext(param1 context.Context, param2 *dynamodb.ListTablesInput, param3 func(*dynamodb.ListTablesOutput, bool) bool, param4 ...request.Option) error {
	f.ListTablesPagesWithContextCall.mutex.Lock()
	defer f.ListTablesPagesWithContextCall.mutex.Unlock()
	f.ListTablesPagesWithContextCall.CallCount++
	f.ListTablesPagesWithContextCall.Receives.Context = param1
	f.ListTablesPagesWithContextCall.Receives.ListTablesInput = param2
	f.ListTablesPagesWithContextCall.Receives.FuncDynamodbListTablesOutputBoolBool = param3
	f.ListTablesPagesWithContextCall.Receives.OptionSlice = param4
	if f.ListTablesPagesWithContextCall.Stub != nil {
		return f.ListTablesPagesWithContextCall.Stub(param1, param2, param3, param4...)
	}
	return f.ListTablesPagesWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) ListTablesRequest(param1 *dynamodb.ListTablesInput) (*request.Request, *dynamodb.ListTablesOutput) {
	f.ListTablesRequestCall.mutex.Lock()
	defer f.ListTablesRequestCall.mutex.Unlock()
	f.ListTablesRequestCall.CallCount++
	f.ListTablesRequestCall.Receives.ListTablesInput = param1
	if f.ListTablesRequestCall.Stub != nil {
		return f.ListTablesRequestCall.Stub(param1)
	}
	return f.ListTablesRequestCall.Returns.Request, f.ListTablesRequestCall.Returns.ListTablesOutput
}
func (f *DynamoDBAPI) ListTablesWithContext(param1 context.Context, param2 *dynamodb.ListTablesInput, param3 ...request.Option) (*dynamodb.ListTablesOutput, error) {
	f.ListTablesWithContextCall.mutex.Lock()
	defer f.ListTablesWithContextCall.mutex.Unlock()
	f.ListTablesWithContextCall.CallCount++
	f.ListTablesWithContextCall.Receives.Context = param1
	f.ListTablesWithContextCall.Receives.ListTablesInput = param2
	f.ListTablesWithContextCall.Receives.OptionSlice = param3
	if f.ListTablesWithContextCall.Stub != nil {
		return f.ListTablesWithContextCall.Stub(param1, param2, param3...)
	}
	return f.ListTablesWithContextCall.Returns.ListTablesOutput, f.ListTablesWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) ListTagsOfResource(param1 *dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error) {
	f.ListTagsOfResourceCall.mutex.Lock()
	defer f.ListTagsOfResourceCall.mutex.Unlock()
	f.ListTagsOfResourceCall.CallCount++
	f.ListTagsOfResourceCall.Receives.ListTagsOfResourceInput = param1
	if f.ListTagsOfResourceCall.Stub != nil {
		return f.ListTagsOfResourceCall.Stub(param1)
	}
	return f.ListTagsOfResourceCall.Returns.ListTagsOfResourceOutput, f.ListTagsOfResourceCall.Returns.Error
}
func (f *DynamoDBAPI) ListTagsOfResourceRequest(param1 *dynamodb.ListTagsOfResourceInput) (*request.Request, *dynamodb.ListTagsOfResourceOutput) {
	f.ListTagsOfResourceRequestCall.mutex.Lock()
	defer f.ListTagsOfResourceRequestCall.mutex.Unlock()
	f.ListTagsOfResourceRequestCall.CallCount++
	f.ListTagsOfResourceRequestCall.Receives.ListTagsOfResourceInput = param1
	if f.ListTagsOfResourceRequestCall.Stub != nil {
		return f.ListTagsOfResourceRequestCall.Stub(param1)
	}
	return f.ListTagsOfResourceRequestCall.Returns.Request, f.ListTagsOfResourceRequestCall.Returns.ListTagsOfResourceOutput
}
func (f *DynamoDBAPI) ListTagsOfResourceWithContext(param1 context.Context, param2 *dynamodb.ListTagsOfResourceInput, param3 ...request.Option) (*dynamodb.ListTagsOfResourceOutput, error) {
	f.ListTagsOfResourceWithContextCall.mutex.Lock()
	defer f.ListTagsOfResourceWithContextCall.mutex.Unlock()
	f.ListTagsOfResourceWithContextCall.CallCount++
	f.ListTagsOfResourceWithContextCall.Receives.Context = param1
	f.ListTagsOfResourceWithContextCall.Receives.ListTagsOfResourceInput = param2
	f.ListTagsOfResourceWithContextCall.Receives.OptionSlice = param3
	if f.ListTagsOfResourceWithContextCall.Stub != nil {
		return f.ListTagsOfResourceWithContextCall.Stub(param1, param2, param3...)
	}
	return f.ListTagsOfResourceWithContextCall.Returns.ListTagsOfResourceOutput, f.ListTagsOfResourceWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) PutItem(param1 *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	f.PutItemCall.mutex.Lock()
	defer f.PutItemCall.mutex.Unlock()
	f.PutItemCall.CallCount++
	f.PutItemCall.Receives.PutItemInput = param1
	if f.PutItemCall.Stub != nil {
		return f.PutItemCall.Stub(param1)
	}
	return f.PutItemCall.Returns.PutItemOutput, f.PutItemCall.Returns.Error
}
func (f *DynamoDBAPI) PutItemRequest(param1 *dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput) {
	f.PutItemRequestCall.mutex.Lock()
	defer f.PutItemRequestCall.mutex.Unlock()
	f.PutItemRequestCall.CallCount++
	f.PutItemRequestCall.Receives.PutItemInput = param1
	if f.PutItemRequestCall.Stub != nil {
		return f.PutItemRequestCall.Stub(param1)
	}
	return f.PutItemRequestCall.Returns.Request, f.PutItemRequestCall.Returns.PutItemOutput
}
func (f *DynamoDBAPI) PutItemWithContext(param1 context.Context, param2 *dynamodb.PutItemInput, param3 ...request.Option) (*dynamodb.PutItemOutput, error) {
	f.PutItemWithContextCall.mutex.Lock()
	defer f.PutItemWithContextCall.mutex.Unlock()
	f.PutItemWithContextCall.CallCount++
	f.PutItemWithContextCall.Receives.Context = param1
	f.PutItemWithContextCall.Receives.PutItemInput = param2
	f.PutItemWithContextCall.Receives.OptionSlice = param3
	if f.PutItemWithContextCall.Stub != nil {
		return f.PutItemWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutItemWithContextCall.Returns.PutItemOutput, f.PutItemWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) Query(param1 *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	f.QueryCall.mutex.Lock()
	defer f.QueryCall.mutex.Unlock()
	f.QueryCall.CallCount++
	f.QueryCall.Receives.QueryInput = param1
	if f.QueryCall.Stub != nil {
		return f.QueryCall.Stub(param1)
	}
	return f.QueryCall.Returns.QueryOutput, f.QueryCall.Returns.Error
}
func (f *DynamoDBAPI) QueryPages(param1 *dynamodb.QueryInput, param2 func(*dynamodb.QueryOutput, bool) bool) error {
	f.QueryPagesCall.mutex.Lock()
	defer f.QueryPagesCall.mutex.Unlock()
	f.QueryPagesCall.CallCount++
	f.QueryPagesCall.Receives.QueryInput = param1
	f.QueryPagesCall.Receives.FuncDynamodbQueryOutputBoolBool = param2
	if f.QueryPagesCall.Stub != nil {
		return f.QueryPagesCall.Stub(param1, param2)
	}
	return f.QueryPagesCall.Returns.Error
}
func (f *DynamoDBAPI) QueryPagesWithContext(param1 context.Context, param2 *dynamodb.QueryInput, param3 func(*dynamodb.QueryOutput, bool) bool, param4 ...request.Option) error {
	f.QueryPagesWithContextCall.mutex.Lock()
	defer f.QueryPagesWithContextCall.mutex.Unlock()
	f.QueryPagesWithContextCall.CallCount++
	f.QueryPagesWithContextCall.Receives.Context = param1
	f.QueryPagesWithContextCall.Receives.QueryInput = param2
	f.QueryPagesWithContextCall.Receives.FuncDynamodbQueryOutputBoolBool = param3
	f.QueryPagesWithContextCall.Receives.OptionSlice = param4
	if f.QueryPagesWithContextCall.Stub != nil {
		return f.QueryPagesWithContextCall.Stub(param1, param2, param3, param4...)
	}
	return f.QueryPagesWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) QueryRequest(param1 *dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput) {
	f.QueryRequestCall.mutex.Lock()
	defer f.QueryRequestCall.mutex.Unlock()
	f.QueryRequestCall.CallCount++
	f.QueryRequestCall.Receives.QueryInput = param1
	if f.QueryRequestCall.Stub != nil {
		return f.QueryRequestCall.Stub(param1)
	}
	return f.QueryRequestCall.Returns.Request, f.QueryRequestCall.Returns.QueryOutput
}
func (f *DynamoDBAPI) QueryWithContext(param1 context.Context, param2 *dynamodb.QueryInput, param3 ...request.Option) (*dynamodb.QueryOutput, error) {
	f.QueryWithContextCall.mutex.Lock()
	defer f.QueryWithContextCall.mutex.Unlock()
	f.QueryWithContextCall.CallCount++
	f.QueryWithContextCall.Receives.Context = param1
	f.QueryWithContextCall.Receives.QueryInput = param2
	f.QueryWithContextCall.Receives.OptionSlice = param3
	if f.QueryWithContextCall.Stub != nil {
		return f.QueryWithContextCall.Stub(param1, param2, param3...)
	}
	return f.QueryWithContextCall.Returns.QueryOutput, f.QueryWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) RestoreTableFromBackup(param1 *dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error) {
	f.RestoreTableFromBackupCall.mutex.Lock()
	defer f.RestoreTableFromBackupCall.mutex.Unlock()
	f.RestoreTableFromBackupCall.CallCount++
	f.RestoreTableFromBackupCall.Receives.RestoreTableFromBackupInput = param1
	if f.RestoreTableFromBackupCall.Stub != nil {
		return f.RestoreTableFromBackupCall.Stub(param1)
	}
	return f.RestoreTableFromBackupCall.Returns.RestoreTableFromBackupOutput, f.RestoreTableFromBackupCall.Returns.Error
}
func (f *DynamoDBAPI) RestoreTableFromBackupRequest(param1 *dynamodb.RestoreTableFromBackupInput) (*request.Request, *dynamodb.RestoreTableFromBackupOutput) {
	f.RestoreTableFromBackupRequestCall.mutex.Lock()
	defer f.RestoreTableFromBackupRequestCall.mutex.Unlock()
	f.RestoreTableFromBackupRequestCall.CallCount++
	f.RestoreTableFromBackupRequestCall.Receives.RestoreTableFromBackupInput = param1
	if f.RestoreTableFromBackupRequestCall.Stub != nil {
		return f.RestoreTableFromBackupRequestCall.Stub(param1)
	}
	return f.RestoreTableFromBackupRequestCall.Returns.Request, f.RestoreTableFromBackupRequestCall.Returns.RestoreTableFromBackupOutput
}
func (f *DynamoDBAPI) RestoreTableFromBackupWithContext(param1 context.Context, param2 *dynamodb.RestoreTableFromBackupInput, param3 ...request.Option) (*dynamodb.RestoreTableFromBackupOutput, error) {
	f.RestoreTableFromBackupWithContextCall.mutex.Lock()
	defer f.RestoreTableFromBackupWithContextCall.mutex.Unlock()
	f.RestoreTableFromBackupWithContextCall.CallCount++
	f.RestoreTableFromBackupWithContextCall.Receives.Context = param1
	f.RestoreTableFromBackupWithContextCall.Receives.RestoreTableFromBackupInput = param2
	f.RestoreTableFromBackupWithContextCall.Receives.OptionSlice = param3
	if f.RestoreTableFromBackupWithContextCall.Stub != nil {
		return f.RestoreTableFromBackupWithContextCall.Stub(param1, param2, param3...)
	}
	return f.RestoreTableFromBackupWithContextCall.Returns.RestoreTableFromBackupOutput, f.RestoreTableFromBackupWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) RestoreTableToPointInTime(param1 *dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	f.RestoreTableToPointInTimeCall.mutex.Lock()
	defer f.RestoreTableToPointInTimeCall.mutex.Unlock()
	f.RestoreTableToPointInTimeCall.CallCount++
	f.RestoreTableToPointInTimeCall.Receives.RestoreTableToPointInTimeInput = param1
	if f.RestoreTableToPointInTimeCall.Stub != nil {
		return f.RestoreTableToPointInTimeCall.Stub(param1)
	}
	return f.RestoreTableToPointInTimeCall.Returns.RestoreTableToPointInTimeOutput, f.RestoreTableToPointInTimeCall.Returns.Error
}
func (f *DynamoDBAPI) RestoreTableToPointInTimeRequest(param1 *dynamodb.RestoreTableToPointInTimeInput) (*request.Request, *dynamodb.RestoreTableToPointInTimeOutput) {
	f.RestoreTableToPointInTimeRequestCall.mutex.Lock()
	defer f.RestoreTableToPointInTimeRequestCall.mutex.Unlock()
	f.RestoreTableToPointInTimeRequestCall.CallCount++
	f.RestoreTableToPointInTimeRequestCall.Receives.RestoreTableToPointInTimeInput = param1
	if f.RestoreTableToPointInTimeRequestCall.Stub != nil {
		return f.RestoreTableToPointInTimeRequestCall.Stub(param1)
	}
	return f.RestoreTableToPointInTimeRequestCall.Returns.Request, f.RestoreTableToPointInTimeRequestCall.Returns.RestoreTableToPointInTimeOutput
}
func (f *DynamoDBAPI) RestoreTableToPointInTimeWithContext(param1 context.Context, param2 *dynamodb.RestoreTableToPointInTimeInput, param3 ...request.Option) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	f.RestoreTableToPointInTimeWithContextCall.mutex.Lock()
	defer f.RestoreTableToPointInTimeWithContextCall.mutex.Unlock()
	f.RestoreTableToPointInTimeWithContextCall.CallCount++
	f.RestoreTableToPointInTimeWithContextCall.Receives.Context = param1
	f.RestoreTableToPointInTimeWithContextCall.Receives.RestoreTableToPointInTimeInput = param2
	f.RestoreTableToPointInTimeWithContextCall.Receives.OptionSlice = param3
	if f.RestoreTableToPointInTimeWithContextCall.Stub != nil {
		return f.RestoreTableToPointInTimeWithContextCall.Stub(param1, param2, param3...)
	}
	return f.RestoreTableToPointInTimeWithContextCall.Returns.RestoreTableToPointInTimeOutput, f.RestoreTableToPointInTimeWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) Scan(param1 *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	f.ScanCall.mutex.Lock()
	defer f.ScanCall.mutex.Unlock()
	f.ScanCall.CallCount++
	f.ScanCall.Receives.ScanInput = param1
	if f.ScanCall.Stub != nil {
		return f.ScanCall.Stub(param1)
	}
	return f.ScanCall.Returns.ScanOutput, f.ScanCall.Returns.Error
}
func (f *DynamoDBAPI) ScanPages(param1 *dynamodb.ScanInput, param2 func(*dynamodb.ScanOutput, bool) bool) error {
	f.ScanPagesCall.mutex.Lock()
	defer f.ScanPagesCall.mutex.Unlock()
	f.ScanPagesCall.CallCount++
	f.ScanPagesCall.Receives.ScanInput = param1
	f.ScanPagesCall.Receives.FuncDynamodbScanOutputBoolBool = param2
	if f.ScanPagesCall.Stub != nil {
		return f.ScanPagesCall.Stub(param1, param2)
	}
	return f.ScanPagesCall.Returns.Error
}
func (f *DynamoDBAPI) ScanPagesWithContext(param1 context.Context, param2 *dynamodb.ScanInput, param3 func(*dynamodb.ScanOutput, bool) bool, param4 ...request.Option) error {
	f.ScanPagesWithContextCall.mutex.Lock()
	defer f.ScanPagesWithContextCall.mutex.Unlock()
	f.ScanPagesWithContextCall.CallCount++
	f.ScanPagesWithContextCall.Receives.Context = param1
	f.ScanPagesWithContextCall.Receives.ScanInput = param2
	f.ScanPagesWithContextCall.Receives.FuncDynamodbScanOutputBoolBool = param3
	f.ScanPagesWithContextCall.Receives.OptionSlice = param4
	if f.ScanPagesWithContextCall.Stub != nil {
		return f.ScanPagesWithContextCall.Stub(param1, param2, param3, param4...)
	}
	return f.ScanPagesWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) ScanRequest(param1 *dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput) {
	f.ScanRequestCall.mutex.Lock()
	defer f.ScanRequestCall.mutex.Unlock()
	f.ScanRequestCall.CallCount++
	f.ScanRequestCall.Receives.ScanInput = param1
	if f.ScanRequestCall.Stub != nil {
		return f.ScanRequestCall.Stub(param1)
	}
	return f.ScanRequestCall.Returns.Request, f.ScanRequestCall.Returns.ScanOutput
}
func (f *DynamoDBAPI) ScanWithContext(param1 context.Context, param2 *dynamodb.ScanInput, param3 ...request.Option) (*dynamodb.ScanOutput, error) {
	f.ScanWithContextCall.mutex.Lock()
	defer f.ScanWithContextCall.mutex.Unlock()
	f.ScanWithContextCall.CallCount++
	f.ScanWithContextCall.Receives.Context = param1
	f.ScanWithContextCall.Receives.ScanInput = param2
	f.ScanWithContextCall.Receives.OptionSlice = param3
	if f.ScanWithContextCall.Stub != nil {
		return f.ScanWithContextCall.Stub(param1, param2, param3...)
	}
	return f.ScanWithContextCall.Returns.ScanOutput, f.ScanWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) TagResource(param1 *dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error) {
	f.TagResourceCall.mutex.Lock()
	defer f.TagResourceCall.mutex.Unlock()
	f.TagResourceCall.CallCount++
	f.TagResourceCall.Receives.TagResourceInput = param1
	if f.TagResourceCall.Stub != nil {
		return f.TagResourceCall.Stub(param1)
	}
	return f.TagResourceCall.Returns.TagResourceOutput, f.TagResourceCall.Returns.Error
}
func (f *DynamoDBAPI) TagResourceRequest(param1 *dynamodb.TagResourceInput) (*request.Request, *dynamodb.TagResourceOutput) {
	f.TagResourceRequestCall.mutex.Lock()
	defer f.TagResourceRequestCall.mutex.Unlock()
	f.TagResourceRequestCall.CallCount++
	f.TagResourceRequestCall.Receives.TagResourceInput = param1
	if f.TagResourceRequestCall.Stub != nil {
		return f.TagResourceRequestCall.Stub(param1)
	}
	return f.TagResourceRequestCall.Returns.Request, f.TagResourceRequestCall.Returns.TagResourceOutput
}
func (f *DynamoDBAPI) TagResourceWithContext(param1 context.Context, param2 *dynamodb.TagResourceInput, param3 ...request.Option) (*dynamodb.TagResourceOutput, error) {
	f.TagResourceWithContextCall.mutex.Lock()
	defer f.TagResourceWithContextCall.mutex.Unlock()
	f.TagResourceWithContextCall.CallCount++
	f.TagResourceWithContextCall.Receives.Context = param1
	f.TagResourceWithContextCall.Receives.TagResourceInput = param2
	f.TagResourceWithContextCall.Receives.OptionSlice = param3
	if f.TagResourceWithContextCall.Stub != nil {
		return f.TagResourceWithContextCall.Stub(param1, param2, param3...)
	}
	return f.TagResourceWithContextCall.Returns.TagResourceOutput, f.TagResourceWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) TransactGetItems(param1 *dynamodb.TransactGetItemsInput) (*dynamodb.TransactGetItemsOutput, error) {
	f.TransactGetItemsCall.mutex.Lock()
	defer f.TransactGetItemsCall.mutex.Unlock()
	f.TransactGetItemsCall.CallCount++
	f.TransactGetItemsCall.Receives.TransactGetItemsInput = param1
	if f.TransactGetItemsCall.Stub != nil {
		return f.TransactGetItemsCall.Stub(param1)
	}
	return f.TransactGetItemsCall.Returns.TransactGetItemsOutput, f.TransactGetItemsCall.Returns.Error
}
func (f *DynamoDBAPI) TransactGetItemsRequest(param1 *dynamodb.TransactGetItemsInput) (*request.Request, *dynamodb.TransactGetItemsOutput) {
	f.TransactGetItemsRequestCall.mutex.Lock()
	defer f.TransactGetItemsRequestCall.mutex.Unlock()
	f.TransactGetItemsRequestCall.CallCount++
	f.TransactGetItemsRequestCall.Receives.TransactGetItemsInput = param1
	if f.TransactGetItemsRequestCall.Stub != nil {
		return f.TransactGetItemsRequestCall.Stub(param1)
	}
	return f.TransactGetItemsRequestCall.Returns.Request, f.TransactGetItemsRequestCall.Returns.TransactGetItemsOutput
}
func (f *DynamoDBAPI) TransactGetItemsWithContext(param1 context.Context, param2 *dynamodb.TransactGetItemsInput, param3 ...request.Option) (*dynamodb.TransactGetItemsOutput, error) {
	f.TransactGetItemsWithContextCall.mutex.Lock()
	defer f.TransactGetItemsWithContextCall.mutex.Unlock()
	f.TransactGetItemsWithContextCall.CallCount++
	f.TransactGetItemsWithContextCall.Receives.Context = param1
	f.TransactGetItemsWithContextCall.Receives.TransactGetItemsInput = param2
	f.TransactGetItemsWithContextCall.Receives.OptionSlice = param3
	if f.TransactGetItemsWithContextCall.Stub != nil {
		return f.TransactGetItemsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.TransactGetItemsWithContextCall.Returns.TransactGetItemsOutput, f.TransactGetItemsWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) TransactWriteItems(param1 *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error) {
	f.TransactWriteItemsCall.mutex.Lock()
	defer f.TransactWriteItemsCall.mutex.Unlock()
	f.TransactWriteItemsCall.CallCount++
	f.TransactWriteItemsCall.Receives.TransactWriteItemsInput = param1
	if f.TransactWriteItemsCall.Stub != nil {
		return f.TransactWriteItemsCall.Stub(param1)
	}
	return f.TransactWriteItemsCall.Returns.TransactWriteItemsOutput, f.TransactWriteItemsCall.Returns.Error
}
func (f *DynamoDBAPI) TransactWriteItemsRequest(param1 *dynamodb.TransactWriteItemsInput) (*request.Request, *dynamodb.TransactWriteItemsOutput) {
	f.TransactWriteItemsRequestCall.mutex.Lock()
	defer f.TransactWriteItemsRequestCall.mutex.Unlock()
	f.TransactWriteItemsRequestCall.CallCount++
	f.TransactWriteItemsRequestCall.Receives.TransactWriteItemsInput = param1
	if f.TransactWriteItemsRequestCall.Stub != nil {
		return f.TransactWriteItemsRequestCall.Stub(param1)
	}
	return f.TransactWriteItemsRequestCall.Returns.Request, f.TransactWriteItemsRequestCall.Returns.TransactWriteItemsOutput
}
func (f *DynamoDBAPI) TransactWriteItemsWithContext(param1 context.Context, param2 *dynamodb.TransactWriteItemsInput, param3 ...request.Option) (*dynamodb.TransactWriteItemsOutput, error) {
	f.TransactWriteItemsWithContextCall.mutex.Lock()
	defer f.TransactWriteItemsWithContextCall.mutex.Unlock()
	f.TransactWriteItemsWithContextCall.CallCount++
	f.TransactWriteItemsWithContextCall.Receives.Context = param1
	f.TransactWriteItemsWithContextCall.Receives.TransactWriteItemsInput = param2
	f.TransactWriteItemsWithContextCall.Receives.OptionSlice = param3
	if f.TransactWriteItemsWithContextCall.Stub != nil {
		return f.TransactWriteItemsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.TransactWriteItemsWithContextCall.Returns.TransactWriteItemsOutput, f.TransactWriteItemsWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) UntagResource(param1 *dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error) {
	f.UntagResourceCall.mutex.Lock()
	defer f.UntagResourceCall.mutex.Unlock()
	f.UntagResourceCall.CallCount++
	f.UntagResourceCall.Receives.UntagResourceInput = param1
	if f.UntagResourceCall.Stub != nil {
		return f.UntagResourceCall.Stub(param1)
	}
	return f.UntagResourceCall.Returns.UntagResourceOutput, f.UntagResourceCall.Returns.Error
}
func (f *DynamoDBAPI) UntagResourceRequest(param1 *dynamodb.UntagResourceInput) (*request.Request, *dynamodb.UntagResourceOutput) {
	f.UntagResourceRequestCall.mutex.Lock()
	defer f.UntagResourceRequestCall.mutex.Unlock()
	f.UntagResourceRequestCall.CallCount++
	f.UntagResourceRequestCall.Receives.UntagResourceInput = param1
	if f.UntagResourceRequestCall.Stub != nil {
		return f.UntagResourceRequestCall.Stub(param1)
	}
	return f.UntagResourceRequestCall.Returns.Request, f.UntagResourceRequestCall.Returns.UntagResourceOutput
}
func (f *DynamoDBAPI) UntagResourceWithContext(param1 context.Context, param2 *dynamodb.UntagResourceInput, param3 ...request.Option) (*dynamodb.UntagResourceOutput, error) {
	f.UntagResourceWithContextCall.mutex.Lock()
	defer f.UntagResourceWithContextCall.mutex.Unlock()
	f.UntagResourceWithContextCall.CallCount++
	f.UntagResourceWithContextCall.Receives.Context = param1
	f.UntagResourceWithContextCall.Receives.UntagResourceInput = param2
	f.UntagResourceWithContextCall.Receives.OptionSlice = param3
	if f.UntagResourceWithContextCall.Stub != nil {
		return f.UntagResourceWithContextCall.Stub(param1, param2, param3...)
	}
	return f.UntagResourceWithContextCall.Returns.UntagResourceOutput, f.UntagResourceWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) UpdateContinuousBackups(param1 *dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	f.UpdateContinuousBackupsCall.mutex.Lock()
	defer f.UpdateContinuousBackupsCall.mutex.Unlock()
	f.UpdateContinuousBackupsCall.CallCount++
	f.UpdateContinuousBackupsCall.Receives.UpdateContinuousBackupsInput = param1
	if f.UpdateContinuousBackupsCall.Stub != nil {
		return f.UpdateContinuousBackupsCall.Stub(param1)
	}
	return f.UpdateContinuousBackupsCall.Returns.UpdateContinuousBackupsOutput, f.UpdateContinuousBackupsCall.Returns.Error
}
func (f *DynamoDBAPI) UpdateContinuousBackupsRequest(param1 *dynamodb.UpdateContinuousBackupsInput) (*request.Request, *dynamodb.UpdateContinuousBackupsOutput) {
	f.UpdateContinuousBackupsRequestCall.mutex.Lock()
	defer f.UpdateContinuousBackupsRequestCall.mutex.Unlock()
	f.UpdateContinuousBackupsRequestCall.CallCount++
	f.UpdateContinuousBackupsRequestCall.Receives.UpdateContinuousBackupsInput = param1
	if f.UpdateContinuousBackupsRequestCall.Stub != nil {
		return f.UpdateContinuousBackupsRequestCall.Stub(param1)
	}
	return f.UpdateContinuousBackupsRequestCall.Returns.Request, f.UpdateContinuousBackupsRequestCall.Returns.UpdateContinuousBackupsOutput
}
func (f *DynamoDBAPI) UpdateContinuousBackupsWithContext(param1 context.Context, param2 *dynamodb.UpdateContinuousBackupsInput, param3 ...request.Option) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	f.UpdateContinuousBackupsWithContextCall.mutex.Lock()
	defer f.UpdateContinuousBackupsWithContextCall.mutex.Unlock()
	f.UpdateContinuousBackupsWithContextCall.CallCount++
	f.UpdateContinuousBackupsWithContextCall.Receives.Context = param1
	f.UpdateContinuousBackupsWithContextCall.Receives.UpdateContinuousBackupsInput = param2
	f.UpdateContinuousBackupsWithContextCall.Receives.OptionSlice = param3
	if f.UpdateContinuousBackupsWithContextCall.Stub != nil {
		return f.UpdateContinuousBackupsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.UpdateContinuousBackupsWithContextCall.Returns.UpdateContinuousBackupsOutput, f.UpdateContinuousBackupsWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) UpdateContributorInsights(param1 *dynamodb.UpdateContributorInsightsInput) (*dynamodb.UpdateContributorInsightsOutput, error) {
	f.UpdateContributorInsightsCall.mutex.Lock()
	defer f.UpdateContributorInsightsCall.mutex.Unlock()
	f.UpdateContributorInsightsCall.CallCount++
	f.UpdateContributorInsightsCall.Receives.UpdateContributorInsightsInput = param1
	if f.UpdateContributorInsightsCall.Stub != nil {
		return f.UpdateContributorInsightsCall.Stub(param1)
	}
	return f.UpdateContributorInsightsCall.Returns.UpdateContributorInsightsOutput, f.UpdateContributorInsightsCall.Returns.Error
}
func (f *DynamoDBAPI) UpdateContributorInsightsRequest(param1 *dynamodb.UpdateContributorInsightsInput) (*request.Request, *dynamodb.UpdateContributorInsightsOutput) {
	f.UpdateContributorInsightsRequestCall.mutex.Lock()
	defer f.UpdateContributorInsightsRequestCall.mutex.Unlock()
	f.UpdateContributorInsightsRequestCall.CallCount++
	f.UpdateContributorInsightsRequestCall.Receives.UpdateContributorInsightsInput = param1
	if f.UpdateContributorInsightsRequestCall.Stub != nil {
		return f.UpdateContributorInsightsRequestCall.Stub(param1)
	}
	return f.UpdateContributorInsightsRequestCall.Returns.Request, f.UpdateContributorInsightsRequestCall.Returns.UpdateContributorInsightsOutput
}
func (f *DynamoDBAPI) UpdateContributorInsightsWithContext(param1 context.Context, param2 *dynamodb.UpdateContributorInsightsInput, param3 ...request.Option) (*dynamodb.UpdateContributorInsightsOutput, error) {
	f.UpdateContributorInsightsWithContextCall.mutex.Lock()
	defer f.UpdateContributorInsightsWithContextCall.mutex.Unlock()
	f.UpdateContributorInsightsWithContextCall.CallCount++
	f.UpdateContributorInsightsWithContextCall.Receives.Context = param1
	f.UpdateContributorInsightsWithContextCall.Receives.UpdateContributorInsightsInput = param2
	f.UpdateContributorInsightsWithContextCall.Receives.OptionSlice = param3
	if f.UpdateContributorInsightsWithContextCall.Stub != nil {
		return f.UpdateContributorInsightsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.UpdateContributorInsightsWithContextCall.Returns.UpdateContributorInsightsOutput, f.UpdateContributorInsightsWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) UpdateGlobalTable(param1 *dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error) {
	f.UpdateGlobalTableCall.mutex.Lock()
	defer f.UpdateGlobalTableCall.mutex.Unlock()
	f.UpdateGlobalTableCall.CallCount++
	f.UpdateGlobalTableCall.Receives.UpdateGlobalTableInput = param1
	if f.UpdateGlobalTableCall.Stub != nil {
		return f.UpdateGlobalTableCall.Stub(param1)
	}
	return f.UpdateGlobalTableCall.Returns.UpdateGlobalTableOutput, f.UpdateGlobalTableCall.Returns.Error
}
func (f *DynamoDBAPI) UpdateGlobalTableRequest(param1 *dynamodb.UpdateGlobalTableInput) (*request.Request, *dynamodb.UpdateGlobalTableOutput) {
	f.UpdateGlobalTableRequestCall.mutex.Lock()
	defer f.UpdateGlobalTableRequestCall.mutex.Unlock()
	f.UpdateGlobalTableRequestCall.CallCount++
	f.UpdateGlobalTableRequestCall.Receives.UpdateGlobalTableInput = param1
	if f.UpdateGlobalTableRequestCall.Stub != nil {
		return f.UpdateGlobalTableRequestCall.Stub(param1)
	}
	return f.UpdateGlobalTableRequestCall.Returns.Request, f.UpdateGlobalTableRequestCall.Returns.UpdateGlobalTableOutput
}
func (f *DynamoDBAPI) UpdateGlobalTableSettings(param1 *dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	f.UpdateGlobalTableSettingsCall.mutex.Lock()
	defer f.UpdateGlobalTableSettingsCall.mutex.Unlock()
	f.UpdateGlobalTableSettingsCall.CallCount++
	f.UpdateGlobalTableSettingsCall.Receives.UpdateGlobalTableSettingsInput = param1
	if f.UpdateGlobalTableSettingsCall.Stub != nil {
		return f.UpdateGlobalTableSettingsCall.Stub(param1)
	}
	return f.UpdateGlobalTableSettingsCall.Returns.UpdateGlobalTableSettingsOutput, f.UpdateGlobalTableSettingsCall.Returns.Error
}
func (f *DynamoDBAPI) UpdateGlobalTableSettingsRequest(param1 *dynamodb.UpdateGlobalTableSettingsInput) (*request.Request, *dynamodb.UpdateGlobalTableSettingsOutput) {
	f.UpdateGlobalTableSettingsRequestCall.mutex.Lock()
	defer f.UpdateGlobalTableSettingsRequestCall.mutex.Unlock()
	f.UpdateGlobalTableSettingsRequestCall.CallCount++
	f.UpdateGlobalTableSettingsRequestCall.Receives.UpdateGlobalTableSettingsInput = param1
	if f.UpdateGlobalTableSettingsRequestCall.Stub != nil {
		return f.UpdateGlobalTableSettingsRequestCall.Stub(param1)
	}
	return f.UpdateGlobalTableSettingsRequestCall.Returns.Request, f.UpdateGlobalTableSettingsRequestCall.Returns.UpdateGlobalTableSettingsOutput
}
func (f *DynamoDBAPI) UpdateGlobalTableSettingsWithContext(param1 context.Context, param2 *dynamodb.UpdateGlobalTableSettingsInput, param3 ...request.Option) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	f.UpdateGlobalTableSettingsWithContextCall.mutex.Lock()
	defer f.UpdateGlobalTableSettingsWithContextCall.mutex.Unlock()
	f.UpdateGlobalTableSettingsWithContextCall.CallCount++
	f.UpdateGlobalTableSettingsWithContextCall.Receives.Context = param1
	f.UpdateGlobalTableSettingsWithContextCall.Receives.UpdateGlobalTableSettingsInput = param2
	f.UpdateGlobalTableSettingsWithContextCall.Receives.OptionSlice = param3
	if f.UpdateGlobalTableSettingsWithContextCall.Stub != nil {
		return f.UpdateGlobalTableSettingsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.UpdateGlobalTableSettingsWithContextCall.Returns.UpdateGlobalTableSettingsOutput, f.UpdateGlobalTableSettingsWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) UpdateGlobalTableWithContext(param1 context.Context, param2 *dynamodb.UpdateGlobalTableInput, param3 ...request.Option) (*dynamodb.UpdateGlobalTableOutput, error) {
	f.UpdateGlobalTableWithContextCall.mutex.Lock()
	defer f.UpdateGlobalTableWithContextCall.mutex.Unlock()
	f.UpdateGlobalTableWithContextCall.CallCount++
	f.UpdateGlobalTableWithContextCall.Receives.Context = param1
	f.UpdateGlobalTableWithContextCall.Receives.UpdateGlobalTableInput = param2
	f.UpdateGlobalTableWithContextCall.Receives.OptionSlice = param3
	if f.UpdateGlobalTableWithContextCall.Stub != nil {
		return f.UpdateGlobalTableWithContextCall.Stub(param1, param2, param3...)
	}
	return f.UpdateGlobalTableWithContextCall.Returns.UpdateGlobalTableOutput, f.UpdateGlobalTableWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) UpdateItem(param1 *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	f.UpdateItemCall.mutex.Lock()
	defer f.UpdateItemCall.mutex.Unlock()
	f.UpdateItemCall.CallCount++
	f.UpdateItemCall.Receives.UpdateItemInput = param1
	if f.UpdateItemCall.Stub != nil {
		return f.UpdateItemCall.Stub(param1)
	}
	return f.UpdateItemCall.Returns.UpdateItemOutput, f.UpdateItemCall.Returns.Error
}
func (f *DynamoDBAPI) UpdateItemRequest(param1 *dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput) {
	f.UpdateItemRequestCall.mutex.Lock()
	defer f.UpdateItemRequestCall.mutex.Unlock()
	f.UpdateItemRequestCall.CallCount++
	f.UpdateItemRequestCall.Receives.UpdateItemInput = param1
	if f.UpdateItemRequestCall.Stub != nil {
		return f.UpdateItemRequestCall.Stub(param1)
	}
	return f.UpdateItemRequestCall.Returns.Request, f.UpdateItemRequestCall.Returns.UpdateItemOutput
}
func (f *DynamoDBAPI) UpdateItemWithContext(param1 context.Context, param2 *dynamodb.UpdateItemInput, param3 ...request.Option) (*dynamodb.UpdateItemOutput, error) {
	f.UpdateItemWithContextCall.mutex.Lock()
	defer f.UpdateItemWithContextCall.mutex.Unlock()
	f.UpdateItemWithContextCall.CallCount++
	f.UpdateItemWithContextCall.Receives.Context = param1
	f.UpdateItemWithContextCall.Receives.UpdateItemInput = param2
	f.UpdateItemWithContextCall.Receives.OptionSlice = param3
	if f.UpdateItemWithContextCall.Stub != nil {
		return f.UpdateItemWithContextCall.Stub(param1, param2, param3...)
	}
	return f.UpdateItemWithContextCall.Returns.UpdateItemOutput, f.UpdateItemWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) UpdateTable(param1 *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
	f.UpdateTableCall.mutex.Lock()
	defer f.UpdateTableCall.mutex.Unlock()
	f.UpdateTableCall.CallCount++
	f.UpdateTableCall.Receives.UpdateTableInput = param1
	if f.UpdateTableCall.Stub != nil {
		return f.UpdateTableCall.Stub(param1)
	}
	return f.UpdateTableCall.Returns.UpdateTableOutput, f.UpdateTableCall.Returns.Error
}
func (f *DynamoDBAPI) UpdateTableReplicaAutoScaling(param1 *dynamodb.UpdateTableReplicaAutoScalingInput) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	f.UpdateTableReplicaAutoScalingCall.mutex.Lock()
	defer f.UpdateTableReplicaAutoScalingCall.mutex.Unlock()
	f.UpdateTableReplicaAutoScalingCall.CallCount++
	f.UpdateTableReplicaAutoScalingCall.Receives.UpdateTableReplicaAutoScalingInput = param1
	if f.UpdateTableReplicaAutoScalingCall.Stub != nil {
		return f.UpdateTableReplicaAutoScalingCall.Stub(param1)
	}
	return f.UpdateTableReplicaAutoScalingCall.Returns.UpdateTableReplicaAutoScalingOutput, f.UpdateTableReplicaAutoScalingCall.Returns.Error
}
func (f *DynamoDBAPI) UpdateTableReplicaAutoScalingRequest(param1 *dynamodb.UpdateTableReplicaAutoScalingInput) (*request.Request, *dynamodb.UpdateTableReplicaAutoScalingOutput) {
	f.UpdateTableReplicaAutoScalingRequestCall.mutex.Lock()
	defer f.UpdateTableReplicaAutoScalingRequestCall.mutex.Unlock()
	f.UpdateTableReplicaAutoScalingRequestCall.CallCount++
	f.UpdateTableReplicaAutoScalingRequestCall.Receives.UpdateTableReplicaAutoScalingInput = param1
	if f.UpdateTableReplicaAutoScalingRequestCall.Stub != nil {
		return f.UpdateTableReplicaAutoScalingRequestCall.Stub(param1)
	}
	return f.UpdateTableReplicaAutoScalingRequestCall.Returns.Request, f.UpdateTableReplicaAutoScalingRequestCall.Returns.UpdateTableReplicaAutoScalingOutput
}
func (f *DynamoDBAPI) UpdateTableReplicaAutoScalingWithContext(param1 context.Context, param2 *dynamodb.UpdateTableReplicaAutoScalingInput, param3 ...request.Option) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	f.UpdateTableReplicaAutoScalingWithContextCall.mutex.Lock()
	defer f.UpdateTableReplicaAutoScalingWithContextCall.mutex.Unlock()
	f.UpdateTableReplicaAutoScalingWithContextCall.CallCount++
	f.UpdateTableReplicaAutoScalingWithContextCall.Receives.Context = param1
	f.UpdateTableReplicaAutoScalingWithContextCall.Receives.UpdateTableReplicaAutoScalingInput = param2
	f.UpdateTableReplicaAutoScalingWithContextCall.Receives.OptionSlice = param3
	if f.UpdateTableReplicaAutoScalingWithContextCall.Stub != nil {
		return f.UpdateTableReplicaAutoScalingWithContextCall.Stub(param1, param2, param3...)
	}
	return f.UpdateTableReplicaAutoScalingWithContextCall.Returns.UpdateTableReplicaAutoScalingOutput, f.UpdateTableReplicaAutoScalingWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) UpdateTableRequest(param1 *dynamodb.UpdateTableInput) (*request.Request, *dynamodb.UpdateTableOutput) {
	f.UpdateTableRequestCall.mutex.Lock()
	defer f.UpdateTableRequestCall.mutex.Unlock()
	f.UpdateTableRequestCall.CallCount++
	f.UpdateTableRequestCall.Receives.UpdateTableInput = param1
	if f.UpdateTableRequestCall.Stub != nil {
		return f.UpdateTableRequestCall.Stub(param1)
	}
	return f.UpdateTableRequestCall.Returns.Request, f.UpdateTableRequestCall.Returns.UpdateTableOutput
}
func (f *DynamoDBAPI) UpdateTableWithContext(param1 context.Context, param2 *dynamodb.UpdateTableInput, param3 ...request.Option) (*dynamodb.UpdateTableOutput, error) {
	f.UpdateTableWithContextCall.mutex.Lock()
	defer f.UpdateTableWithContextCall.mutex.Unlock()
	f.UpdateTableWithContextCall.CallCount++
	f.UpdateTableWithContextCall.Receives.Context = param1
	f.UpdateTableWithContextCall.Receives.UpdateTableInput = param2
	f.UpdateTableWithContextCall.Receives.OptionSlice = param3
	if f.UpdateTableWithContextCall.Stub != nil {
		return f.UpdateTableWithContextCall.Stub(param1, param2, param3...)
	}
	return f.UpdateTableWithContextCall.Returns.UpdateTableOutput, f.UpdateTableWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) UpdateTimeToLive(param1 *dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error) {
	f.UpdateTimeToLiveCall.mutex.Lock()
	defer f.UpdateTimeToLiveCall.mutex.Unlock()
	f.UpdateTimeToLiveCall.CallCount++
	f.UpdateTimeToLiveCall.Receives.UpdateTimeToLiveInput = param1
	if f.UpdateTimeToLiveCall.Stub != nil {
		return f.UpdateTimeToLiveCall.Stub(param1)
	}
	return f.UpdateTimeToLiveCall.Returns.UpdateTimeToLiveOutput, f.UpdateTimeToLiveCall.Returns.Error
}
func (f *DynamoDBAPI) UpdateTimeToLiveRequest(param1 *dynamodb.UpdateTimeToLiveInput) (*request.Request, *dynamodb.UpdateTimeToLiveOutput) {
	f.UpdateTimeToLiveRequestCall.mutex.Lock()
	defer f.UpdateTimeToLiveRequestCall.mutex.Unlock()
	f.UpdateTimeToLiveRequestCall.CallCount++
	f.UpdateTimeToLiveRequestCall.Receives.UpdateTimeToLiveInput = param1
	if f.UpdateTimeToLiveRequestCall.Stub != nil {
		return f.UpdateTimeToLiveRequestCall.Stub(param1)
	}
	return f.UpdateTimeToLiveRequestCall.Returns.Request, f.UpdateTimeToLiveRequestCall.Returns.UpdateTimeToLiveOutput
}
func (f *DynamoDBAPI) UpdateTimeToLiveWithContext(param1 context.Context, param2 *dynamodb.UpdateTimeToLiveInput, param3 ...request.Option) (*dynamodb.UpdateTimeToLiveOutput, error) {
	f.UpdateTimeToLiveWithContextCall.mutex.Lock()
	defer f.UpdateTimeToLiveWithContextCall.mutex.Unlock()
	f.UpdateTimeToLiveWithContextCall.CallCount++
	f.UpdateTimeToLiveWithContextCall.Receives.Context = param1
	f.UpdateTimeToLiveWithContextCall.Receives.UpdateTimeToLiveInput = param2
	f.UpdateTimeToLiveWithContextCall.Receives.OptionSlice = param3
	if f.UpdateTimeToLiveWithContextCall.Stub != nil {
		return f.UpdateTimeToLiveWithContextCall.Stub(param1, param2, param3...)
	}
	return f.UpdateTimeToLiveWithContextCall.Returns.UpdateTimeToLiveOutput, f.UpdateTimeToLiveWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) WaitUntilTableExists(param1 *dynamodb.DescribeTableInput) error {
	f.WaitUntilTableExistsCall.mutex.Lock()
	defer f.WaitUntilTableExistsCall.mutex.Unlock()
	f.WaitUntilTableExistsCall.CallCount++
	f.WaitUntilTableExistsCall.Receives.DescribeTableInput = param1
	if f.WaitUntilTableExistsCall.Stub != nil {
		return f.WaitUntilTableExistsCall.Stub(param1)
	}
	return f.WaitUntilTableExistsCall.Returns.Error
}
func (f *DynamoDBAPI) WaitUntilTableExistsWithContext(param1 context.Context, param2 *dynamodb.DescribeTableInput, param3 ...request.WaiterOption) error {
	f.WaitUntilTableExistsWithContextCall.mutex.Lock()
	defer f.WaitUntilTableExistsWithContextCall.mutex.Unlock()
	f.WaitUntilTableExistsWithContextCall.CallCount++
	f.WaitUntilTableExistsWithContextCall.Receives.Context = param1
	f.WaitUntilTableExistsWithContextCall.Receives.DescribeTableInput = param2
	f.WaitUntilTableExistsWithContextCall.Receives.WaiterOptionSlice = param3
	if f.WaitUntilTableExistsWithContextCall.Stub != nil {
		return f.WaitUntilTableExistsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.WaitUntilTableExistsWithContextCall.Returns.Error
}
func (f *DynamoDBAPI) WaitUntilTableNotExists(param1 *dynamodb.DescribeTableInput) error {
	f.WaitUntilTableNotExistsCall.mutex.Lock()
	defer f.WaitUntilTableNotExistsCall.mutex.Unlock()
	f.WaitUntilTableNotExistsCall.CallCount++
	f.WaitUntilTableNotExistsCall.Receives.DescribeTableInput = param1
	if f.WaitUntilTableNotExistsCall.Stub != nil {
		return f.WaitUntilTableNotExistsCall.Stub(param1)
	}
	return f.WaitUntilTableNotExistsCall.Returns.Error
}
func (f *DynamoDBAPI) WaitUntilTableNotExistsWithContext(param1 context.Context, param2 *dynamodb.DescribeTableInput, param3 ...request.WaiterOption) error {
	f.WaitUntilTableNotExistsWithContextCall.mutex.Lock()
	defer f.WaitUntilTableNotExistsWithContextCall.mutex.Unlock()
	f.WaitUntilTableNotExistsWithContextCall.CallCount++
	f.WaitUntilTableNotExistsWithContextCall.Receives.Context = param1
	f.WaitUntilTableNotExistsWithContextCall.Receives.DescribeTableInput = param2
	f.WaitUntilTableNotExistsWithContextCall.Receives.WaiterOptionSlice = param3
	if f.WaitUntilTableNotExistsWithContextCall.Stub != nil {
		return f.WaitUntilTableNotExistsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.WaitUntilTableNotExistsWithContextCall.Returns.Error
}
