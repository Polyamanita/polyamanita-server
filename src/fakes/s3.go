package fakes

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3API struct {
	AbortMultipartUploadCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			AbortMultipartUploadInput *s3.AbortMultipartUploadInput
		}
		Returns struct {
			AbortMultipartUploadOutput *s3.AbortMultipartUploadOutput
			Error                      error
		}
		Stub func(*s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error)
	}
	AbortMultipartUploadRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			AbortMultipartUploadInput *s3.AbortMultipartUploadInput
		}
		Returns struct {
			Request                    *request.Request
			AbortMultipartUploadOutput *s3.AbortMultipartUploadOutput
		}
		Stub func(*s3.AbortMultipartUploadInput) (*request.Request, *s3.AbortMultipartUploadOutput)
	}
	AbortMultipartUploadWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                   context.Context
			AbortMultipartUploadInput *s3.AbortMultipartUploadInput
			OptionSlice               []request.Option
		}
		Returns struct {
			AbortMultipartUploadOutput *s3.AbortMultipartUploadOutput
			Error                      error
		}
		Stub func(context.Context, *s3.AbortMultipartUploadInput, ...request.Option) (*s3.AbortMultipartUploadOutput, error)
	}
	CompleteMultipartUploadCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			CompleteMultipartUploadInput *s3.CompleteMultipartUploadInput
		}
		Returns struct {
			CompleteMultipartUploadOutput *s3.CompleteMultipartUploadOutput
			Error                         error
		}
		Stub func(*s3.CompleteMultipartUploadInput) (*s3.CompleteMultipartUploadOutput, error)
	}
	CompleteMultipartUploadRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			CompleteMultipartUploadInput *s3.CompleteMultipartUploadInput
		}
		Returns struct {
			Request                       *request.Request
			CompleteMultipartUploadOutput *s3.CompleteMultipartUploadOutput
		}
		Stub func(*s3.CompleteMultipartUploadInput) (*request.Request, *s3.CompleteMultipartUploadOutput)
	}
	CompleteMultipartUploadWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                      context.Context
			CompleteMultipartUploadInput *s3.CompleteMultipartUploadInput
			OptionSlice                  []request.Option
		}
		Returns struct {
			CompleteMultipartUploadOutput *s3.CompleteMultipartUploadOutput
			Error                         error
		}
		Stub func(context.Context, *s3.CompleteMultipartUploadInput, ...request.Option) (*s3.CompleteMultipartUploadOutput, error)
	}
	CopyObjectCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			CopyObjectInput *s3.CopyObjectInput
		}
		Returns struct {
			CopyObjectOutput *s3.CopyObjectOutput
			Error            error
		}
		Stub func(*s3.CopyObjectInput) (*s3.CopyObjectOutput, error)
	}
	CopyObjectRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			CopyObjectInput *s3.CopyObjectInput
		}
		Returns struct {
			Request          *request.Request
			CopyObjectOutput *s3.CopyObjectOutput
		}
		Stub func(*s3.CopyObjectInput) (*request.Request, *s3.CopyObjectOutput)
	}
	CopyObjectWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context         context.Context
			CopyObjectInput *s3.CopyObjectInput
			OptionSlice     []request.Option
		}
		Returns struct {
			CopyObjectOutput *s3.CopyObjectOutput
			Error            error
		}
		Stub func(context.Context, *s3.CopyObjectInput, ...request.Option) (*s3.CopyObjectOutput, error)
	}
	CreateBucketCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			CreateBucketInput *s3.CreateBucketInput
		}
		Returns struct {
			CreateBucketOutput *s3.CreateBucketOutput
			Error              error
		}
		Stub func(*s3.CreateBucketInput) (*s3.CreateBucketOutput, error)
	}
	CreateBucketRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			CreateBucketInput *s3.CreateBucketInput
		}
		Returns struct {
			Request            *request.Request
			CreateBucketOutput *s3.CreateBucketOutput
		}
		Stub func(*s3.CreateBucketInput) (*request.Request, *s3.CreateBucketOutput)
	}
	CreateBucketWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context           context.Context
			CreateBucketInput *s3.CreateBucketInput
			OptionSlice       []request.Option
		}
		Returns struct {
			CreateBucketOutput *s3.CreateBucketOutput
			Error              error
		}
		Stub func(context.Context, *s3.CreateBucketInput, ...request.Option) (*s3.CreateBucketOutput, error)
	}
	CreateMultipartUploadCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			CreateMultipartUploadInput *s3.CreateMultipartUploadInput
		}
		Returns struct {
			CreateMultipartUploadOutput *s3.CreateMultipartUploadOutput
			Error                       error
		}
		Stub func(*s3.CreateMultipartUploadInput) (*s3.CreateMultipartUploadOutput, error)
	}
	CreateMultipartUploadRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			CreateMultipartUploadInput *s3.CreateMultipartUploadInput
		}
		Returns struct {
			Request                     *request.Request
			CreateMultipartUploadOutput *s3.CreateMultipartUploadOutput
		}
		Stub func(*s3.CreateMultipartUploadInput) (*request.Request, *s3.CreateMultipartUploadOutput)
	}
	CreateMultipartUploadWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                    context.Context
			CreateMultipartUploadInput *s3.CreateMultipartUploadInput
			OptionSlice                []request.Option
		}
		Returns struct {
			CreateMultipartUploadOutput *s3.CreateMultipartUploadOutput
			Error                       error
		}
		Stub func(context.Context, *s3.CreateMultipartUploadInput, ...request.Option) (*s3.CreateMultipartUploadOutput, error)
	}
	DeleteBucketCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketInput *s3.DeleteBucketInput
		}
		Returns struct {
			DeleteBucketOutput *s3.DeleteBucketOutput
			Error              error
		}
		Stub func(*s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error)
	}
	DeleteBucketAnalyticsConfigurationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketAnalyticsConfigurationInput *s3.DeleteBucketAnalyticsConfigurationInput
		}
		Returns struct {
			DeleteBucketAnalyticsConfigurationOutput *s3.DeleteBucketAnalyticsConfigurationOutput
			Error                                    error
		}
		Stub func(*s3.DeleteBucketAnalyticsConfigurationInput) (*s3.DeleteBucketAnalyticsConfigurationOutput, error)
	}
	DeleteBucketAnalyticsConfigurationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketAnalyticsConfigurationInput *s3.DeleteBucketAnalyticsConfigurationInput
		}
		Returns struct {
			Request                                  *request.Request
			DeleteBucketAnalyticsConfigurationOutput *s3.DeleteBucketAnalyticsConfigurationOutput
		}
		Stub func(*s3.DeleteBucketAnalyticsConfigurationInput) (*request.Request, *s3.DeleteBucketAnalyticsConfigurationOutput)
	}
	DeleteBucketAnalyticsConfigurationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                                 context.Context
			DeleteBucketAnalyticsConfigurationInput *s3.DeleteBucketAnalyticsConfigurationInput
			OptionSlice                             []request.Option
		}
		Returns struct {
			DeleteBucketAnalyticsConfigurationOutput *s3.DeleteBucketAnalyticsConfigurationOutput
			Error                                    error
		}
		Stub func(context.Context, *s3.DeleteBucketAnalyticsConfigurationInput, ...request.Option) (*s3.DeleteBucketAnalyticsConfigurationOutput, error)
	}
	DeleteBucketCorsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketCorsInput *s3.DeleteBucketCorsInput
		}
		Returns struct {
			DeleteBucketCorsOutput *s3.DeleteBucketCorsOutput
			Error                  error
		}
		Stub func(*s3.DeleteBucketCorsInput) (*s3.DeleteBucketCorsOutput, error)
	}
	DeleteBucketCorsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketCorsInput *s3.DeleteBucketCorsInput
		}
		Returns struct {
			Request                *request.Request
			DeleteBucketCorsOutput *s3.DeleteBucketCorsOutput
		}
		Stub func(*s3.DeleteBucketCorsInput) (*request.Request, *s3.DeleteBucketCorsOutput)
	}
	DeleteBucketCorsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context               context.Context
			DeleteBucketCorsInput *s3.DeleteBucketCorsInput
			OptionSlice           []request.Option
		}
		Returns struct {
			DeleteBucketCorsOutput *s3.DeleteBucketCorsOutput
			Error                  error
		}
		Stub func(context.Context, *s3.DeleteBucketCorsInput, ...request.Option) (*s3.DeleteBucketCorsOutput, error)
	}
	DeleteBucketEncryptionCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketEncryptionInput *s3.DeleteBucketEncryptionInput
		}
		Returns struct {
			DeleteBucketEncryptionOutput *s3.DeleteBucketEncryptionOutput
			Error                        error
		}
		Stub func(*s3.DeleteBucketEncryptionInput) (*s3.DeleteBucketEncryptionOutput, error)
	}
	DeleteBucketEncryptionRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketEncryptionInput *s3.DeleteBucketEncryptionInput
		}
		Returns struct {
			Request                      *request.Request
			DeleteBucketEncryptionOutput *s3.DeleteBucketEncryptionOutput
		}
		Stub func(*s3.DeleteBucketEncryptionInput) (*request.Request, *s3.DeleteBucketEncryptionOutput)
	}
	DeleteBucketEncryptionWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                     context.Context
			DeleteBucketEncryptionInput *s3.DeleteBucketEncryptionInput
			OptionSlice                 []request.Option
		}
		Returns struct {
			DeleteBucketEncryptionOutput *s3.DeleteBucketEncryptionOutput
			Error                        error
		}
		Stub func(context.Context, *s3.DeleteBucketEncryptionInput, ...request.Option) (*s3.DeleteBucketEncryptionOutput, error)
	}
	DeleteBucketIntelligentTieringConfigurationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketIntelligentTieringConfigurationInput *s3.DeleteBucketIntelligentTieringConfigurationInput
		}
		Returns struct {
			DeleteBucketIntelligentTieringConfigurationOutput *s3.DeleteBucketIntelligentTieringConfigurationOutput
			Error                                             error
		}
		Stub func(*s3.DeleteBucketIntelligentTieringConfigurationInput) (*s3.DeleteBucketIntelligentTieringConfigurationOutput, error)
	}
	DeleteBucketIntelligentTieringConfigurationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketIntelligentTieringConfigurationInput *s3.DeleteBucketIntelligentTieringConfigurationInput
		}
		Returns struct {
			Request                                           *request.Request
			DeleteBucketIntelligentTieringConfigurationOutput *s3.DeleteBucketIntelligentTieringConfigurationOutput
		}
		Stub func(*s3.DeleteBucketIntelligentTieringConfigurationInput) (*request.Request, *s3.DeleteBucketIntelligentTieringConfigurationOutput)
	}
	DeleteBucketIntelligentTieringConfigurationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                                          context.Context
			DeleteBucketIntelligentTieringConfigurationInput *s3.DeleteBucketIntelligentTieringConfigurationInput
			OptionSlice                                      []request.Option
		}
		Returns struct {
			DeleteBucketIntelligentTieringConfigurationOutput *s3.DeleteBucketIntelligentTieringConfigurationOutput
			Error                                             error
		}
		Stub func(context.Context, *s3.DeleteBucketIntelligentTieringConfigurationInput, ...request.Option) (*s3.DeleteBucketIntelligentTieringConfigurationOutput, error)
	}
	DeleteBucketInventoryConfigurationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketInventoryConfigurationInput *s3.DeleteBucketInventoryConfigurationInput
		}
		Returns struct {
			DeleteBucketInventoryConfigurationOutput *s3.DeleteBucketInventoryConfigurationOutput
			Error                                    error
		}
		Stub func(*s3.DeleteBucketInventoryConfigurationInput) (*s3.DeleteBucketInventoryConfigurationOutput, error)
	}
	DeleteBucketInventoryConfigurationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketInventoryConfigurationInput *s3.DeleteBucketInventoryConfigurationInput
		}
		Returns struct {
			Request                                  *request.Request
			DeleteBucketInventoryConfigurationOutput *s3.DeleteBucketInventoryConfigurationOutput
		}
		Stub func(*s3.DeleteBucketInventoryConfigurationInput) (*request.Request, *s3.DeleteBucketInventoryConfigurationOutput)
	}
	DeleteBucketInventoryConfigurationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                                 context.Context
			DeleteBucketInventoryConfigurationInput *s3.DeleteBucketInventoryConfigurationInput
			OptionSlice                             []request.Option
		}
		Returns struct {
			DeleteBucketInventoryConfigurationOutput *s3.DeleteBucketInventoryConfigurationOutput
			Error                                    error
		}
		Stub func(context.Context, *s3.DeleteBucketInventoryConfigurationInput, ...request.Option) (*s3.DeleteBucketInventoryConfigurationOutput, error)
	}
	DeleteBucketLifecycleCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketLifecycleInput *s3.DeleteBucketLifecycleInput
		}
		Returns struct {
			DeleteBucketLifecycleOutput *s3.DeleteBucketLifecycleOutput
			Error                       error
		}
		Stub func(*s3.DeleteBucketLifecycleInput) (*s3.DeleteBucketLifecycleOutput, error)
	}
	DeleteBucketLifecycleRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketLifecycleInput *s3.DeleteBucketLifecycleInput
		}
		Returns struct {
			Request                     *request.Request
			DeleteBucketLifecycleOutput *s3.DeleteBucketLifecycleOutput
		}
		Stub func(*s3.DeleteBucketLifecycleInput) (*request.Request, *s3.DeleteBucketLifecycleOutput)
	}
	DeleteBucketLifecycleWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                    context.Context
			DeleteBucketLifecycleInput *s3.DeleteBucketLifecycleInput
			OptionSlice                []request.Option
		}
		Returns struct {
			DeleteBucketLifecycleOutput *s3.DeleteBucketLifecycleOutput
			Error                       error
		}
		Stub func(context.Context, *s3.DeleteBucketLifecycleInput, ...request.Option) (*s3.DeleteBucketLifecycleOutput, error)
	}
	DeleteBucketMetricsConfigurationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketMetricsConfigurationInput *s3.DeleteBucketMetricsConfigurationInput
		}
		Returns struct {
			DeleteBucketMetricsConfigurationOutput *s3.DeleteBucketMetricsConfigurationOutput
			Error                                  error
		}
		Stub func(*s3.DeleteBucketMetricsConfigurationInput) (*s3.DeleteBucketMetricsConfigurationOutput, error)
	}
	DeleteBucketMetricsConfigurationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketMetricsConfigurationInput *s3.DeleteBucketMetricsConfigurationInput
		}
		Returns struct {
			Request                                *request.Request
			DeleteBucketMetricsConfigurationOutput *s3.DeleteBucketMetricsConfigurationOutput
		}
		Stub func(*s3.DeleteBucketMetricsConfigurationInput) (*request.Request, *s3.DeleteBucketMetricsConfigurationOutput)
	}
	DeleteBucketMetricsConfigurationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                               context.Context
			DeleteBucketMetricsConfigurationInput *s3.DeleteBucketMetricsConfigurationInput
			OptionSlice                           []request.Option
		}
		Returns struct {
			DeleteBucketMetricsConfigurationOutput *s3.DeleteBucketMetricsConfigurationOutput
			Error                                  error
		}
		Stub func(context.Context, *s3.DeleteBucketMetricsConfigurationInput, ...request.Option) (*s3.DeleteBucketMetricsConfigurationOutput, error)
	}
	DeleteBucketOwnershipControlsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketOwnershipControlsInput *s3.DeleteBucketOwnershipControlsInput
		}
		Returns struct {
			DeleteBucketOwnershipControlsOutput *s3.DeleteBucketOwnershipControlsOutput
			Error                               error
		}
		Stub func(*s3.DeleteBucketOwnershipControlsInput) (*s3.DeleteBucketOwnershipControlsOutput, error)
	}
	DeleteBucketOwnershipControlsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketOwnershipControlsInput *s3.DeleteBucketOwnershipControlsInput
		}
		Returns struct {
			Request                             *request.Request
			DeleteBucketOwnershipControlsOutput *s3.DeleteBucketOwnershipControlsOutput
		}
		Stub func(*s3.DeleteBucketOwnershipControlsInput) (*request.Request, *s3.DeleteBucketOwnershipControlsOutput)
	}
	DeleteBucketOwnershipControlsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                            context.Context
			DeleteBucketOwnershipControlsInput *s3.DeleteBucketOwnershipControlsInput
			OptionSlice                        []request.Option
		}
		Returns struct {
			DeleteBucketOwnershipControlsOutput *s3.DeleteBucketOwnershipControlsOutput
			Error                               error
		}
		Stub func(context.Context, *s3.DeleteBucketOwnershipControlsInput, ...request.Option) (*s3.DeleteBucketOwnershipControlsOutput, error)
	}
	DeleteBucketPolicyCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketPolicyInput *s3.DeleteBucketPolicyInput
		}
		Returns struct {
			DeleteBucketPolicyOutput *s3.DeleteBucketPolicyOutput
			Error                    error
		}
		Stub func(*s3.DeleteBucketPolicyInput) (*s3.DeleteBucketPolicyOutput, error)
	}
	DeleteBucketPolicyRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketPolicyInput *s3.DeleteBucketPolicyInput
		}
		Returns struct {
			Request                  *request.Request
			DeleteBucketPolicyOutput *s3.DeleteBucketPolicyOutput
		}
		Stub func(*s3.DeleteBucketPolicyInput) (*request.Request, *s3.DeleteBucketPolicyOutput)
	}
	DeleteBucketPolicyWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                 context.Context
			DeleteBucketPolicyInput *s3.DeleteBucketPolicyInput
			OptionSlice             []request.Option
		}
		Returns struct {
			DeleteBucketPolicyOutput *s3.DeleteBucketPolicyOutput
			Error                    error
		}
		Stub func(context.Context, *s3.DeleteBucketPolicyInput, ...request.Option) (*s3.DeleteBucketPolicyOutput, error)
	}
	DeleteBucketReplicationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketReplicationInput *s3.DeleteBucketReplicationInput
		}
		Returns struct {
			DeleteBucketReplicationOutput *s3.DeleteBucketReplicationOutput
			Error                         error
		}
		Stub func(*s3.DeleteBucketReplicationInput) (*s3.DeleteBucketReplicationOutput, error)
	}
	DeleteBucketReplicationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketReplicationInput *s3.DeleteBucketReplicationInput
		}
		Returns struct {
			Request                       *request.Request
			DeleteBucketReplicationOutput *s3.DeleteBucketReplicationOutput
		}
		Stub func(*s3.DeleteBucketReplicationInput) (*request.Request, *s3.DeleteBucketReplicationOutput)
	}
	DeleteBucketReplicationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                      context.Context
			DeleteBucketReplicationInput *s3.DeleteBucketReplicationInput
			OptionSlice                  []request.Option
		}
		Returns struct {
			DeleteBucketReplicationOutput *s3.DeleteBucketReplicationOutput
			Error                         error
		}
		Stub func(context.Context, *s3.DeleteBucketReplicationInput, ...request.Option) (*s3.DeleteBucketReplicationOutput, error)
	}
	DeleteBucketRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketInput *s3.DeleteBucketInput
		}
		Returns struct {
			Request            *request.Request
			DeleteBucketOutput *s3.DeleteBucketOutput
		}
		Stub func(*s3.DeleteBucketInput) (*request.Request, *s3.DeleteBucketOutput)
	}
	DeleteBucketTaggingCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketTaggingInput *s3.DeleteBucketTaggingInput
		}
		Returns struct {
			DeleteBucketTaggingOutput *s3.DeleteBucketTaggingOutput
			Error                     error
		}
		Stub func(*s3.DeleteBucketTaggingInput) (*s3.DeleteBucketTaggingOutput, error)
	}
	DeleteBucketTaggingRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketTaggingInput *s3.DeleteBucketTaggingInput
		}
		Returns struct {
			Request                   *request.Request
			DeleteBucketTaggingOutput *s3.DeleteBucketTaggingOutput
		}
		Stub func(*s3.DeleteBucketTaggingInput) (*request.Request, *s3.DeleteBucketTaggingOutput)
	}
	DeleteBucketTaggingWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                  context.Context
			DeleteBucketTaggingInput *s3.DeleteBucketTaggingInput
			OptionSlice              []request.Option
		}
		Returns struct {
			DeleteBucketTaggingOutput *s3.DeleteBucketTaggingOutput
			Error                     error
		}
		Stub func(context.Context, *s3.DeleteBucketTaggingInput, ...request.Option) (*s3.DeleteBucketTaggingOutput, error)
	}
	DeleteBucketWebsiteCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketWebsiteInput *s3.DeleteBucketWebsiteInput
		}
		Returns struct {
			DeleteBucketWebsiteOutput *s3.DeleteBucketWebsiteOutput
			Error                     error
		}
		Stub func(*s3.DeleteBucketWebsiteInput) (*s3.DeleteBucketWebsiteOutput, error)
	}
	DeleteBucketWebsiteRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketWebsiteInput *s3.DeleteBucketWebsiteInput
		}
		Returns struct {
			Request                   *request.Request
			DeleteBucketWebsiteOutput *s3.DeleteBucketWebsiteOutput
		}
		Stub func(*s3.DeleteBucketWebsiteInput) (*request.Request, *s3.DeleteBucketWebsiteOutput)
	}
	DeleteBucketWebsiteWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                  context.Context
			DeleteBucketWebsiteInput *s3.DeleteBucketWebsiteInput
			OptionSlice              []request.Option
		}
		Returns struct {
			DeleteBucketWebsiteOutput *s3.DeleteBucketWebsiteOutput
			Error                     error
		}
		Stub func(context.Context, *s3.DeleteBucketWebsiteInput, ...request.Option) (*s3.DeleteBucketWebsiteOutput, error)
	}
	DeleteBucketWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context           context.Context
			DeleteBucketInput *s3.DeleteBucketInput
			OptionSlice       []request.Option
		}
		Returns struct {
			DeleteBucketOutput *s3.DeleteBucketOutput
			Error              error
		}
		Stub func(context.Context, *s3.DeleteBucketInput, ...request.Option) (*s3.DeleteBucketOutput, error)
	}
	DeleteObjectCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteObjectInput *s3.DeleteObjectInput
		}
		Returns struct {
			DeleteObjectOutput *s3.DeleteObjectOutput
			Error              error
		}
		Stub func(*s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error)
	}
	DeleteObjectRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteObjectInput *s3.DeleteObjectInput
		}
		Returns struct {
			Request            *request.Request
			DeleteObjectOutput *s3.DeleteObjectOutput
		}
		Stub func(*s3.DeleteObjectInput) (*request.Request, *s3.DeleteObjectOutput)
	}
	DeleteObjectTaggingCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteObjectTaggingInput *s3.DeleteObjectTaggingInput
		}
		Returns struct {
			DeleteObjectTaggingOutput *s3.DeleteObjectTaggingOutput
			Error                     error
		}
		Stub func(*s3.DeleteObjectTaggingInput) (*s3.DeleteObjectTaggingOutput, error)
	}
	DeleteObjectTaggingRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteObjectTaggingInput *s3.DeleteObjectTaggingInput
		}
		Returns struct {
			Request                   *request.Request
			DeleteObjectTaggingOutput *s3.DeleteObjectTaggingOutput
		}
		Stub func(*s3.DeleteObjectTaggingInput) (*request.Request, *s3.DeleteObjectTaggingOutput)
	}
	DeleteObjectTaggingWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                  context.Context
			DeleteObjectTaggingInput *s3.DeleteObjectTaggingInput
			OptionSlice              []request.Option
		}
		Returns struct {
			DeleteObjectTaggingOutput *s3.DeleteObjectTaggingOutput
			Error                     error
		}
		Stub func(context.Context, *s3.DeleteObjectTaggingInput, ...request.Option) (*s3.DeleteObjectTaggingOutput, error)
	}
	DeleteObjectWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context           context.Context
			DeleteObjectInput *s3.DeleteObjectInput
			OptionSlice       []request.Option
		}
		Returns struct {
			DeleteObjectOutput *s3.DeleteObjectOutput
			Error              error
		}
		Stub func(context.Context, *s3.DeleteObjectInput, ...request.Option) (*s3.DeleteObjectOutput, error)
	}
	DeleteObjectsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteObjectsInput *s3.DeleteObjectsInput
		}
		Returns struct {
			DeleteObjectsOutput *s3.DeleteObjectsOutput
			Error               error
		}
		Stub func(*s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error)
	}
	DeleteObjectsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeleteObjectsInput *s3.DeleteObjectsInput
		}
		Returns struct {
			Request             *request.Request
			DeleteObjectsOutput *s3.DeleteObjectsOutput
		}
		Stub func(*s3.DeleteObjectsInput) (*request.Request, *s3.DeleteObjectsOutput)
	}
	DeleteObjectsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context            context.Context
			DeleteObjectsInput *s3.DeleteObjectsInput
			OptionSlice        []request.Option
		}
		Returns struct {
			DeleteObjectsOutput *s3.DeleteObjectsOutput
			Error               error
		}
		Stub func(context.Context, *s3.DeleteObjectsInput, ...request.Option) (*s3.DeleteObjectsOutput, error)
	}
	DeletePublicAccessBlockCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeletePublicAccessBlockInput *s3.DeletePublicAccessBlockInput
		}
		Returns struct {
			DeletePublicAccessBlockOutput *s3.DeletePublicAccessBlockOutput
			Error                         error
		}
		Stub func(*s3.DeletePublicAccessBlockInput) (*s3.DeletePublicAccessBlockOutput, error)
	}
	DeletePublicAccessBlockRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			DeletePublicAccessBlockInput *s3.DeletePublicAccessBlockInput
		}
		Returns struct {
			Request                       *request.Request
			DeletePublicAccessBlockOutput *s3.DeletePublicAccessBlockOutput
		}
		Stub func(*s3.DeletePublicAccessBlockInput) (*request.Request, *s3.DeletePublicAccessBlockOutput)
	}
	DeletePublicAccessBlockWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                      context.Context
			DeletePublicAccessBlockInput *s3.DeletePublicAccessBlockInput
			OptionSlice                  []request.Option
		}
		Returns struct {
			DeletePublicAccessBlockOutput *s3.DeletePublicAccessBlockOutput
			Error                         error
		}
		Stub func(context.Context, *s3.DeletePublicAccessBlockInput, ...request.Option) (*s3.DeletePublicAccessBlockOutput, error)
	}
	GetBucketAccelerateConfigurationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketAccelerateConfigurationInput *s3.GetBucketAccelerateConfigurationInput
		}
		Returns struct {
			GetBucketAccelerateConfigurationOutput *s3.GetBucketAccelerateConfigurationOutput
			Error                                  error
		}
		Stub func(*s3.GetBucketAccelerateConfigurationInput) (*s3.GetBucketAccelerateConfigurationOutput, error)
	}
	GetBucketAccelerateConfigurationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketAccelerateConfigurationInput *s3.GetBucketAccelerateConfigurationInput
		}
		Returns struct {
			Request                                *request.Request
			GetBucketAccelerateConfigurationOutput *s3.GetBucketAccelerateConfigurationOutput
		}
		Stub func(*s3.GetBucketAccelerateConfigurationInput) (*request.Request, *s3.GetBucketAccelerateConfigurationOutput)
	}
	GetBucketAccelerateConfigurationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                               context.Context
			GetBucketAccelerateConfigurationInput *s3.GetBucketAccelerateConfigurationInput
			OptionSlice                           []request.Option
		}
		Returns struct {
			GetBucketAccelerateConfigurationOutput *s3.GetBucketAccelerateConfigurationOutput
			Error                                  error
		}
		Stub func(context.Context, *s3.GetBucketAccelerateConfigurationInput, ...request.Option) (*s3.GetBucketAccelerateConfigurationOutput, error)
	}
	GetBucketAclCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketAclInput *s3.GetBucketAclInput
		}
		Returns struct {
			GetBucketAclOutput *s3.GetBucketAclOutput
			Error              error
		}
		Stub func(*s3.GetBucketAclInput) (*s3.GetBucketAclOutput, error)
	}
	GetBucketAclRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketAclInput *s3.GetBucketAclInput
		}
		Returns struct {
			Request            *request.Request
			GetBucketAclOutput *s3.GetBucketAclOutput
		}
		Stub func(*s3.GetBucketAclInput) (*request.Request, *s3.GetBucketAclOutput)
	}
	GetBucketAclWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context           context.Context
			GetBucketAclInput *s3.GetBucketAclInput
			OptionSlice       []request.Option
		}
		Returns struct {
			GetBucketAclOutput *s3.GetBucketAclOutput
			Error              error
		}
		Stub func(context.Context, *s3.GetBucketAclInput, ...request.Option) (*s3.GetBucketAclOutput, error)
	}
	GetBucketAnalyticsConfigurationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketAnalyticsConfigurationInput *s3.GetBucketAnalyticsConfigurationInput
		}
		Returns struct {
			GetBucketAnalyticsConfigurationOutput *s3.GetBucketAnalyticsConfigurationOutput
			Error                                 error
		}
		Stub func(*s3.GetBucketAnalyticsConfigurationInput) (*s3.GetBucketAnalyticsConfigurationOutput, error)
	}
	GetBucketAnalyticsConfigurationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketAnalyticsConfigurationInput *s3.GetBucketAnalyticsConfigurationInput
		}
		Returns struct {
			Request                               *request.Request
			GetBucketAnalyticsConfigurationOutput *s3.GetBucketAnalyticsConfigurationOutput
		}
		Stub func(*s3.GetBucketAnalyticsConfigurationInput) (*request.Request, *s3.GetBucketAnalyticsConfigurationOutput)
	}
	GetBucketAnalyticsConfigurationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                              context.Context
			GetBucketAnalyticsConfigurationInput *s3.GetBucketAnalyticsConfigurationInput
			OptionSlice                          []request.Option
		}
		Returns struct {
			GetBucketAnalyticsConfigurationOutput *s3.GetBucketAnalyticsConfigurationOutput
			Error                                 error
		}
		Stub func(context.Context, *s3.GetBucketAnalyticsConfigurationInput, ...request.Option) (*s3.GetBucketAnalyticsConfigurationOutput, error)
	}
	GetBucketCorsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketCorsInput *s3.GetBucketCorsInput
		}
		Returns struct {
			GetBucketCorsOutput *s3.GetBucketCorsOutput
			Error               error
		}
		Stub func(*s3.GetBucketCorsInput) (*s3.GetBucketCorsOutput, error)
	}
	GetBucketCorsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketCorsInput *s3.GetBucketCorsInput
		}
		Returns struct {
			Request             *request.Request
			GetBucketCorsOutput *s3.GetBucketCorsOutput
		}
		Stub func(*s3.GetBucketCorsInput) (*request.Request, *s3.GetBucketCorsOutput)
	}
	GetBucketCorsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context            context.Context
			GetBucketCorsInput *s3.GetBucketCorsInput
			OptionSlice        []request.Option
		}
		Returns struct {
			GetBucketCorsOutput *s3.GetBucketCorsOutput
			Error               error
		}
		Stub func(context.Context, *s3.GetBucketCorsInput, ...request.Option) (*s3.GetBucketCorsOutput, error)
	}
	GetBucketEncryptionCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketEncryptionInput *s3.GetBucketEncryptionInput
		}
		Returns struct {
			GetBucketEncryptionOutput *s3.GetBucketEncryptionOutput
			Error                     error
		}
		Stub func(*s3.GetBucketEncryptionInput) (*s3.GetBucketEncryptionOutput, error)
	}
	GetBucketEncryptionRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketEncryptionInput *s3.GetBucketEncryptionInput
		}
		Returns struct {
			Request                   *request.Request
			GetBucketEncryptionOutput *s3.GetBucketEncryptionOutput
		}
		Stub func(*s3.GetBucketEncryptionInput) (*request.Request, *s3.GetBucketEncryptionOutput)
	}
	GetBucketEncryptionWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                  context.Context
			GetBucketEncryptionInput *s3.GetBucketEncryptionInput
			OptionSlice              []request.Option
		}
		Returns struct {
			GetBucketEncryptionOutput *s3.GetBucketEncryptionOutput
			Error                     error
		}
		Stub func(context.Context, *s3.GetBucketEncryptionInput, ...request.Option) (*s3.GetBucketEncryptionOutput, error)
	}
	GetBucketIntelligentTieringConfigurationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketIntelligentTieringConfigurationInput *s3.GetBucketIntelligentTieringConfigurationInput
		}
		Returns struct {
			GetBucketIntelligentTieringConfigurationOutput *s3.GetBucketIntelligentTieringConfigurationOutput
			Error                                          error
		}
		Stub func(*s3.GetBucketIntelligentTieringConfigurationInput) (*s3.GetBucketIntelligentTieringConfigurationOutput, error)
	}
	GetBucketIntelligentTieringConfigurationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketIntelligentTieringConfigurationInput *s3.GetBucketIntelligentTieringConfigurationInput
		}
		Returns struct {
			Request                                        *request.Request
			GetBucketIntelligentTieringConfigurationOutput *s3.GetBucketIntelligentTieringConfigurationOutput
		}
		Stub func(*s3.GetBucketIntelligentTieringConfigurationInput) (*request.Request, *s3.GetBucketIntelligentTieringConfigurationOutput)
	}
	GetBucketIntelligentTieringConfigurationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                                       context.Context
			GetBucketIntelligentTieringConfigurationInput *s3.GetBucketIntelligentTieringConfigurationInput
			OptionSlice                                   []request.Option
		}
		Returns struct {
			GetBucketIntelligentTieringConfigurationOutput *s3.GetBucketIntelligentTieringConfigurationOutput
			Error                                          error
		}
		Stub func(context.Context, *s3.GetBucketIntelligentTieringConfigurationInput, ...request.Option) (*s3.GetBucketIntelligentTieringConfigurationOutput, error)
	}
	GetBucketInventoryConfigurationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketInventoryConfigurationInput *s3.GetBucketInventoryConfigurationInput
		}
		Returns struct {
			GetBucketInventoryConfigurationOutput *s3.GetBucketInventoryConfigurationOutput
			Error                                 error
		}
		Stub func(*s3.GetBucketInventoryConfigurationInput) (*s3.GetBucketInventoryConfigurationOutput, error)
	}
	GetBucketInventoryConfigurationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketInventoryConfigurationInput *s3.GetBucketInventoryConfigurationInput
		}
		Returns struct {
			Request                               *request.Request
			GetBucketInventoryConfigurationOutput *s3.GetBucketInventoryConfigurationOutput
		}
		Stub func(*s3.GetBucketInventoryConfigurationInput) (*request.Request, *s3.GetBucketInventoryConfigurationOutput)
	}
	GetBucketInventoryConfigurationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                              context.Context
			GetBucketInventoryConfigurationInput *s3.GetBucketInventoryConfigurationInput
			OptionSlice                          []request.Option
		}
		Returns struct {
			GetBucketInventoryConfigurationOutput *s3.GetBucketInventoryConfigurationOutput
			Error                                 error
		}
		Stub func(context.Context, *s3.GetBucketInventoryConfigurationInput, ...request.Option) (*s3.GetBucketInventoryConfigurationOutput, error)
	}
	GetBucketLifecycleCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketLifecycleInput *s3.GetBucketLifecycleInput
		}
		Returns struct {
			GetBucketLifecycleOutput *s3.GetBucketLifecycleOutput
			Error                    error
		}
		Stub func(*s3.GetBucketLifecycleInput) (*s3.GetBucketLifecycleOutput, error)
	}
	GetBucketLifecycleConfigurationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketLifecycleConfigurationInput *s3.GetBucketLifecycleConfigurationInput
		}
		Returns struct {
			GetBucketLifecycleConfigurationOutput *s3.GetBucketLifecycleConfigurationOutput
			Error                                 error
		}
		Stub func(*s3.GetBucketLifecycleConfigurationInput) (*s3.GetBucketLifecycleConfigurationOutput, error)
	}
	GetBucketLifecycleConfigurationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketLifecycleConfigurationInput *s3.GetBucketLifecycleConfigurationInput
		}
		Returns struct {
			Request                               *request.Request
			GetBucketLifecycleConfigurationOutput *s3.GetBucketLifecycleConfigurationOutput
		}
		Stub func(*s3.GetBucketLifecycleConfigurationInput) (*request.Request, *s3.GetBucketLifecycleConfigurationOutput)
	}
	GetBucketLifecycleConfigurationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                              context.Context
			GetBucketLifecycleConfigurationInput *s3.GetBucketLifecycleConfigurationInput
			OptionSlice                          []request.Option
		}
		Returns struct {
			GetBucketLifecycleConfigurationOutput *s3.GetBucketLifecycleConfigurationOutput
			Error                                 error
		}
		Stub func(context.Context, *s3.GetBucketLifecycleConfigurationInput, ...request.Option) (*s3.GetBucketLifecycleConfigurationOutput, error)
	}
	GetBucketLifecycleRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketLifecycleInput *s3.GetBucketLifecycleInput
		}
		Returns struct {
			Request                  *request.Request
			GetBucketLifecycleOutput *s3.GetBucketLifecycleOutput
		}
		Stub func(*s3.GetBucketLifecycleInput) (*request.Request, *s3.GetBucketLifecycleOutput)
	}
	GetBucketLifecycleWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                 context.Context
			GetBucketLifecycleInput *s3.GetBucketLifecycleInput
			OptionSlice             []request.Option
		}
		Returns struct {
			GetBucketLifecycleOutput *s3.GetBucketLifecycleOutput
			Error                    error
		}
		Stub func(context.Context, *s3.GetBucketLifecycleInput, ...request.Option) (*s3.GetBucketLifecycleOutput, error)
	}
	GetBucketLocationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketLocationInput *s3.GetBucketLocationInput
		}
		Returns struct {
			GetBucketLocationOutput *s3.GetBucketLocationOutput
			Error                   error
		}
		Stub func(*s3.GetBucketLocationInput) (*s3.GetBucketLocationOutput, error)
	}
	GetBucketLocationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketLocationInput *s3.GetBucketLocationInput
		}
		Returns struct {
			Request                 *request.Request
			GetBucketLocationOutput *s3.GetBucketLocationOutput
		}
		Stub func(*s3.GetBucketLocationInput) (*request.Request, *s3.GetBucketLocationOutput)
	}
	GetBucketLocationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                context.Context
			GetBucketLocationInput *s3.GetBucketLocationInput
			OptionSlice            []request.Option
		}
		Returns struct {
			GetBucketLocationOutput *s3.GetBucketLocationOutput
			Error                   error
		}
		Stub func(context.Context, *s3.GetBucketLocationInput, ...request.Option) (*s3.GetBucketLocationOutput, error)
	}
	GetBucketLoggingCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketLoggingInput *s3.GetBucketLoggingInput
		}
		Returns struct {
			GetBucketLoggingOutput *s3.GetBucketLoggingOutput
			Error                  error
		}
		Stub func(*s3.GetBucketLoggingInput) (*s3.GetBucketLoggingOutput, error)
	}
	GetBucketLoggingRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketLoggingInput *s3.GetBucketLoggingInput
		}
		Returns struct {
			Request                *request.Request
			GetBucketLoggingOutput *s3.GetBucketLoggingOutput
		}
		Stub func(*s3.GetBucketLoggingInput) (*request.Request, *s3.GetBucketLoggingOutput)
	}
	GetBucketLoggingWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context               context.Context
			GetBucketLoggingInput *s3.GetBucketLoggingInput
			OptionSlice           []request.Option
		}
		Returns struct {
			GetBucketLoggingOutput *s3.GetBucketLoggingOutput
			Error                  error
		}
		Stub func(context.Context, *s3.GetBucketLoggingInput, ...request.Option) (*s3.GetBucketLoggingOutput, error)
	}
	GetBucketMetricsConfigurationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketMetricsConfigurationInput *s3.GetBucketMetricsConfigurationInput
		}
		Returns struct {
			GetBucketMetricsConfigurationOutput *s3.GetBucketMetricsConfigurationOutput
			Error                               error
		}
		Stub func(*s3.GetBucketMetricsConfigurationInput) (*s3.GetBucketMetricsConfigurationOutput, error)
	}
	GetBucketMetricsConfigurationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketMetricsConfigurationInput *s3.GetBucketMetricsConfigurationInput
		}
		Returns struct {
			Request                             *request.Request
			GetBucketMetricsConfigurationOutput *s3.GetBucketMetricsConfigurationOutput
		}
		Stub func(*s3.GetBucketMetricsConfigurationInput) (*request.Request, *s3.GetBucketMetricsConfigurationOutput)
	}
	GetBucketMetricsConfigurationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                            context.Context
			GetBucketMetricsConfigurationInput *s3.GetBucketMetricsConfigurationInput
			OptionSlice                        []request.Option
		}
		Returns struct {
			GetBucketMetricsConfigurationOutput *s3.GetBucketMetricsConfigurationOutput
			Error                               error
		}
		Stub func(context.Context, *s3.GetBucketMetricsConfigurationInput, ...request.Option) (*s3.GetBucketMetricsConfigurationOutput, error)
	}
	GetBucketNotificationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketNotificationConfigurationRequest *s3.GetBucketNotificationConfigurationRequest
		}
		Returns struct {
			NotificationConfigurationDeprecated *s3.NotificationConfigurationDeprecated
			Error                               error
		}
		Stub func(*s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfigurationDeprecated, error)
	}
	GetBucketNotificationConfigurationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketNotificationConfigurationRequest *s3.GetBucketNotificationConfigurationRequest
		}
		Returns struct {
			NotificationConfiguration *s3.NotificationConfiguration
			Error                     error
		}
		Stub func(*s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfiguration, error)
	}
	GetBucketNotificationConfigurationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketNotificationConfigurationRequest *s3.GetBucketNotificationConfigurationRequest
		}
		Returns struct {
			Request                   *request.Request
			NotificationConfiguration *s3.NotificationConfiguration
		}
		Stub func(*s3.GetBucketNotificationConfigurationRequest) (*request.Request, *s3.NotificationConfiguration)
	}
	GetBucketNotificationConfigurationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                                   context.Context
			GetBucketNotificationConfigurationRequest *s3.GetBucketNotificationConfigurationRequest
			OptionSlice                               []request.Option
		}
		Returns struct {
			NotificationConfiguration *s3.NotificationConfiguration
			Error                     error
		}
		Stub func(context.Context, *s3.GetBucketNotificationConfigurationRequest, ...request.Option) (*s3.NotificationConfiguration, error)
	}
	GetBucketNotificationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketNotificationConfigurationRequest *s3.GetBucketNotificationConfigurationRequest
		}
		Returns struct {
			Request                             *request.Request
			NotificationConfigurationDeprecated *s3.NotificationConfigurationDeprecated
		}
		Stub func(*s3.GetBucketNotificationConfigurationRequest) (*request.Request, *s3.NotificationConfigurationDeprecated)
	}
	GetBucketNotificationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                                   context.Context
			GetBucketNotificationConfigurationRequest *s3.GetBucketNotificationConfigurationRequest
			OptionSlice                               []request.Option
		}
		Returns struct {
			NotificationConfigurationDeprecated *s3.NotificationConfigurationDeprecated
			Error                               error
		}
		Stub func(context.Context, *s3.GetBucketNotificationConfigurationRequest, ...request.Option) (*s3.NotificationConfigurationDeprecated, error)
	}
	GetBucketOwnershipControlsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketOwnershipControlsInput *s3.GetBucketOwnershipControlsInput
		}
		Returns struct {
			GetBucketOwnershipControlsOutput *s3.GetBucketOwnershipControlsOutput
			Error                            error
		}
		Stub func(*s3.GetBucketOwnershipControlsInput) (*s3.GetBucketOwnershipControlsOutput, error)
	}
	GetBucketOwnershipControlsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketOwnershipControlsInput *s3.GetBucketOwnershipControlsInput
		}
		Returns struct {
			Request                          *request.Request
			GetBucketOwnershipControlsOutput *s3.GetBucketOwnershipControlsOutput
		}
		Stub func(*s3.GetBucketOwnershipControlsInput) (*request.Request, *s3.GetBucketOwnershipControlsOutput)
	}
	GetBucketOwnershipControlsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                         context.Context
			GetBucketOwnershipControlsInput *s3.GetBucketOwnershipControlsInput
			OptionSlice                     []request.Option
		}
		Returns struct {
			GetBucketOwnershipControlsOutput *s3.GetBucketOwnershipControlsOutput
			Error                            error
		}
		Stub func(context.Context, *s3.GetBucketOwnershipControlsInput, ...request.Option) (*s3.GetBucketOwnershipControlsOutput, error)
	}
	GetBucketPolicyCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketPolicyInput *s3.GetBucketPolicyInput
		}
		Returns struct {
			GetBucketPolicyOutput *s3.GetBucketPolicyOutput
			Error                 error
		}
		Stub func(*s3.GetBucketPolicyInput) (*s3.GetBucketPolicyOutput, error)
	}
	GetBucketPolicyRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketPolicyInput *s3.GetBucketPolicyInput
		}
		Returns struct {
			Request               *request.Request
			GetBucketPolicyOutput *s3.GetBucketPolicyOutput
		}
		Stub func(*s3.GetBucketPolicyInput) (*request.Request, *s3.GetBucketPolicyOutput)
	}
	GetBucketPolicyStatusCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketPolicyStatusInput *s3.GetBucketPolicyStatusInput
		}
		Returns struct {
			GetBucketPolicyStatusOutput *s3.GetBucketPolicyStatusOutput
			Error                       error
		}
		Stub func(*s3.GetBucketPolicyStatusInput) (*s3.GetBucketPolicyStatusOutput, error)
	}
	GetBucketPolicyStatusRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketPolicyStatusInput *s3.GetBucketPolicyStatusInput
		}
		Returns struct {
			Request                     *request.Request
			GetBucketPolicyStatusOutput *s3.GetBucketPolicyStatusOutput
		}
		Stub func(*s3.GetBucketPolicyStatusInput) (*request.Request, *s3.GetBucketPolicyStatusOutput)
	}
	GetBucketPolicyStatusWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                    context.Context
			GetBucketPolicyStatusInput *s3.GetBucketPolicyStatusInput
			OptionSlice                []request.Option
		}
		Returns struct {
			GetBucketPolicyStatusOutput *s3.GetBucketPolicyStatusOutput
			Error                       error
		}
		Stub func(context.Context, *s3.GetBucketPolicyStatusInput, ...request.Option) (*s3.GetBucketPolicyStatusOutput, error)
	}
	GetBucketPolicyWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context              context.Context
			GetBucketPolicyInput *s3.GetBucketPolicyInput
			OptionSlice          []request.Option
		}
		Returns struct {
			GetBucketPolicyOutput *s3.GetBucketPolicyOutput
			Error                 error
		}
		Stub func(context.Context, *s3.GetBucketPolicyInput, ...request.Option) (*s3.GetBucketPolicyOutput, error)
	}
	GetBucketReplicationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketReplicationInput *s3.GetBucketReplicationInput
		}
		Returns struct {
			GetBucketReplicationOutput *s3.GetBucketReplicationOutput
			Error                      error
		}
		Stub func(*s3.GetBucketReplicationInput) (*s3.GetBucketReplicationOutput, error)
	}
	GetBucketReplicationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketReplicationInput *s3.GetBucketReplicationInput
		}
		Returns struct {
			Request                    *request.Request
			GetBucketReplicationOutput *s3.GetBucketReplicationOutput
		}
		Stub func(*s3.GetBucketReplicationInput) (*request.Request, *s3.GetBucketReplicationOutput)
	}
	GetBucketReplicationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                   context.Context
			GetBucketReplicationInput *s3.GetBucketReplicationInput
			OptionSlice               []request.Option
		}
		Returns struct {
			GetBucketReplicationOutput *s3.GetBucketReplicationOutput
			Error                      error
		}
		Stub func(context.Context, *s3.GetBucketReplicationInput, ...request.Option) (*s3.GetBucketReplicationOutput, error)
	}
	GetBucketRequestPaymentCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketRequestPaymentInput *s3.GetBucketRequestPaymentInput
		}
		Returns struct {
			GetBucketRequestPaymentOutput *s3.GetBucketRequestPaymentOutput
			Error                         error
		}
		Stub func(*s3.GetBucketRequestPaymentInput) (*s3.GetBucketRequestPaymentOutput, error)
	}
	GetBucketRequestPaymentRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketRequestPaymentInput *s3.GetBucketRequestPaymentInput
		}
		Returns struct {
			Request                       *request.Request
			GetBucketRequestPaymentOutput *s3.GetBucketRequestPaymentOutput
		}
		Stub func(*s3.GetBucketRequestPaymentInput) (*request.Request, *s3.GetBucketRequestPaymentOutput)
	}
	GetBucketRequestPaymentWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                      context.Context
			GetBucketRequestPaymentInput *s3.GetBucketRequestPaymentInput
			OptionSlice                  []request.Option
		}
		Returns struct {
			GetBucketRequestPaymentOutput *s3.GetBucketRequestPaymentOutput
			Error                         error
		}
		Stub func(context.Context, *s3.GetBucketRequestPaymentInput, ...request.Option) (*s3.GetBucketRequestPaymentOutput, error)
	}
	GetBucketTaggingCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketTaggingInput *s3.GetBucketTaggingInput
		}
		Returns struct {
			GetBucketTaggingOutput *s3.GetBucketTaggingOutput
			Error                  error
		}
		Stub func(*s3.GetBucketTaggingInput) (*s3.GetBucketTaggingOutput, error)
	}
	GetBucketTaggingRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketTaggingInput *s3.GetBucketTaggingInput
		}
		Returns struct {
			Request                *request.Request
			GetBucketTaggingOutput *s3.GetBucketTaggingOutput
		}
		Stub func(*s3.GetBucketTaggingInput) (*request.Request, *s3.GetBucketTaggingOutput)
	}
	GetBucketTaggingWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context               context.Context
			GetBucketTaggingInput *s3.GetBucketTaggingInput
			OptionSlice           []request.Option
		}
		Returns struct {
			GetBucketTaggingOutput *s3.GetBucketTaggingOutput
			Error                  error
		}
		Stub func(context.Context, *s3.GetBucketTaggingInput, ...request.Option) (*s3.GetBucketTaggingOutput, error)
	}
	GetBucketVersioningCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketVersioningInput *s3.GetBucketVersioningInput
		}
		Returns struct {
			GetBucketVersioningOutput *s3.GetBucketVersioningOutput
			Error                     error
		}
		Stub func(*s3.GetBucketVersioningInput) (*s3.GetBucketVersioningOutput, error)
	}
	GetBucketVersioningRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketVersioningInput *s3.GetBucketVersioningInput
		}
		Returns struct {
			Request                   *request.Request
			GetBucketVersioningOutput *s3.GetBucketVersioningOutput
		}
		Stub func(*s3.GetBucketVersioningInput) (*request.Request, *s3.GetBucketVersioningOutput)
	}
	GetBucketVersioningWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                  context.Context
			GetBucketVersioningInput *s3.GetBucketVersioningInput
			OptionSlice              []request.Option
		}
		Returns struct {
			GetBucketVersioningOutput *s3.GetBucketVersioningOutput
			Error                     error
		}
		Stub func(context.Context, *s3.GetBucketVersioningInput, ...request.Option) (*s3.GetBucketVersioningOutput, error)
	}
	GetBucketWebsiteCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketWebsiteInput *s3.GetBucketWebsiteInput
		}
		Returns struct {
			GetBucketWebsiteOutput *s3.GetBucketWebsiteOutput
			Error                  error
		}
		Stub func(*s3.GetBucketWebsiteInput) (*s3.GetBucketWebsiteOutput, error)
	}
	GetBucketWebsiteRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetBucketWebsiteInput *s3.GetBucketWebsiteInput
		}
		Returns struct {
			Request                *request.Request
			GetBucketWebsiteOutput *s3.GetBucketWebsiteOutput
		}
		Stub func(*s3.GetBucketWebsiteInput) (*request.Request, *s3.GetBucketWebsiteOutput)
	}
	GetBucketWebsiteWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context               context.Context
			GetBucketWebsiteInput *s3.GetBucketWebsiteInput
			OptionSlice           []request.Option
		}
		Returns struct {
			GetBucketWebsiteOutput *s3.GetBucketWebsiteOutput
			Error                  error
		}
		Stub func(context.Context, *s3.GetBucketWebsiteInput, ...request.Option) (*s3.GetBucketWebsiteOutput, error)
	}
	GetObjectCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetObjectInput *s3.GetObjectInput
		}
		Returns struct {
			GetObjectOutput *s3.GetObjectOutput
			Error           error
		}
		Stub func(*s3.GetObjectInput) (*s3.GetObjectOutput, error)
	}
	GetObjectAclCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetObjectAclInput *s3.GetObjectAclInput
		}
		Returns struct {
			GetObjectAclOutput *s3.GetObjectAclOutput
			Error              error
		}
		Stub func(*s3.GetObjectAclInput) (*s3.GetObjectAclOutput, error)
	}
	GetObjectAclRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetObjectAclInput *s3.GetObjectAclInput
		}
		Returns struct {
			Request            *request.Request
			GetObjectAclOutput *s3.GetObjectAclOutput
		}
		Stub func(*s3.GetObjectAclInput) (*request.Request, *s3.GetObjectAclOutput)
	}
	GetObjectAclWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context           context.Context
			GetObjectAclInput *s3.GetObjectAclInput
			OptionSlice       []request.Option
		}
		Returns struct {
			GetObjectAclOutput *s3.GetObjectAclOutput
			Error              error
		}
		Stub func(context.Context, *s3.GetObjectAclInput, ...request.Option) (*s3.GetObjectAclOutput, error)
	}
	GetObjectAttributesCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetObjectAttributesInput *s3.GetObjectAttributesInput
		}
		Returns struct {
			GetObjectAttributesOutput *s3.GetObjectAttributesOutput
			Error                     error
		}
		Stub func(*s3.GetObjectAttributesInput) (*s3.GetObjectAttributesOutput, error)
	}
	GetObjectAttributesRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetObjectAttributesInput *s3.GetObjectAttributesInput
		}
		Returns struct {
			Request                   *request.Request
			GetObjectAttributesOutput *s3.GetObjectAttributesOutput
		}
		Stub func(*s3.GetObjectAttributesInput) (*request.Request, *s3.GetObjectAttributesOutput)
	}
	GetObjectAttributesWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                  context.Context
			GetObjectAttributesInput *s3.GetObjectAttributesInput
			OptionSlice              []request.Option
		}
		Returns struct {
			GetObjectAttributesOutput *s3.GetObjectAttributesOutput
			Error                     error
		}
		Stub func(context.Context, *s3.GetObjectAttributesInput, ...request.Option) (*s3.GetObjectAttributesOutput, error)
	}
	GetObjectLegalHoldCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetObjectLegalHoldInput *s3.GetObjectLegalHoldInput
		}
		Returns struct {
			GetObjectLegalHoldOutput *s3.GetObjectLegalHoldOutput
			Error                    error
		}
		Stub func(*s3.GetObjectLegalHoldInput) (*s3.GetObjectLegalHoldOutput, error)
	}
	GetObjectLegalHoldRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetObjectLegalHoldInput *s3.GetObjectLegalHoldInput
		}
		Returns struct {
			Request                  *request.Request
			GetObjectLegalHoldOutput *s3.GetObjectLegalHoldOutput
		}
		Stub func(*s3.GetObjectLegalHoldInput) (*request.Request, *s3.GetObjectLegalHoldOutput)
	}
	GetObjectLegalHoldWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                 context.Context
			GetObjectLegalHoldInput *s3.GetObjectLegalHoldInput
			OptionSlice             []request.Option
		}
		Returns struct {
			GetObjectLegalHoldOutput *s3.GetObjectLegalHoldOutput
			Error                    error
		}
		Stub func(context.Context, *s3.GetObjectLegalHoldInput, ...request.Option) (*s3.GetObjectLegalHoldOutput, error)
	}
	GetObjectLockConfigurationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetObjectLockConfigurationInput *s3.GetObjectLockConfigurationInput
		}
		Returns struct {
			GetObjectLockConfigurationOutput *s3.GetObjectLockConfigurationOutput
			Error                            error
		}
		Stub func(*s3.GetObjectLockConfigurationInput) (*s3.GetObjectLockConfigurationOutput, error)
	}
	GetObjectLockConfigurationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetObjectLockConfigurationInput *s3.GetObjectLockConfigurationInput
		}
		Returns struct {
			Request                          *request.Request
			GetObjectLockConfigurationOutput *s3.GetObjectLockConfigurationOutput
		}
		Stub func(*s3.GetObjectLockConfigurationInput) (*request.Request, *s3.GetObjectLockConfigurationOutput)
	}
	GetObjectLockConfigurationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                         context.Context
			GetObjectLockConfigurationInput *s3.GetObjectLockConfigurationInput
			OptionSlice                     []request.Option
		}
		Returns struct {
			GetObjectLockConfigurationOutput *s3.GetObjectLockConfigurationOutput
			Error                            error
		}
		Stub func(context.Context, *s3.GetObjectLockConfigurationInput, ...request.Option) (*s3.GetObjectLockConfigurationOutput, error)
	}
	GetObjectRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetObjectInput *s3.GetObjectInput
		}
		Returns struct {
			Request         *request.Request
			GetObjectOutput *s3.GetObjectOutput
		}
		Stub func(*s3.GetObjectInput) (*request.Request, *s3.GetObjectOutput)
	}
	GetObjectRetentionCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetObjectRetentionInput *s3.GetObjectRetentionInput
		}
		Returns struct {
			GetObjectRetentionOutput *s3.GetObjectRetentionOutput
			Error                    error
		}
		Stub func(*s3.GetObjectRetentionInput) (*s3.GetObjectRetentionOutput, error)
	}
	GetObjectRetentionRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetObjectRetentionInput *s3.GetObjectRetentionInput
		}
		Returns struct {
			Request                  *request.Request
			GetObjectRetentionOutput *s3.GetObjectRetentionOutput
		}
		Stub func(*s3.GetObjectRetentionInput) (*request.Request, *s3.GetObjectRetentionOutput)
	}
	GetObjectRetentionWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                 context.Context
			GetObjectRetentionInput *s3.GetObjectRetentionInput
			OptionSlice             []request.Option
		}
		Returns struct {
			GetObjectRetentionOutput *s3.GetObjectRetentionOutput
			Error                    error
		}
		Stub func(context.Context, *s3.GetObjectRetentionInput, ...request.Option) (*s3.GetObjectRetentionOutput, error)
	}
	GetObjectTaggingCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetObjectTaggingInput *s3.GetObjectTaggingInput
		}
		Returns struct {
			GetObjectTaggingOutput *s3.GetObjectTaggingOutput
			Error                  error
		}
		Stub func(*s3.GetObjectTaggingInput) (*s3.GetObjectTaggingOutput, error)
	}
	GetObjectTaggingRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetObjectTaggingInput *s3.GetObjectTaggingInput
		}
		Returns struct {
			Request                *request.Request
			GetObjectTaggingOutput *s3.GetObjectTaggingOutput
		}
		Stub func(*s3.GetObjectTaggingInput) (*request.Request, *s3.GetObjectTaggingOutput)
	}
	GetObjectTaggingWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context               context.Context
			GetObjectTaggingInput *s3.GetObjectTaggingInput
			OptionSlice           []request.Option
		}
		Returns struct {
			GetObjectTaggingOutput *s3.GetObjectTaggingOutput
			Error                  error
		}
		Stub func(context.Context, *s3.GetObjectTaggingInput, ...request.Option) (*s3.GetObjectTaggingOutput, error)
	}
	GetObjectTorrentCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetObjectTorrentInput *s3.GetObjectTorrentInput
		}
		Returns struct {
			GetObjectTorrentOutput *s3.GetObjectTorrentOutput
			Error                  error
		}
		Stub func(*s3.GetObjectTorrentInput) (*s3.GetObjectTorrentOutput, error)
	}
	GetObjectTorrentRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetObjectTorrentInput *s3.GetObjectTorrentInput
		}
		Returns struct {
			Request                *request.Request
			GetObjectTorrentOutput *s3.GetObjectTorrentOutput
		}
		Stub func(*s3.GetObjectTorrentInput) (*request.Request, *s3.GetObjectTorrentOutput)
	}
	GetObjectTorrentWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context               context.Context
			GetObjectTorrentInput *s3.GetObjectTorrentInput
			OptionSlice           []request.Option
		}
		Returns struct {
			GetObjectTorrentOutput *s3.GetObjectTorrentOutput
			Error                  error
		}
		Stub func(context.Context, *s3.GetObjectTorrentInput, ...request.Option) (*s3.GetObjectTorrentOutput, error)
	}
	GetObjectWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context        context.Context
			GetObjectInput *s3.GetObjectInput
			OptionSlice    []request.Option
		}
		Returns struct {
			GetObjectOutput *s3.GetObjectOutput
			Error           error
		}
		Stub func(context.Context, *s3.GetObjectInput, ...request.Option) (*s3.GetObjectOutput, error)
	}
	GetPublicAccessBlockCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetPublicAccessBlockInput *s3.GetPublicAccessBlockInput
		}
		Returns struct {
			GetPublicAccessBlockOutput *s3.GetPublicAccessBlockOutput
			Error                      error
		}
		Stub func(*s3.GetPublicAccessBlockInput) (*s3.GetPublicAccessBlockOutput, error)
	}
	GetPublicAccessBlockRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			GetPublicAccessBlockInput *s3.GetPublicAccessBlockInput
		}
		Returns struct {
			Request                    *request.Request
			GetPublicAccessBlockOutput *s3.GetPublicAccessBlockOutput
		}
		Stub func(*s3.GetPublicAccessBlockInput) (*request.Request, *s3.GetPublicAccessBlockOutput)
	}
	GetPublicAccessBlockWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                   context.Context
			GetPublicAccessBlockInput *s3.GetPublicAccessBlockInput
			OptionSlice               []request.Option
		}
		Returns struct {
			GetPublicAccessBlockOutput *s3.GetPublicAccessBlockOutput
			Error                      error
		}
		Stub func(context.Context, *s3.GetPublicAccessBlockInput, ...request.Option) (*s3.GetPublicAccessBlockOutput, error)
	}
	HeadBucketCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			HeadBucketInput *s3.HeadBucketInput
		}
		Returns struct {
			HeadBucketOutput *s3.HeadBucketOutput
			Error            error
		}
		Stub func(*s3.HeadBucketInput) (*s3.HeadBucketOutput, error)
	}
	HeadBucketRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			HeadBucketInput *s3.HeadBucketInput
		}
		Returns struct {
			Request          *request.Request
			HeadBucketOutput *s3.HeadBucketOutput
		}
		Stub func(*s3.HeadBucketInput) (*request.Request, *s3.HeadBucketOutput)
	}
	HeadBucketWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context         context.Context
			HeadBucketInput *s3.HeadBucketInput
			OptionSlice     []request.Option
		}
		Returns struct {
			HeadBucketOutput *s3.HeadBucketOutput
			Error            error
		}
		Stub func(context.Context, *s3.HeadBucketInput, ...request.Option) (*s3.HeadBucketOutput, error)
	}
	HeadObjectCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			HeadObjectInput *s3.HeadObjectInput
		}
		Returns struct {
			HeadObjectOutput *s3.HeadObjectOutput
			Error            error
		}
		Stub func(*s3.HeadObjectInput) (*s3.HeadObjectOutput, error)
	}
	HeadObjectRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			HeadObjectInput *s3.HeadObjectInput
		}
		Returns struct {
			Request          *request.Request
			HeadObjectOutput *s3.HeadObjectOutput
		}
		Stub func(*s3.HeadObjectInput) (*request.Request, *s3.HeadObjectOutput)
	}
	HeadObjectWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context         context.Context
			HeadObjectInput *s3.HeadObjectInput
			OptionSlice     []request.Option
		}
		Returns struct {
			HeadObjectOutput *s3.HeadObjectOutput
			Error            error
		}
		Stub func(context.Context, *s3.HeadObjectInput, ...request.Option) (*s3.HeadObjectOutput, error)
	}
	ListBucketAnalyticsConfigurationsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListBucketAnalyticsConfigurationsInput *s3.ListBucketAnalyticsConfigurationsInput
		}
		Returns struct {
			ListBucketAnalyticsConfigurationsOutput *s3.ListBucketAnalyticsConfigurationsOutput
			Error                                   error
		}
		Stub func(*s3.ListBucketAnalyticsConfigurationsInput) (*s3.ListBucketAnalyticsConfigurationsOutput, error)
	}
	ListBucketAnalyticsConfigurationsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListBucketAnalyticsConfigurationsInput *s3.ListBucketAnalyticsConfigurationsInput
		}
		Returns struct {
			Request                                 *request.Request
			ListBucketAnalyticsConfigurationsOutput *s3.ListBucketAnalyticsConfigurationsOutput
		}
		Stub func(*s3.ListBucketAnalyticsConfigurationsInput) (*request.Request, *s3.ListBucketAnalyticsConfigurationsOutput)
	}
	ListBucketAnalyticsConfigurationsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                                context.Context
			ListBucketAnalyticsConfigurationsInput *s3.ListBucketAnalyticsConfigurationsInput
			OptionSlice                            []request.Option
		}
		Returns struct {
			ListBucketAnalyticsConfigurationsOutput *s3.ListBucketAnalyticsConfigurationsOutput
			Error                                   error
		}
		Stub func(context.Context, *s3.ListBucketAnalyticsConfigurationsInput, ...request.Option) (*s3.ListBucketAnalyticsConfigurationsOutput, error)
	}
	ListBucketIntelligentTieringConfigurationsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListBucketIntelligentTieringConfigurationsInput *s3.ListBucketIntelligentTieringConfigurationsInput
		}
		Returns struct {
			ListBucketIntelligentTieringConfigurationsOutput *s3.ListBucketIntelligentTieringConfigurationsOutput
			Error                                            error
		}
		Stub func(*s3.ListBucketIntelligentTieringConfigurationsInput) (*s3.ListBucketIntelligentTieringConfigurationsOutput, error)
	}
	ListBucketIntelligentTieringConfigurationsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListBucketIntelligentTieringConfigurationsInput *s3.ListBucketIntelligentTieringConfigurationsInput
		}
		Returns struct {
			Request                                          *request.Request
			ListBucketIntelligentTieringConfigurationsOutput *s3.ListBucketIntelligentTieringConfigurationsOutput
		}
		Stub func(*s3.ListBucketIntelligentTieringConfigurationsInput) (*request.Request, *s3.ListBucketIntelligentTieringConfigurationsOutput)
	}
	ListBucketIntelligentTieringConfigurationsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                                         context.Context
			ListBucketIntelligentTieringConfigurationsInput *s3.ListBucketIntelligentTieringConfigurationsInput
			OptionSlice                                     []request.Option
		}
		Returns struct {
			ListBucketIntelligentTieringConfigurationsOutput *s3.ListBucketIntelligentTieringConfigurationsOutput
			Error                                            error
		}
		Stub func(context.Context, *s3.ListBucketIntelligentTieringConfigurationsInput, ...request.Option) (*s3.ListBucketIntelligentTieringConfigurationsOutput, error)
	}
	ListBucketInventoryConfigurationsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListBucketInventoryConfigurationsInput *s3.ListBucketInventoryConfigurationsInput
		}
		Returns struct {
			ListBucketInventoryConfigurationsOutput *s3.ListBucketInventoryConfigurationsOutput
			Error                                   error
		}
		Stub func(*s3.ListBucketInventoryConfigurationsInput) (*s3.ListBucketInventoryConfigurationsOutput, error)
	}
	ListBucketInventoryConfigurationsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListBucketInventoryConfigurationsInput *s3.ListBucketInventoryConfigurationsInput
		}
		Returns struct {
			Request                                 *request.Request
			ListBucketInventoryConfigurationsOutput *s3.ListBucketInventoryConfigurationsOutput
		}
		Stub func(*s3.ListBucketInventoryConfigurationsInput) (*request.Request, *s3.ListBucketInventoryConfigurationsOutput)
	}
	ListBucketInventoryConfigurationsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                                context.Context
			ListBucketInventoryConfigurationsInput *s3.ListBucketInventoryConfigurationsInput
			OptionSlice                            []request.Option
		}
		Returns struct {
			ListBucketInventoryConfigurationsOutput *s3.ListBucketInventoryConfigurationsOutput
			Error                                   error
		}
		Stub func(context.Context, *s3.ListBucketInventoryConfigurationsInput, ...request.Option) (*s3.ListBucketInventoryConfigurationsOutput, error)
	}
	ListBucketMetricsConfigurationsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListBucketMetricsConfigurationsInput *s3.ListBucketMetricsConfigurationsInput
		}
		Returns struct {
			ListBucketMetricsConfigurationsOutput *s3.ListBucketMetricsConfigurationsOutput
			Error                                 error
		}
		Stub func(*s3.ListBucketMetricsConfigurationsInput) (*s3.ListBucketMetricsConfigurationsOutput, error)
	}
	ListBucketMetricsConfigurationsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListBucketMetricsConfigurationsInput *s3.ListBucketMetricsConfigurationsInput
		}
		Returns struct {
			Request                               *request.Request
			ListBucketMetricsConfigurationsOutput *s3.ListBucketMetricsConfigurationsOutput
		}
		Stub func(*s3.ListBucketMetricsConfigurationsInput) (*request.Request, *s3.ListBucketMetricsConfigurationsOutput)
	}
	ListBucketMetricsConfigurationsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                              context.Context
			ListBucketMetricsConfigurationsInput *s3.ListBucketMetricsConfigurationsInput
			OptionSlice                          []request.Option
		}
		Returns struct {
			ListBucketMetricsConfigurationsOutput *s3.ListBucketMetricsConfigurationsOutput
			Error                                 error
		}
		Stub func(context.Context, *s3.ListBucketMetricsConfigurationsInput, ...request.Option) (*s3.ListBucketMetricsConfigurationsOutput, error)
	}
	ListBucketsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListBucketsInput *s3.ListBucketsInput
		}
		Returns struct {
			ListBucketsOutput *s3.ListBucketsOutput
			Error             error
		}
		Stub func(*s3.ListBucketsInput) (*s3.ListBucketsOutput, error)
	}
	ListBucketsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListBucketsInput *s3.ListBucketsInput
		}
		Returns struct {
			Request           *request.Request
			ListBucketsOutput *s3.ListBucketsOutput
		}
		Stub func(*s3.ListBucketsInput) (*request.Request, *s3.ListBucketsOutput)
	}
	ListBucketsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context          context.Context
			ListBucketsInput *s3.ListBucketsInput
			OptionSlice      []request.Option
		}
		Returns struct {
			ListBucketsOutput *s3.ListBucketsOutput
			Error             error
		}
		Stub func(context.Context, *s3.ListBucketsInput, ...request.Option) (*s3.ListBucketsOutput, error)
	}
	ListMultipartUploadsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListMultipartUploadsInput *s3.ListMultipartUploadsInput
		}
		Returns struct {
			ListMultipartUploadsOutput *s3.ListMultipartUploadsOutput
			Error                      error
		}
		Stub func(*s3.ListMultipartUploadsInput) (*s3.ListMultipartUploadsOutput, error)
	}
	ListMultipartUploadsPagesCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListMultipartUploadsInput                *s3.ListMultipartUploadsInput
			FuncS3ListMultipartUploadsOutputBoolBool func(*s3.ListMultipartUploadsOutput, bool) bool
		}
		Returns struct {
			Error error
		}
		Stub func(*s3.ListMultipartUploadsInput, func(*s3.ListMultipartUploadsOutput, bool) bool) error
	}
	ListMultipartUploadsPagesWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                                  context.Context
			ListMultipartUploadsInput                *s3.ListMultipartUploadsInput
			FuncS3ListMultipartUploadsOutputBoolBool func(*s3.ListMultipartUploadsOutput, bool) bool
			OptionSlice                              []request.Option
		}
		Returns struct {
			Error error
		}
		Stub func(context.Context, *s3.ListMultipartUploadsInput, func(*s3.ListMultipartUploadsOutput, bool) bool, ...request.Option) error
	}
	ListMultipartUploadsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListMultipartUploadsInput *s3.ListMultipartUploadsInput
		}
		Returns struct {
			Request                    *request.Request
			ListMultipartUploadsOutput *s3.ListMultipartUploadsOutput
		}
		Stub func(*s3.ListMultipartUploadsInput) (*request.Request, *s3.ListMultipartUploadsOutput)
	}
	ListMultipartUploadsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                   context.Context
			ListMultipartUploadsInput *s3.ListMultipartUploadsInput
			OptionSlice               []request.Option
		}
		Returns struct {
			ListMultipartUploadsOutput *s3.ListMultipartUploadsOutput
			Error                      error
		}
		Stub func(context.Context, *s3.ListMultipartUploadsInput, ...request.Option) (*s3.ListMultipartUploadsOutput, error)
	}
	ListObjectVersionsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListObjectVersionsInput *s3.ListObjectVersionsInput
		}
		Returns struct {
			ListObjectVersionsOutput *s3.ListObjectVersionsOutput
			Error                    error
		}
		Stub func(*s3.ListObjectVersionsInput) (*s3.ListObjectVersionsOutput, error)
	}
	ListObjectVersionsPagesCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListObjectVersionsInput                *s3.ListObjectVersionsInput
			FuncS3ListObjectVersionsOutputBoolBool func(*s3.ListObjectVersionsOutput, bool) bool
		}
		Returns struct {
			Error error
		}
		Stub func(*s3.ListObjectVersionsInput, func(*s3.ListObjectVersionsOutput, bool) bool) error
	}
	ListObjectVersionsPagesWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                                context.Context
			ListObjectVersionsInput                *s3.ListObjectVersionsInput
			FuncS3ListObjectVersionsOutputBoolBool func(*s3.ListObjectVersionsOutput, bool) bool
			OptionSlice                            []request.Option
		}
		Returns struct {
			Error error
		}
		Stub func(context.Context, *s3.ListObjectVersionsInput, func(*s3.ListObjectVersionsOutput, bool) bool, ...request.Option) error
	}
	ListObjectVersionsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListObjectVersionsInput *s3.ListObjectVersionsInput
		}
		Returns struct {
			Request                  *request.Request
			ListObjectVersionsOutput *s3.ListObjectVersionsOutput
		}
		Stub func(*s3.ListObjectVersionsInput) (*request.Request, *s3.ListObjectVersionsOutput)
	}
	ListObjectVersionsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                 context.Context
			ListObjectVersionsInput *s3.ListObjectVersionsInput
			OptionSlice             []request.Option
		}
		Returns struct {
			ListObjectVersionsOutput *s3.ListObjectVersionsOutput
			Error                    error
		}
		Stub func(context.Context, *s3.ListObjectVersionsInput, ...request.Option) (*s3.ListObjectVersionsOutput, error)
	}
	ListObjectsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListObjectsInput *s3.ListObjectsInput
		}
		Returns struct {
			ListObjectsOutput *s3.ListObjectsOutput
			Error             error
		}
		Stub func(*s3.ListObjectsInput) (*s3.ListObjectsOutput, error)
	}
	ListObjectsPagesCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListObjectsInput                *s3.ListObjectsInput
			FuncS3ListObjectsOutputBoolBool func(*s3.ListObjectsOutput, bool) bool
		}
		Returns struct {
			Error error
		}
		Stub func(*s3.ListObjectsInput, func(*s3.ListObjectsOutput, bool) bool) error
	}
	ListObjectsPagesWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                         context.Context
			ListObjectsInput                *s3.ListObjectsInput
			FuncS3ListObjectsOutputBoolBool func(*s3.ListObjectsOutput, bool) bool
			OptionSlice                     []request.Option
		}
		Returns struct {
			Error error
		}
		Stub func(context.Context, *s3.ListObjectsInput, func(*s3.ListObjectsOutput, bool) bool, ...request.Option) error
	}
	ListObjectsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListObjectsInput *s3.ListObjectsInput
		}
		Returns struct {
			Request           *request.Request
			ListObjectsOutput *s3.ListObjectsOutput
		}
		Stub func(*s3.ListObjectsInput) (*request.Request, *s3.ListObjectsOutput)
	}
	ListObjectsV2Call struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListObjectsV2Input *s3.ListObjectsV2Input
		}
		Returns struct {
			ListObjectsV2Output *s3.ListObjectsV2Output
			Error               error
		}
		Stub func(*s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error)
	}
	ListObjectsV2PagesCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListObjectsV2Input                *s3.ListObjectsV2Input
			FuncS3ListObjectsV2OutputBoolBool func(*s3.ListObjectsV2Output, bool) bool
		}
		Returns struct {
			Error error
		}
		Stub func(*s3.ListObjectsV2Input, func(*s3.ListObjectsV2Output, bool) bool) error
	}
	ListObjectsV2PagesWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                           context.Context
			ListObjectsV2Input                *s3.ListObjectsV2Input
			FuncS3ListObjectsV2OutputBoolBool func(*s3.ListObjectsV2Output, bool) bool
			OptionSlice                       []request.Option
		}
		Returns struct {
			Error error
		}
		Stub func(context.Context, *s3.ListObjectsV2Input, func(*s3.ListObjectsV2Output, bool) bool, ...request.Option) error
	}
	ListObjectsV2RequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListObjectsV2Input *s3.ListObjectsV2Input
		}
		Returns struct {
			Request             *request.Request
			ListObjectsV2Output *s3.ListObjectsV2Output
		}
		Stub func(*s3.ListObjectsV2Input) (*request.Request, *s3.ListObjectsV2Output)
	}
	ListObjectsV2WithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context            context.Context
			ListObjectsV2Input *s3.ListObjectsV2Input
			OptionSlice        []request.Option
		}
		Returns struct {
			ListObjectsV2Output *s3.ListObjectsV2Output
			Error               error
		}
		Stub func(context.Context, *s3.ListObjectsV2Input, ...request.Option) (*s3.ListObjectsV2Output, error)
	}
	ListObjectsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context          context.Context
			ListObjectsInput *s3.ListObjectsInput
			OptionSlice      []request.Option
		}
		Returns struct {
			ListObjectsOutput *s3.ListObjectsOutput
			Error             error
		}
		Stub func(context.Context, *s3.ListObjectsInput, ...request.Option) (*s3.ListObjectsOutput, error)
	}
	ListPartsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListPartsInput *s3.ListPartsInput
		}
		Returns struct {
			ListPartsOutput *s3.ListPartsOutput
			Error           error
		}
		Stub func(*s3.ListPartsInput) (*s3.ListPartsOutput, error)
	}
	ListPartsPagesCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListPartsInput                *s3.ListPartsInput
			FuncS3ListPartsOutputBoolBool func(*s3.ListPartsOutput, bool) bool
		}
		Returns struct {
			Error error
		}
		Stub func(*s3.ListPartsInput, func(*s3.ListPartsOutput, bool) bool) error
	}
	ListPartsPagesWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                       context.Context
			ListPartsInput                *s3.ListPartsInput
			FuncS3ListPartsOutputBoolBool func(*s3.ListPartsOutput, bool) bool
			OptionSlice                   []request.Option
		}
		Returns struct {
			Error error
		}
		Stub func(context.Context, *s3.ListPartsInput, func(*s3.ListPartsOutput, bool) bool, ...request.Option) error
	}
	ListPartsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ListPartsInput *s3.ListPartsInput
		}
		Returns struct {
			Request         *request.Request
			ListPartsOutput *s3.ListPartsOutput
		}
		Stub func(*s3.ListPartsInput) (*request.Request, *s3.ListPartsOutput)
	}
	ListPartsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context        context.Context
			ListPartsInput *s3.ListPartsInput
			OptionSlice    []request.Option
		}
		Returns struct {
			ListPartsOutput *s3.ListPartsOutput
			Error           error
		}
		Stub func(context.Context, *s3.ListPartsInput, ...request.Option) (*s3.ListPartsOutput, error)
	}
	PutBucketAccelerateConfigurationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketAccelerateConfigurationInput *s3.PutBucketAccelerateConfigurationInput
		}
		Returns struct {
			PutBucketAccelerateConfigurationOutput *s3.PutBucketAccelerateConfigurationOutput
			Error                                  error
		}
		Stub func(*s3.PutBucketAccelerateConfigurationInput) (*s3.PutBucketAccelerateConfigurationOutput, error)
	}
	PutBucketAccelerateConfigurationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketAccelerateConfigurationInput *s3.PutBucketAccelerateConfigurationInput
		}
		Returns struct {
			Request                                *request.Request
			PutBucketAccelerateConfigurationOutput *s3.PutBucketAccelerateConfigurationOutput
		}
		Stub func(*s3.PutBucketAccelerateConfigurationInput) (*request.Request, *s3.PutBucketAccelerateConfigurationOutput)
	}
	PutBucketAccelerateConfigurationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                               context.Context
			PutBucketAccelerateConfigurationInput *s3.PutBucketAccelerateConfigurationInput
			OptionSlice                           []request.Option
		}
		Returns struct {
			PutBucketAccelerateConfigurationOutput *s3.PutBucketAccelerateConfigurationOutput
			Error                                  error
		}
		Stub func(context.Context, *s3.PutBucketAccelerateConfigurationInput, ...request.Option) (*s3.PutBucketAccelerateConfigurationOutput, error)
	}
	PutBucketAclCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketAclInput *s3.PutBucketAclInput
		}
		Returns struct {
			PutBucketAclOutput *s3.PutBucketAclOutput
			Error              error
		}
		Stub func(*s3.PutBucketAclInput) (*s3.PutBucketAclOutput, error)
	}
	PutBucketAclRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketAclInput *s3.PutBucketAclInput
		}
		Returns struct {
			Request            *request.Request
			PutBucketAclOutput *s3.PutBucketAclOutput
		}
		Stub func(*s3.PutBucketAclInput) (*request.Request, *s3.PutBucketAclOutput)
	}
	PutBucketAclWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context           context.Context
			PutBucketAclInput *s3.PutBucketAclInput
			OptionSlice       []request.Option
		}
		Returns struct {
			PutBucketAclOutput *s3.PutBucketAclOutput
			Error              error
		}
		Stub func(context.Context, *s3.PutBucketAclInput, ...request.Option) (*s3.PutBucketAclOutput, error)
	}
	PutBucketAnalyticsConfigurationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketAnalyticsConfigurationInput *s3.PutBucketAnalyticsConfigurationInput
		}
		Returns struct {
			PutBucketAnalyticsConfigurationOutput *s3.PutBucketAnalyticsConfigurationOutput
			Error                                 error
		}
		Stub func(*s3.PutBucketAnalyticsConfigurationInput) (*s3.PutBucketAnalyticsConfigurationOutput, error)
	}
	PutBucketAnalyticsConfigurationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketAnalyticsConfigurationInput *s3.PutBucketAnalyticsConfigurationInput
		}
		Returns struct {
			Request                               *request.Request
			PutBucketAnalyticsConfigurationOutput *s3.PutBucketAnalyticsConfigurationOutput
		}
		Stub func(*s3.PutBucketAnalyticsConfigurationInput) (*request.Request, *s3.PutBucketAnalyticsConfigurationOutput)
	}
	PutBucketAnalyticsConfigurationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                              context.Context
			PutBucketAnalyticsConfigurationInput *s3.PutBucketAnalyticsConfigurationInput
			OptionSlice                          []request.Option
		}
		Returns struct {
			PutBucketAnalyticsConfigurationOutput *s3.PutBucketAnalyticsConfigurationOutput
			Error                                 error
		}
		Stub func(context.Context, *s3.PutBucketAnalyticsConfigurationInput, ...request.Option) (*s3.PutBucketAnalyticsConfigurationOutput, error)
	}
	PutBucketCorsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketCorsInput *s3.PutBucketCorsInput
		}
		Returns struct {
			PutBucketCorsOutput *s3.PutBucketCorsOutput
			Error               error
		}
		Stub func(*s3.PutBucketCorsInput) (*s3.PutBucketCorsOutput, error)
	}
	PutBucketCorsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketCorsInput *s3.PutBucketCorsInput
		}
		Returns struct {
			Request             *request.Request
			PutBucketCorsOutput *s3.PutBucketCorsOutput
		}
		Stub func(*s3.PutBucketCorsInput) (*request.Request, *s3.PutBucketCorsOutput)
	}
	PutBucketCorsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context            context.Context
			PutBucketCorsInput *s3.PutBucketCorsInput
			OptionSlice        []request.Option
		}
		Returns struct {
			PutBucketCorsOutput *s3.PutBucketCorsOutput
			Error               error
		}
		Stub func(context.Context, *s3.PutBucketCorsInput, ...request.Option) (*s3.PutBucketCorsOutput, error)
	}
	PutBucketEncryptionCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketEncryptionInput *s3.PutBucketEncryptionInput
		}
		Returns struct {
			PutBucketEncryptionOutput *s3.PutBucketEncryptionOutput
			Error                     error
		}
		Stub func(*s3.PutBucketEncryptionInput) (*s3.PutBucketEncryptionOutput, error)
	}
	PutBucketEncryptionRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketEncryptionInput *s3.PutBucketEncryptionInput
		}
		Returns struct {
			Request                   *request.Request
			PutBucketEncryptionOutput *s3.PutBucketEncryptionOutput
		}
		Stub func(*s3.PutBucketEncryptionInput) (*request.Request, *s3.PutBucketEncryptionOutput)
	}
	PutBucketEncryptionWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                  context.Context
			PutBucketEncryptionInput *s3.PutBucketEncryptionInput
			OptionSlice              []request.Option
		}
		Returns struct {
			PutBucketEncryptionOutput *s3.PutBucketEncryptionOutput
			Error                     error
		}
		Stub func(context.Context, *s3.PutBucketEncryptionInput, ...request.Option) (*s3.PutBucketEncryptionOutput, error)
	}
	PutBucketIntelligentTieringConfigurationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketIntelligentTieringConfigurationInput *s3.PutBucketIntelligentTieringConfigurationInput
		}
		Returns struct {
			PutBucketIntelligentTieringConfigurationOutput *s3.PutBucketIntelligentTieringConfigurationOutput
			Error                                          error
		}
		Stub func(*s3.PutBucketIntelligentTieringConfigurationInput) (*s3.PutBucketIntelligentTieringConfigurationOutput, error)
	}
	PutBucketIntelligentTieringConfigurationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketIntelligentTieringConfigurationInput *s3.PutBucketIntelligentTieringConfigurationInput
		}
		Returns struct {
			Request                                        *request.Request
			PutBucketIntelligentTieringConfigurationOutput *s3.PutBucketIntelligentTieringConfigurationOutput
		}
		Stub func(*s3.PutBucketIntelligentTieringConfigurationInput) (*request.Request, *s3.PutBucketIntelligentTieringConfigurationOutput)
	}
	PutBucketIntelligentTieringConfigurationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                                       context.Context
			PutBucketIntelligentTieringConfigurationInput *s3.PutBucketIntelligentTieringConfigurationInput
			OptionSlice                                   []request.Option
		}
		Returns struct {
			PutBucketIntelligentTieringConfigurationOutput *s3.PutBucketIntelligentTieringConfigurationOutput
			Error                                          error
		}
		Stub func(context.Context, *s3.PutBucketIntelligentTieringConfigurationInput, ...request.Option) (*s3.PutBucketIntelligentTieringConfigurationOutput, error)
	}
	PutBucketInventoryConfigurationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketInventoryConfigurationInput *s3.PutBucketInventoryConfigurationInput
		}
		Returns struct {
			PutBucketInventoryConfigurationOutput *s3.PutBucketInventoryConfigurationOutput
			Error                                 error
		}
		Stub func(*s3.PutBucketInventoryConfigurationInput) (*s3.PutBucketInventoryConfigurationOutput, error)
	}
	PutBucketInventoryConfigurationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketInventoryConfigurationInput *s3.PutBucketInventoryConfigurationInput
		}
		Returns struct {
			Request                               *request.Request
			PutBucketInventoryConfigurationOutput *s3.PutBucketInventoryConfigurationOutput
		}
		Stub func(*s3.PutBucketInventoryConfigurationInput) (*request.Request, *s3.PutBucketInventoryConfigurationOutput)
	}
	PutBucketInventoryConfigurationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                              context.Context
			PutBucketInventoryConfigurationInput *s3.PutBucketInventoryConfigurationInput
			OptionSlice                          []request.Option
		}
		Returns struct {
			PutBucketInventoryConfigurationOutput *s3.PutBucketInventoryConfigurationOutput
			Error                                 error
		}
		Stub func(context.Context, *s3.PutBucketInventoryConfigurationInput, ...request.Option) (*s3.PutBucketInventoryConfigurationOutput, error)
	}
	PutBucketLifecycleCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketLifecycleInput *s3.PutBucketLifecycleInput
		}
		Returns struct {
			PutBucketLifecycleOutput *s3.PutBucketLifecycleOutput
			Error                    error
		}
		Stub func(*s3.PutBucketLifecycleInput) (*s3.PutBucketLifecycleOutput, error)
	}
	PutBucketLifecycleConfigurationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketLifecycleConfigurationInput *s3.PutBucketLifecycleConfigurationInput
		}
		Returns struct {
			PutBucketLifecycleConfigurationOutput *s3.PutBucketLifecycleConfigurationOutput
			Error                                 error
		}
		Stub func(*s3.PutBucketLifecycleConfigurationInput) (*s3.PutBucketLifecycleConfigurationOutput, error)
	}
	PutBucketLifecycleConfigurationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketLifecycleConfigurationInput *s3.PutBucketLifecycleConfigurationInput
		}
		Returns struct {
			Request                               *request.Request
			PutBucketLifecycleConfigurationOutput *s3.PutBucketLifecycleConfigurationOutput
		}
		Stub func(*s3.PutBucketLifecycleConfigurationInput) (*request.Request, *s3.PutBucketLifecycleConfigurationOutput)
	}
	PutBucketLifecycleConfigurationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                              context.Context
			PutBucketLifecycleConfigurationInput *s3.PutBucketLifecycleConfigurationInput
			OptionSlice                          []request.Option
		}
		Returns struct {
			PutBucketLifecycleConfigurationOutput *s3.PutBucketLifecycleConfigurationOutput
			Error                                 error
		}
		Stub func(context.Context, *s3.PutBucketLifecycleConfigurationInput, ...request.Option) (*s3.PutBucketLifecycleConfigurationOutput, error)
	}
	PutBucketLifecycleRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketLifecycleInput *s3.PutBucketLifecycleInput
		}
		Returns struct {
			Request                  *request.Request
			PutBucketLifecycleOutput *s3.PutBucketLifecycleOutput
		}
		Stub func(*s3.PutBucketLifecycleInput) (*request.Request, *s3.PutBucketLifecycleOutput)
	}
	PutBucketLifecycleWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                 context.Context
			PutBucketLifecycleInput *s3.PutBucketLifecycleInput
			OptionSlice             []request.Option
		}
		Returns struct {
			PutBucketLifecycleOutput *s3.PutBucketLifecycleOutput
			Error                    error
		}
		Stub func(context.Context, *s3.PutBucketLifecycleInput, ...request.Option) (*s3.PutBucketLifecycleOutput, error)
	}
	PutBucketLoggingCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketLoggingInput *s3.PutBucketLoggingInput
		}
		Returns struct {
			PutBucketLoggingOutput *s3.PutBucketLoggingOutput
			Error                  error
		}
		Stub func(*s3.PutBucketLoggingInput) (*s3.PutBucketLoggingOutput, error)
	}
	PutBucketLoggingRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketLoggingInput *s3.PutBucketLoggingInput
		}
		Returns struct {
			Request                *request.Request
			PutBucketLoggingOutput *s3.PutBucketLoggingOutput
		}
		Stub func(*s3.PutBucketLoggingInput) (*request.Request, *s3.PutBucketLoggingOutput)
	}
	PutBucketLoggingWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context               context.Context
			PutBucketLoggingInput *s3.PutBucketLoggingInput
			OptionSlice           []request.Option
		}
		Returns struct {
			PutBucketLoggingOutput *s3.PutBucketLoggingOutput
			Error                  error
		}
		Stub func(context.Context, *s3.PutBucketLoggingInput, ...request.Option) (*s3.PutBucketLoggingOutput, error)
	}
	PutBucketMetricsConfigurationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketMetricsConfigurationInput *s3.PutBucketMetricsConfigurationInput
		}
		Returns struct {
			PutBucketMetricsConfigurationOutput *s3.PutBucketMetricsConfigurationOutput
			Error                               error
		}
		Stub func(*s3.PutBucketMetricsConfigurationInput) (*s3.PutBucketMetricsConfigurationOutput, error)
	}
	PutBucketMetricsConfigurationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketMetricsConfigurationInput *s3.PutBucketMetricsConfigurationInput
		}
		Returns struct {
			Request                             *request.Request
			PutBucketMetricsConfigurationOutput *s3.PutBucketMetricsConfigurationOutput
		}
		Stub func(*s3.PutBucketMetricsConfigurationInput) (*request.Request, *s3.PutBucketMetricsConfigurationOutput)
	}
	PutBucketMetricsConfigurationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                            context.Context
			PutBucketMetricsConfigurationInput *s3.PutBucketMetricsConfigurationInput
			OptionSlice                        []request.Option
		}
		Returns struct {
			PutBucketMetricsConfigurationOutput *s3.PutBucketMetricsConfigurationOutput
			Error                               error
		}
		Stub func(context.Context, *s3.PutBucketMetricsConfigurationInput, ...request.Option) (*s3.PutBucketMetricsConfigurationOutput, error)
	}
	PutBucketNotificationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketNotificationInput *s3.PutBucketNotificationInput
		}
		Returns struct {
			PutBucketNotificationOutput *s3.PutBucketNotificationOutput
			Error                       error
		}
		Stub func(*s3.PutBucketNotificationInput) (*s3.PutBucketNotificationOutput, error)
	}
	PutBucketNotificationConfigurationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketNotificationConfigurationInput *s3.PutBucketNotificationConfigurationInput
		}
		Returns struct {
			PutBucketNotificationConfigurationOutput *s3.PutBucketNotificationConfigurationOutput
			Error                                    error
		}
		Stub func(*s3.PutBucketNotificationConfigurationInput) (*s3.PutBucketNotificationConfigurationOutput, error)
	}
	PutBucketNotificationConfigurationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketNotificationConfigurationInput *s3.PutBucketNotificationConfigurationInput
		}
		Returns struct {
			Request                                  *request.Request
			PutBucketNotificationConfigurationOutput *s3.PutBucketNotificationConfigurationOutput
		}
		Stub func(*s3.PutBucketNotificationConfigurationInput) (*request.Request, *s3.PutBucketNotificationConfigurationOutput)
	}
	PutBucketNotificationConfigurationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                                 context.Context
			PutBucketNotificationConfigurationInput *s3.PutBucketNotificationConfigurationInput
			OptionSlice                             []request.Option
		}
		Returns struct {
			PutBucketNotificationConfigurationOutput *s3.PutBucketNotificationConfigurationOutput
			Error                                    error
		}
		Stub func(context.Context, *s3.PutBucketNotificationConfigurationInput, ...request.Option) (*s3.PutBucketNotificationConfigurationOutput, error)
	}
	PutBucketNotificationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketNotificationInput *s3.PutBucketNotificationInput
		}
		Returns struct {
			Request                     *request.Request
			PutBucketNotificationOutput *s3.PutBucketNotificationOutput
		}
		Stub func(*s3.PutBucketNotificationInput) (*request.Request, *s3.PutBucketNotificationOutput)
	}
	PutBucketNotificationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                    context.Context
			PutBucketNotificationInput *s3.PutBucketNotificationInput
			OptionSlice                []request.Option
		}
		Returns struct {
			PutBucketNotificationOutput *s3.PutBucketNotificationOutput
			Error                       error
		}
		Stub func(context.Context, *s3.PutBucketNotificationInput, ...request.Option) (*s3.PutBucketNotificationOutput, error)
	}
	PutBucketOwnershipControlsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketOwnershipControlsInput *s3.PutBucketOwnershipControlsInput
		}
		Returns struct {
			PutBucketOwnershipControlsOutput *s3.PutBucketOwnershipControlsOutput
			Error                            error
		}
		Stub func(*s3.PutBucketOwnershipControlsInput) (*s3.PutBucketOwnershipControlsOutput, error)
	}
	PutBucketOwnershipControlsRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketOwnershipControlsInput *s3.PutBucketOwnershipControlsInput
		}
		Returns struct {
			Request                          *request.Request
			PutBucketOwnershipControlsOutput *s3.PutBucketOwnershipControlsOutput
		}
		Stub func(*s3.PutBucketOwnershipControlsInput) (*request.Request, *s3.PutBucketOwnershipControlsOutput)
	}
	PutBucketOwnershipControlsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                         context.Context
			PutBucketOwnershipControlsInput *s3.PutBucketOwnershipControlsInput
			OptionSlice                     []request.Option
		}
		Returns struct {
			PutBucketOwnershipControlsOutput *s3.PutBucketOwnershipControlsOutput
			Error                            error
		}
		Stub func(context.Context, *s3.PutBucketOwnershipControlsInput, ...request.Option) (*s3.PutBucketOwnershipControlsOutput, error)
	}
	PutBucketPolicyCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketPolicyInput *s3.PutBucketPolicyInput
		}
		Returns struct {
			PutBucketPolicyOutput *s3.PutBucketPolicyOutput
			Error                 error
		}
		Stub func(*s3.PutBucketPolicyInput) (*s3.PutBucketPolicyOutput, error)
	}
	PutBucketPolicyRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketPolicyInput *s3.PutBucketPolicyInput
		}
		Returns struct {
			Request               *request.Request
			PutBucketPolicyOutput *s3.PutBucketPolicyOutput
		}
		Stub func(*s3.PutBucketPolicyInput) (*request.Request, *s3.PutBucketPolicyOutput)
	}
	PutBucketPolicyWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context              context.Context
			PutBucketPolicyInput *s3.PutBucketPolicyInput
			OptionSlice          []request.Option
		}
		Returns struct {
			PutBucketPolicyOutput *s3.PutBucketPolicyOutput
			Error                 error
		}
		Stub func(context.Context, *s3.PutBucketPolicyInput, ...request.Option) (*s3.PutBucketPolicyOutput, error)
	}
	PutBucketReplicationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketReplicationInput *s3.PutBucketReplicationInput
		}
		Returns struct {
			PutBucketReplicationOutput *s3.PutBucketReplicationOutput
			Error                      error
		}
		Stub func(*s3.PutBucketReplicationInput) (*s3.PutBucketReplicationOutput, error)
	}
	PutBucketReplicationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketReplicationInput *s3.PutBucketReplicationInput
		}
		Returns struct {
			Request                    *request.Request
			PutBucketReplicationOutput *s3.PutBucketReplicationOutput
		}
		Stub func(*s3.PutBucketReplicationInput) (*request.Request, *s3.PutBucketReplicationOutput)
	}
	PutBucketReplicationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                   context.Context
			PutBucketReplicationInput *s3.PutBucketReplicationInput
			OptionSlice               []request.Option
		}
		Returns struct {
			PutBucketReplicationOutput *s3.PutBucketReplicationOutput
			Error                      error
		}
		Stub func(context.Context, *s3.PutBucketReplicationInput, ...request.Option) (*s3.PutBucketReplicationOutput, error)
	}
	PutBucketRequestPaymentCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketRequestPaymentInput *s3.PutBucketRequestPaymentInput
		}
		Returns struct {
			PutBucketRequestPaymentOutput *s3.PutBucketRequestPaymentOutput
			Error                         error
		}
		Stub func(*s3.PutBucketRequestPaymentInput) (*s3.PutBucketRequestPaymentOutput, error)
	}
	PutBucketRequestPaymentRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketRequestPaymentInput *s3.PutBucketRequestPaymentInput
		}
		Returns struct {
			Request                       *request.Request
			PutBucketRequestPaymentOutput *s3.PutBucketRequestPaymentOutput
		}
		Stub func(*s3.PutBucketRequestPaymentInput) (*request.Request, *s3.PutBucketRequestPaymentOutput)
	}
	PutBucketRequestPaymentWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                      context.Context
			PutBucketRequestPaymentInput *s3.PutBucketRequestPaymentInput
			OptionSlice                  []request.Option
		}
		Returns struct {
			PutBucketRequestPaymentOutput *s3.PutBucketRequestPaymentOutput
			Error                         error
		}
		Stub func(context.Context, *s3.PutBucketRequestPaymentInput, ...request.Option) (*s3.PutBucketRequestPaymentOutput, error)
	}
	PutBucketTaggingCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketTaggingInput *s3.PutBucketTaggingInput
		}
		Returns struct {
			PutBucketTaggingOutput *s3.PutBucketTaggingOutput
			Error                  error
		}
		Stub func(*s3.PutBucketTaggingInput) (*s3.PutBucketTaggingOutput, error)
	}
	PutBucketTaggingRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketTaggingInput *s3.PutBucketTaggingInput
		}
		Returns struct {
			Request                *request.Request
			PutBucketTaggingOutput *s3.PutBucketTaggingOutput
		}
		Stub func(*s3.PutBucketTaggingInput) (*request.Request, *s3.PutBucketTaggingOutput)
	}
	PutBucketTaggingWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context               context.Context
			PutBucketTaggingInput *s3.PutBucketTaggingInput
			OptionSlice           []request.Option
		}
		Returns struct {
			PutBucketTaggingOutput *s3.PutBucketTaggingOutput
			Error                  error
		}
		Stub func(context.Context, *s3.PutBucketTaggingInput, ...request.Option) (*s3.PutBucketTaggingOutput, error)
	}
	PutBucketVersioningCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketVersioningInput *s3.PutBucketVersioningInput
		}
		Returns struct {
			PutBucketVersioningOutput *s3.PutBucketVersioningOutput
			Error                     error
		}
		Stub func(*s3.PutBucketVersioningInput) (*s3.PutBucketVersioningOutput, error)
	}
	PutBucketVersioningRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketVersioningInput *s3.PutBucketVersioningInput
		}
		Returns struct {
			Request                   *request.Request
			PutBucketVersioningOutput *s3.PutBucketVersioningOutput
		}
		Stub func(*s3.PutBucketVersioningInput) (*request.Request, *s3.PutBucketVersioningOutput)
	}
	PutBucketVersioningWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                  context.Context
			PutBucketVersioningInput *s3.PutBucketVersioningInput
			OptionSlice              []request.Option
		}
		Returns struct {
			PutBucketVersioningOutput *s3.PutBucketVersioningOutput
			Error                     error
		}
		Stub func(context.Context, *s3.PutBucketVersioningInput, ...request.Option) (*s3.PutBucketVersioningOutput, error)
	}
	PutBucketWebsiteCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketWebsiteInput *s3.PutBucketWebsiteInput
		}
		Returns struct {
			PutBucketWebsiteOutput *s3.PutBucketWebsiteOutput
			Error                  error
		}
		Stub func(*s3.PutBucketWebsiteInput) (*s3.PutBucketWebsiteOutput, error)
	}
	PutBucketWebsiteRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutBucketWebsiteInput *s3.PutBucketWebsiteInput
		}
		Returns struct {
			Request                *request.Request
			PutBucketWebsiteOutput *s3.PutBucketWebsiteOutput
		}
		Stub func(*s3.PutBucketWebsiteInput) (*request.Request, *s3.PutBucketWebsiteOutput)
	}
	PutBucketWebsiteWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context               context.Context
			PutBucketWebsiteInput *s3.PutBucketWebsiteInput
			OptionSlice           []request.Option
		}
		Returns struct {
			PutBucketWebsiteOutput *s3.PutBucketWebsiteOutput
			Error                  error
		}
		Stub func(context.Context, *s3.PutBucketWebsiteInput, ...request.Option) (*s3.PutBucketWebsiteOutput, error)
	}
	PutObjectCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutObjectInput *s3.PutObjectInput
		}
		Returns struct {
			PutObjectOutput *s3.PutObjectOutput
			Error           error
		}
		Stub func(*s3.PutObjectInput) (*s3.PutObjectOutput, error)
	}
	PutObjectAclCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutObjectAclInput *s3.PutObjectAclInput
		}
		Returns struct {
			PutObjectAclOutput *s3.PutObjectAclOutput
			Error              error
		}
		Stub func(*s3.PutObjectAclInput) (*s3.PutObjectAclOutput, error)
	}
	PutObjectAclRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutObjectAclInput *s3.PutObjectAclInput
		}
		Returns struct {
			Request            *request.Request
			PutObjectAclOutput *s3.PutObjectAclOutput
		}
		Stub func(*s3.PutObjectAclInput) (*request.Request, *s3.PutObjectAclOutput)
	}
	PutObjectAclWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context           context.Context
			PutObjectAclInput *s3.PutObjectAclInput
			OptionSlice       []request.Option
		}
		Returns struct {
			PutObjectAclOutput *s3.PutObjectAclOutput
			Error              error
		}
		Stub func(context.Context, *s3.PutObjectAclInput, ...request.Option) (*s3.PutObjectAclOutput, error)
	}
	PutObjectLegalHoldCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutObjectLegalHoldInput *s3.PutObjectLegalHoldInput
		}
		Returns struct {
			PutObjectLegalHoldOutput *s3.PutObjectLegalHoldOutput
			Error                    error
		}
		Stub func(*s3.PutObjectLegalHoldInput) (*s3.PutObjectLegalHoldOutput, error)
	}
	PutObjectLegalHoldRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutObjectLegalHoldInput *s3.PutObjectLegalHoldInput
		}
		Returns struct {
			Request                  *request.Request
			PutObjectLegalHoldOutput *s3.PutObjectLegalHoldOutput
		}
		Stub func(*s3.PutObjectLegalHoldInput) (*request.Request, *s3.PutObjectLegalHoldOutput)
	}
	PutObjectLegalHoldWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                 context.Context
			PutObjectLegalHoldInput *s3.PutObjectLegalHoldInput
			OptionSlice             []request.Option
		}
		Returns struct {
			PutObjectLegalHoldOutput *s3.PutObjectLegalHoldOutput
			Error                    error
		}
		Stub func(context.Context, *s3.PutObjectLegalHoldInput, ...request.Option) (*s3.PutObjectLegalHoldOutput, error)
	}
	PutObjectLockConfigurationCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutObjectLockConfigurationInput *s3.PutObjectLockConfigurationInput
		}
		Returns struct {
			PutObjectLockConfigurationOutput *s3.PutObjectLockConfigurationOutput
			Error                            error
		}
		Stub func(*s3.PutObjectLockConfigurationInput) (*s3.PutObjectLockConfigurationOutput, error)
	}
	PutObjectLockConfigurationRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutObjectLockConfigurationInput *s3.PutObjectLockConfigurationInput
		}
		Returns struct {
			Request                          *request.Request
			PutObjectLockConfigurationOutput *s3.PutObjectLockConfigurationOutput
		}
		Stub func(*s3.PutObjectLockConfigurationInput) (*request.Request, *s3.PutObjectLockConfigurationOutput)
	}
	PutObjectLockConfigurationWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                         context.Context
			PutObjectLockConfigurationInput *s3.PutObjectLockConfigurationInput
			OptionSlice                     []request.Option
		}
		Returns struct {
			PutObjectLockConfigurationOutput *s3.PutObjectLockConfigurationOutput
			Error                            error
		}
		Stub func(context.Context, *s3.PutObjectLockConfigurationInput, ...request.Option) (*s3.PutObjectLockConfigurationOutput, error)
	}
	PutObjectRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutObjectInput *s3.PutObjectInput
		}
		Returns struct {
			Request         *request.Request
			PutObjectOutput *s3.PutObjectOutput
		}
		Stub func(*s3.PutObjectInput) (*request.Request, *s3.PutObjectOutput)
	}
	PutObjectRetentionCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutObjectRetentionInput *s3.PutObjectRetentionInput
		}
		Returns struct {
			PutObjectRetentionOutput *s3.PutObjectRetentionOutput
			Error                    error
		}
		Stub func(*s3.PutObjectRetentionInput) (*s3.PutObjectRetentionOutput, error)
	}
	PutObjectRetentionRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutObjectRetentionInput *s3.PutObjectRetentionInput
		}
		Returns struct {
			Request                  *request.Request
			PutObjectRetentionOutput *s3.PutObjectRetentionOutput
		}
		Stub func(*s3.PutObjectRetentionInput) (*request.Request, *s3.PutObjectRetentionOutput)
	}
	PutObjectRetentionWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                 context.Context
			PutObjectRetentionInput *s3.PutObjectRetentionInput
			OptionSlice             []request.Option
		}
		Returns struct {
			PutObjectRetentionOutput *s3.PutObjectRetentionOutput
			Error                    error
		}
		Stub func(context.Context, *s3.PutObjectRetentionInput, ...request.Option) (*s3.PutObjectRetentionOutput, error)
	}
	PutObjectTaggingCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutObjectTaggingInput *s3.PutObjectTaggingInput
		}
		Returns struct {
			PutObjectTaggingOutput *s3.PutObjectTaggingOutput
			Error                  error
		}
		Stub func(*s3.PutObjectTaggingInput) (*s3.PutObjectTaggingOutput, error)
	}
	PutObjectTaggingRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutObjectTaggingInput *s3.PutObjectTaggingInput
		}
		Returns struct {
			Request                *request.Request
			PutObjectTaggingOutput *s3.PutObjectTaggingOutput
		}
		Stub func(*s3.PutObjectTaggingInput) (*request.Request, *s3.PutObjectTaggingOutput)
	}
	PutObjectTaggingWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context               context.Context
			PutObjectTaggingInput *s3.PutObjectTaggingInput
			OptionSlice           []request.Option
		}
		Returns struct {
			PutObjectTaggingOutput *s3.PutObjectTaggingOutput
			Error                  error
		}
		Stub func(context.Context, *s3.PutObjectTaggingInput, ...request.Option) (*s3.PutObjectTaggingOutput, error)
	}
	PutObjectWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context        context.Context
			PutObjectInput *s3.PutObjectInput
			OptionSlice    []request.Option
		}
		Returns struct {
			PutObjectOutput *s3.PutObjectOutput
			Error           error
		}
		Stub func(context.Context, *s3.PutObjectInput, ...request.Option) (*s3.PutObjectOutput, error)
	}
	PutPublicAccessBlockCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutPublicAccessBlockInput *s3.PutPublicAccessBlockInput
		}
		Returns struct {
			PutPublicAccessBlockOutput *s3.PutPublicAccessBlockOutput
			Error                      error
		}
		Stub func(*s3.PutPublicAccessBlockInput) (*s3.PutPublicAccessBlockOutput, error)
	}
	PutPublicAccessBlockRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			PutPublicAccessBlockInput *s3.PutPublicAccessBlockInput
		}
		Returns struct {
			Request                    *request.Request
			PutPublicAccessBlockOutput *s3.PutPublicAccessBlockOutput
		}
		Stub func(*s3.PutPublicAccessBlockInput) (*request.Request, *s3.PutPublicAccessBlockOutput)
	}
	PutPublicAccessBlockWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                   context.Context
			PutPublicAccessBlockInput *s3.PutPublicAccessBlockInput
			OptionSlice               []request.Option
		}
		Returns struct {
			PutPublicAccessBlockOutput *s3.PutPublicAccessBlockOutput
			Error                      error
		}
		Stub func(context.Context, *s3.PutPublicAccessBlockInput, ...request.Option) (*s3.PutPublicAccessBlockOutput, error)
	}
	RestoreObjectCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			RestoreObjectInput *s3.RestoreObjectInput
		}
		Returns struct {
			RestoreObjectOutput *s3.RestoreObjectOutput
			Error               error
		}
		Stub func(*s3.RestoreObjectInput) (*s3.RestoreObjectOutput, error)
	}
	RestoreObjectRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			RestoreObjectInput *s3.RestoreObjectInput
		}
		Returns struct {
			Request             *request.Request
			RestoreObjectOutput *s3.RestoreObjectOutput
		}
		Stub func(*s3.RestoreObjectInput) (*request.Request, *s3.RestoreObjectOutput)
	}
	RestoreObjectWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context            context.Context
			RestoreObjectInput *s3.RestoreObjectInput
			OptionSlice        []request.Option
		}
		Returns struct {
			RestoreObjectOutput *s3.RestoreObjectOutput
			Error               error
		}
		Stub func(context.Context, *s3.RestoreObjectInput, ...request.Option) (*s3.RestoreObjectOutput, error)
	}
	SelectObjectContentCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			SelectObjectContentInput *s3.SelectObjectContentInput
		}
		Returns struct {
			SelectObjectContentOutput *s3.SelectObjectContentOutput
			Error                     error
		}
		Stub func(*s3.SelectObjectContentInput) (*s3.SelectObjectContentOutput, error)
	}
	SelectObjectContentRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			SelectObjectContentInput *s3.SelectObjectContentInput
		}
		Returns struct {
			Request                   *request.Request
			SelectObjectContentOutput *s3.SelectObjectContentOutput
		}
		Stub func(*s3.SelectObjectContentInput) (*request.Request, *s3.SelectObjectContentOutput)
	}
	SelectObjectContentWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                  context.Context
			SelectObjectContentInput *s3.SelectObjectContentInput
			OptionSlice              []request.Option
		}
		Returns struct {
			SelectObjectContentOutput *s3.SelectObjectContentOutput
			Error                     error
		}
		Stub func(context.Context, *s3.SelectObjectContentInput, ...request.Option) (*s3.SelectObjectContentOutput, error)
	}
	UploadPartCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UploadPartInput *s3.UploadPartInput
		}
		Returns struct {
			UploadPartOutput *s3.UploadPartOutput
			Error            error
		}
		Stub func(*s3.UploadPartInput) (*s3.UploadPartOutput, error)
	}
	UploadPartCopyCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UploadPartCopyInput *s3.UploadPartCopyInput
		}
		Returns struct {
			UploadPartCopyOutput *s3.UploadPartCopyOutput
			Error                error
		}
		Stub func(*s3.UploadPartCopyInput) (*s3.UploadPartCopyOutput, error)
	}
	UploadPartCopyRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UploadPartCopyInput *s3.UploadPartCopyInput
		}
		Returns struct {
			Request              *request.Request
			UploadPartCopyOutput *s3.UploadPartCopyOutput
		}
		Stub func(*s3.UploadPartCopyInput) (*request.Request, *s3.UploadPartCopyOutput)
	}
	UploadPartCopyWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context             context.Context
			UploadPartCopyInput *s3.UploadPartCopyInput
			OptionSlice         []request.Option
		}
		Returns struct {
			UploadPartCopyOutput *s3.UploadPartCopyOutput
			Error                error
		}
		Stub func(context.Context, *s3.UploadPartCopyInput, ...request.Option) (*s3.UploadPartCopyOutput, error)
	}
	UploadPartRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			UploadPartInput *s3.UploadPartInput
		}
		Returns struct {
			Request          *request.Request
			UploadPartOutput *s3.UploadPartOutput
		}
		Stub func(*s3.UploadPartInput) (*request.Request, *s3.UploadPartOutput)
	}
	UploadPartWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context         context.Context
			UploadPartInput *s3.UploadPartInput
			OptionSlice     []request.Option
		}
		Returns struct {
			UploadPartOutput *s3.UploadPartOutput
			Error            error
		}
		Stub func(context.Context, *s3.UploadPartInput, ...request.Option) (*s3.UploadPartOutput, error)
	}
	WaitUntilBucketExistsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			HeadBucketInput *s3.HeadBucketInput
		}
		Returns struct {
			Error error
		}
		Stub func(*s3.HeadBucketInput) error
	}
	WaitUntilBucketExistsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context           context.Context
			HeadBucketInput   *s3.HeadBucketInput
			WaiterOptionSlice []request.WaiterOption
		}
		Returns struct {
			Error error
		}
		Stub func(context.Context, *s3.HeadBucketInput, ...request.WaiterOption) error
	}
	WaitUntilBucketNotExistsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			HeadBucketInput *s3.HeadBucketInput
		}
		Returns struct {
			Error error
		}
		Stub func(*s3.HeadBucketInput) error
	}
	WaitUntilBucketNotExistsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context           context.Context
			HeadBucketInput   *s3.HeadBucketInput
			WaiterOptionSlice []request.WaiterOption
		}
		Returns struct {
			Error error
		}
		Stub func(context.Context, *s3.HeadBucketInput, ...request.WaiterOption) error
	}
	WaitUntilObjectExistsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			HeadObjectInput *s3.HeadObjectInput
		}
		Returns struct {
			Error error
		}
		Stub func(*s3.HeadObjectInput) error
	}
	WaitUntilObjectExistsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context           context.Context
			HeadObjectInput   *s3.HeadObjectInput
			WaiterOptionSlice []request.WaiterOption
		}
		Returns struct {
			Error error
		}
		Stub func(context.Context, *s3.HeadObjectInput, ...request.WaiterOption) error
	}
	WaitUntilObjectNotExistsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			HeadObjectInput *s3.HeadObjectInput
		}
		Returns struct {
			Error error
		}
		Stub func(*s3.HeadObjectInput) error
	}
	WaitUntilObjectNotExistsWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context           context.Context
			HeadObjectInput   *s3.HeadObjectInput
			WaiterOptionSlice []request.WaiterOption
		}
		Returns struct {
			Error error
		}
		Stub func(context.Context, *s3.HeadObjectInput, ...request.WaiterOption) error
	}
	WriteGetObjectResponseCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			WriteGetObjectResponseInput *s3.WriteGetObjectResponseInput
		}
		Returns struct {
			WriteGetObjectResponseOutput *s3.WriteGetObjectResponseOutput
			Error                        error
		}
		Stub func(*s3.WriteGetObjectResponseInput) (*s3.WriteGetObjectResponseOutput, error)
	}
	WriteGetObjectResponseRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			WriteGetObjectResponseInput *s3.WriteGetObjectResponseInput
		}
		Returns struct {
			Request                      *request.Request
			WriteGetObjectResponseOutput *s3.WriteGetObjectResponseOutput
		}
		Stub func(*s3.WriteGetObjectResponseInput) (*request.Request, *s3.WriteGetObjectResponseOutput)
	}
	WriteGetObjectResponseWithContextCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Context                     context.Context
			WriteGetObjectResponseInput *s3.WriteGetObjectResponseInput
			OptionSlice                 []request.Option
		}
		Returns struct {
			WriteGetObjectResponseOutput *s3.WriteGetObjectResponseOutput
			Error                        error
		}
		Stub func(context.Context, *s3.WriteGetObjectResponseInput, ...request.Option) (*s3.WriteGetObjectResponseOutput, error)
	}
}

func (f *S3API) AbortMultipartUpload(param1 *s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error) {
	f.AbortMultipartUploadCall.mutex.Lock()
	defer f.AbortMultipartUploadCall.mutex.Unlock()
	f.AbortMultipartUploadCall.CallCount++
	f.AbortMultipartUploadCall.Receives.AbortMultipartUploadInput = param1
	if f.AbortMultipartUploadCall.Stub != nil {
		return f.AbortMultipartUploadCall.Stub(param1)
	}
	return f.AbortMultipartUploadCall.Returns.AbortMultipartUploadOutput, f.AbortMultipartUploadCall.Returns.Error
}
func (f *S3API) AbortMultipartUploadRequest(param1 *s3.AbortMultipartUploadInput) (*request.Request, *s3.AbortMultipartUploadOutput) {
	f.AbortMultipartUploadRequestCall.mutex.Lock()
	defer f.AbortMultipartUploadRequestCall.mutex.Unlock()
	f.AbortMultipartUploadRequestCall.CallCount++
	f.AbortMultipartUploadRequestCall.Receives.AbortMultipartUploadInput = param1
	if f.AbortMultipartUploadRequestCall.Stub != nil {
		return f.AbortMultipartUploadRequestCall.Stub(param1)
	}
	return f.AbortMultipartUploadRequestCall.Returns.Request, f.AbortMultipartUploadRequestCall.Returns.AbortMultipartUploadOutput
}
func (f *S3API) AbortMultipartUploadWithContext(param1 context.Context, param2 *s3.AbortMultipartUploadInput, param3 ...request.Option) (*s3.AbortMultipartUploadOutput, error) {
	f.AbortMultipartUploadWithContextCall.mutex.Lock()
	defer f.AbortMultipartUploadWithContextCall.mutex.Unlock()
	f.AbortMultipartUploadWithContextCall.CallCount++
	f.AbortMultipartUploadWithContextCall.Receives.Context = param1
	f.AbortMultipartUploadWithContextCall.Receives.AbortMultipartUploadInput = param2
	f.AbortMultipartUploadWithContextCall.Receives.OptionSlice = param3
	if f.AbortMultipartUploadWithContextCall.Stub != nil {
		return f.AbortMultipartUploadWithContextCall.Stub(param1, param2, param3...)
	}
	return f.AbortMultipartUploadWithContextCall.Returns.AbortMultipartUploadOutput, f.AbortMultipartUploadWithContextCall.Returns.Error
}
func (f *S3API) CompleteMultipartUpload(param1 *s3.CompleteMultipartUploadInput) (*s3.CompleteMultipartUploadOutput, error) {
	f.CompleteMultipartUploadCall.mutex.Lock()
	defer f.CompleteMultipartUploadCall.mutex.Unlock()
	f.CompleteMultipartUploadCall.CallCount++
	f.CompleteMultipartUploadCall.Receives.CompleteMultipartUploadInput = param1
	if f.CompleteMultipartUploadCall.Stub != nil {
		return f.CompleteMultipartUploadCall.Stub(param1)
	}
	return f.CompleteMultipartUploadCall.Returns.CompleteMultipartUploadOutput, f.CompleteMultipartUploadCall.Returns.Error
}
func (f *S3API) CompleteMultipartUploadRequest(param1 *s3.CompleteMultipartUploadInput) (*request.Request, *s3.CompleteMultipartUploadOutput) {
	f.CompleteMultipartUploadRequestCall.mutex.Lock()
	defer f.CompleteMultipartUploadRequestCall.mutex.Unlock()
	f.CompleteMultipartUploadRequestCall.CallCount++
	f.CompleteMultipartUploadRequestCall.Receives.CompleteMultipartUploadInput = param1
	if f.CompleteMultipartUploadRequestCall.Stub != nil {
		return f.CompleteMultipartUploadRequestCall.Stub(param1)
	}
	return f.CompleteMultipartUploadRequestCall.Returns.Request, f.CompleteMultipartUploadRequestCall.Returns.CompleteMultipartUploadOutput
}
func (f *S3API) CompleteMultipartUploadWithContext(param1 context.Context, param2 *s3.CompleteMultipartUploadInput, param3 ...request.Option) (*s3.CompleteMultipartUploadOutput, error) {
	f.CompleteMultipartUploadWithContextCall.mutex.Lock()
	defer f.CompleteMultipartUploadWithContextCall.mutex.Unlock()
	f.CompleteMultipartUploadWithContextCall.CallCount++
	f.CompleteMultipartUploadWithContextCall.Receives.Context = param1
	f.CompleteMultipartUploadWithContextCall.Receives.CompleteMultipartUploadInput = param2
	f.CompleteMultipartUploadWithContextCall.Receives.OptionSlice = param3
	if f.CompleteMultipartUploadWithContextCall.Stub != nil {
		return f.CompleteMultipartUploadWithContextCall.Stub(param1, param2, param3...)
	}
	return f.CompleteMultipartUploadWithContextCall.Returns.CompleteMultipartUploadOutput, f.CompleteMultipartUploadWithContextCall.Returns.Error
}
func (f *S3API) CopyObject(param1 *s3.CopyObjectInput) (*s3.CopyObjectOutput, error) {
	f.CopyObjectCall.mutex.Lock()
	defer f.CopyObjectCall.mutex.Unlock()
	f.CopyObjectCall.CallCount++
	f.CopyObjectCall.Receives.CopyObjectInput = param1
	if f.CopyObjectCall.Stub != nil {
		return f.CopyObjectCall.Stub(param1)
	}
	return f.CopyObjectCall.Returns.CopyObjectOutput, f.CopyObjectCall.Returns.Error
}
func (f *S3API) CopyObjectRequest(param1 *s3.CopyObjectInput) (*request.Request, *s3.CopyObjectOutput) {
	f.CopyObjectRequestCall.mutex.Lock()
	defer f.CopyObjectRequestCall.mutex.Unlock()
	f.CopyObjectRequestCall.CallCount++
	f.CopyObjectRequestCall.Receives.CopyObjectInput = param1
	if f.CopyObjectRequestCall.Stub != nil {
		return f.CopyObjectRequestCall.Stub(param1)
	}
	return f.CopyObjectRequestCall.Returns.Request, f.CopyObjectRequestCall.Returns.CopyObjectOutput
}
func (f *S3API) CopyObjectWithContext(param1 context.Context, param2 *s3.CopyObjectInput, param3 ...request.Option) (*s3.CopyObjectOutput, error) {
	f.CopyObjectWithContextCall.mutex.Lock()
	defer f.CopyObjectWithContextCall.mutex.Unlock()
	f.CopyObjectWithContextCall.CallCount++
	f.CopyObjectWithContextCall.Receives.Context = param1
	f.CopyObjectWithContextCall.Receives.CopyObjectInput = param2
	f.CopyObjectWithContextCall.Receives.OptionSlice = param3
	if f.CopyObjectWithContextCall.Stub != nil {
		return f.CopyObjectWithContextCall.Stub(param1, param2, param3...)
	}
	return f.CopyObjectWithContextCall.Returns.CopyObjectOutput, f.CopyObjectWithContextCall.Returns.Error
}
func (f *S3API) CreateBucket(param1 *s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
	f.CreateBucketCall.mutex.Lock()
	defer f.CreateBucketCall.mutex.Unlock()
	f.CreateBucketCall.CallCount++
	f.CreateBucketCall.Receives.CreateBucketInput = param1
	if f.CreateBucketCall.Stub != nil {
		return f.CreateBucketCall.Stub(param1)
	}
	return f.CreateBucketCall.Returns.CreateBucketOutput, f.CreateBucketCall.Returns.Error
}
func (f *S3API) CreateBucketRequest(param1 *s3.CreateBucketInput) (*request.Request, *s3.CreateBucketOutput) {
	f.CreateBucketRequestCall.mutex.Lock()
	defer f.CreateBucketRequestCall.mutex.Unlock()
	f.CreateBucketRequestCall.CallCount++
	f.CreateBucketRequestCall.Receives.CreateBucketInput = param1
	if f.CreateBucketRequestCall.Stub != nil {
		return f.CreateBucketRequestCall.Stub(param1)
	}
	return f.CreateBucketRequestCall.Returns.Request, f.CreateBucketRequestCall.Returns.CreateBucketOutput
}
func (f *S3API) CreateBucketWithContext(param1 context.Context, param2 *s3.CreateBucketInput, param3 ...request.Option) (*s3.CreateBucketOutput, error) {
	f.CreateBucketWithContextCall.mutex.Lock()
	defer f.CreateBucketWithContextCall.mutex.Unlock()
	f.CreateBucketWithContextCall.CallCount++
	f.CreateBucketWithContextCall.Receives.Context = param1
	f.CreateBucketWithContextCall.Receives.CreateBucketInput = param2
	f.CreateBucketWithContextCall.Receives.OptionSlice = param3
	if f.CreateBucketWithContextCall.Stub != nil {
		return f.CreateBucketWithContextCall.Stub(param1, param2, param3...)
	}
	return f.CreateBucketWithContextCall.Returns.CreateBucketOutput, f.CreateBucketWithContextCall.Returns.Error
}
func (f *S3API) CreateMultipartUpload(param1 *s3.CreateMultipartUploadInput) (*s3.CreateMultipartUploadOutput, error) {
	f.CreateMultipartUploadCall.mutex.Lock()
	defer f.CreateMultipartUploadCall.mutex.Unlock()
	f.CreateMultipartUploadCall.CallCount++
	f.CreateMultipartUploadCall.Receives.CreateMultipartUploadInput = param1
	if f.CreateMultipartUploadCall.Stub != nil {
		return f.CreateMultipartUploadCall.Stub(param1)
	}
	return f.CreateMultipartUploadCall.Returns.CreateMultipartUploadOutput, f.CreateMultipartUploadCall.Returns.Error
}
func (f *S3API) CreateMultipartUploadRequest(param1 *s3.CreateMultipartUploadInput) (*request.Request, *s3.CreateMultipartUploadOutput) {
	f.CreateMultipartUploadRequestCall.mutex.Lock()
	defer f.CreateMultipartUploadRequestCall.mutex.Unlock()
	f.CreateMultipartUploadRequestCall.CallCount++
	f.CreateMultipartUploadRequestCall.Receives.CreateMultipartUploadInput = param1
	if f.CreateMultipartUploadRequestCall.Stub != nil {
		return f.CreateMultipartUploadRequestCall.Stub(param1)
	}
	return f.CreateMultipartUploadRequestCall.Returns.Request, f.CreateMultipartUploadRequestCall.Returns.CreateMultipartUploadOutput
}
func (f *S3API) CreateMultipartUploadWithContext(param1 context.Context, param2 *s3.CreateMultipartUploadInput, param3 ...request.Option) (*s3.CreateMultipartUploadOutput, error) {
	f.CreateMultipartUploadWithContextCall.mutex.Lock()
	defer f.CreateMultipartUploadWithContextCall.mutex.Unlock()
	f.CreateMultipartUploadWithContextCall.CallCount++
	f.CreateMultipartUploadWithContextCall.Receives.Context = param1
	f.CreateMultipartUploadWithContextCall.Receives.CreateMultipartUploadInput = param2
	f.CreateMultipartUploadWithContextCall.Receives.OptionSlice = param3
	if f.CreateMultipartUploadWithContextCall.Stub != nil {
		return f.CreateMultipartUploadWithContextCall.Stub(param1, param2, param3...)
	}
	return f.CreateMultipartUploadWithContextCall.Returns.CreateMultipartUploadOutput, f.CreateMultipartUploadWithContextCall.Returns.Error
}
func (f *S3API) DeleteBucket(param1 *s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error) {
	f.DeleteBucketCall.mutex.Lock()
	defer f.DeleteBucketCall.mutex.Unlock()
	f.DeleteBucketCall.CallCount++
	f.DeleteBucketCall.Receives.DeleteBucketInput = param1
	if f.DeleteBucketCall.Stub != nil {
		return f.DeleteBucketCall.Stub(param1)
	}
	return f.DeleteBucketCall.Returns.DeleteBucketOutput, f.DeleteBucketCall.Returns.Error
}
func (f *S3API) DeleteBucketAnalyticsConfiguration(param1 *s3.DeleteBucketAnalyticsConfigurationInput) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
	f.DeleteBucketAnalyticsConfigurationCall.mutex.Lock()
	defer f.DeleteBucketAnalyticsConfigurationCall.mutex.Unlock()
	f.DeleteBucketAnalyticsConfigurationCall.CallCount++
	f.DeleteBucketAnalyticsConfigurationCall.Receives.DeleteBucketAnalyticsConfigurationInput = param1
	if f.DeleteBucketAnalyticsConfigurationCall.Stub != nil {
		return f.DeleteBucketAnalyticsConfigurationCall.Stub(param1)
	}
	return f.DeleteBucketAnalyticsConfigurationCall.Returns.DeleteBucketAnalyticsConfigurationOutput, f.DeleteBucketAnalyticsConfigurationCall.Returns.Error
}
func (f *S3API) DeleteBucketAnalyticsConfigurationRequest(param1 *s3.DeleteBucketAnalyticsConfigurationInput) (*request.Request, *s3.DeleteBucketAnalyticsConfigurationOutput) {
	f.DeleteBucketAnalyticsConfigurationRequestCall.mutex.Lock()
	defer f.DeleteBucketAnalyticsConfigurationRequestCall.mutex.Unlock()
	f.DeleteBucketAnalyticsConfigurationRequestCall.CallCount++
	f.DeleteBucketAnalyticsConfigurationRequestCall.Receives.DeleteBucketAnalyticsConfigurationInput = param1
	if f.DeleteBucketAnalyticsConfigurationRequestCall.Stub != nil {
		return f.DeleteBucketAnalyticsConfigurationRequestCall.Stub(param1)
	}
	return f.DeleteBucketAnalyticsConfigurationRequestCall.Returns.Request, f.DeleteBucketAnalyticsConfigurationRequestCall.Returns.DeleteBucketAnalyticsConfigurationOutput
}
func (f *S3API) DeleteBucketAnalyticsConfigurationWithContext(param1 context.Context, param2 *s3.DeleteBucketAnalyticsConfigurationInput, param3 ...request.Option) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
	f.DeleteBucketAnalyticsConfigurationWithContextCall.mutex.Lock()
	defer f.DeleteBucketAnalyticsConfigurationWithContextCall.mutex.Unlock()
	f.DeleteBucketAnalyticsConfigurationWithContextCall.CallCount++
	f.DeleteBucketAnalyticsConfigurationWithContextCall.Receives.Context = param1
	f.DeleteBucketAnalyticsConfigurationWithContextCall.Receives.DeleteBucketAnalyticsConfigurationInput = param2
	f.DeleteBucketAnalyticsConfigurationWithContextCall.Receives.OptionSlice = param3
	if f.DeleteBucketAnalyticsConfigurationWithContextCall.Stub != nil {
		return f.DeleteBucketAnalyticsConfigurationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DeleteBucketAnalyticsConfigurationWithContextCall.Returns.DeleteBucketAnalyticsConfigurationOutput, f.DeleteBucketAnalyticsConfigurationWithContextCall.Returns.Error
}
func (f *S3API) DeleteBucketCors(param1 *s3.DeleteBucketCorsInput) (*s3.DeleteBucketCorsOutput, error) {
	f.DeleteBucketCorsCall.mutex.Lock()
	defer f.DeleteBucketCorsCall.mutex.Unlock()
	f.DeleteBucketCorsCall.CallCount++
	f.DeleteBucketCorsCall.Receives.DeleteBucketCorsInput = param1
	if f.DeleteBucketCorsCall.Stub != nil {
		return f.DeleteBucketCorsCall.Stub(param1)
	}
	return f.DeleteBucketCorsCall.Returns.DeleteBucketCorsOutput, f.DeleteBucketCorsCall.Returns.Error
}
func (f *S3API) DeleteBucketCorsRequest(param1 *s3.DeleteBucketCorsInput) (*request.Request, *s3.DeleteBucketCorsOutput) {
	f.DeleteBucketCorsRequestCall.mutex.Lock()
	defer f.DeleteBucketCorsRequestCall.mutex.Unlock()
	f.DeleteBucketCorsRequestCall.CallCount++
	f.DeleteBucketCorsRequestCall.Receives.DeleteBucketCorsInput = param1
	if f.DeleteBucketCorsRequestCall.Stub != nil {
		return f.DeleteBucketCorsRequestCall.Stub(param1)
	}
	return f.DeleteBucketCorsRequestCall.Returns.Request, f.DeleteBucketCorsRequestCall.Returns.DeleteBucketCorsOutput
}
func (f *S3API) DeleteBucketCorsWithContext(param1 context.Context, param2 *s3.DeleteBucketCorsInput, param3 ...request.Option) (*s3.DeleteBucketCorsOutput, error) {
	f.DeleteBucketCorsWithContextCall.mutex.Lock()
	defer f.DeleteBucketCorsWithContextCall.mutex.Unlock()
	f.DeleteBucketCorsWithContextCall.CallCount++
	f.DeleteBucketCorsWithContextCall.Receives.Context = param1
	f.DeleteBucketCorsWithContextCall.Receives.DeleteBucketCorsInput = param2
	f.DeleteBucketCorsWithContextCall.Receives.OptionSlice = param3
	if f.DeleteBucketCorsWithContextCall.Stub != nil {
		return f.DeleteBucketCorsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DeleteBucketCorsWithContextCall.Returns.DeleteBucketCorsOutput, f.DeleteBucketCorsWithContextCall.Returns.Error
}
func (f *S3API) DeleteBucketEncryption(param1 *s3.DeleteBucketEncryptionInput) (*s3.DeleteBucketEncryptionOutput, error) {
	f.DeleteBucketEncryptionCall.mutex.Lock()
	defer f.DeleteBucketEncryptionCall.mutex.Unlock()
	f.DeleteBucketEncryptionCall.CallCount++
	f.DeleteBucketEncryptionCall.Receives.DeleteBucketEncryptionInput = param1
	if f.DeleteBucketEncryptionCall.Stub != nil {
		return f.DeleteBucketEncryptionCall.Stub(param1)
	}
	return f.DeleteBucketEncryptionCall.Returns.DeleteBucketEncryptionOutput, f.DeleteBucketEncryptionCall.Returns.Error
}
func (f *S3API) DeleteBucketEncryptionRequest(param1 *s3.DeleteBucketEncryptionInput) (*request.Request, *s3.DeleteBucketEncryptionOutput) {
	f.DeleteBucketEncryptionRequestCall.mutex.Lock()
	defer f.DeleteBucketEncryptionRequestCall.mutex.Unlock()
	f.DeleteBucketEncryptionRequestCall.CallCount++
	f.DeleteBucketEncryptionRequestCall.Receives.DeleteBucketEncryptionInput = param1
	if f.DeleteBucketEncryptionRequestCall.Stub != nil {
		return f.DeleteBucketEncryptionRequestCall.Stub(param1)
	}
	return f.DeleteBucketEncryptionRequestCall.Returns.Request, f.DeleteBucketEncryptionRequestCall.Returns.DeleteBucketEncryptionOutput
}
func (f *S3API) DeleteBucketEncryptionWithContext(param1 context.Context, param2 *s3.DeleteBucketEncryptionInput, param3 ...request.Option) (*s3.DeleteBucketEncryptionOutput, error) {
	f.DeleteBucketEncryptionWithContextCall.mutex.Lock()
	defer f.DeleteBucketEncryptionWithContextCall.mutex.Unlock()
	f.DeleteBucketEncryptionWithContextCall.CallCount++
	f.DeleteBucketEncryptionWithContextCall.Receives.Context = param1
	f.DeleteBucketEncryptionWithContextCall.Receives.DeleteBucketEncryptionInput = param2
	f.DeleteBucketEncryptionWithContextCall.Receives.OptionSlice = param3
	if f.DeleteBucketEncryptionWithContextCall.Stub != nil {
		return f.DeleteBucketEncryptionWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DeleteBucketEncryptionWithContextCall.Returns.DeleteBucketEncryptionOutput, f.DeleteBucketEncryptionWithContextCall.Returns.Error
}
func (f *S3API) DeleteBucketIntelligentTieringConfiguration(param1 *s3.DeleteBucketIntelligentTieringConfigurationInput) (*s3.DeleteBucketIntelligentTieringConfigurationOutput, error) {
	f.DeleteBucketIntelligentTieringConfigurationCall.mutex.Lock()
	defer f.DeleteBucketIntelligentTieringConfigurationCall.mutex.Unlock()
	f.DeleteBucketIntelligentTieringConfigurationCall.CallCount++
	f.DeleteBucketIntelligentTieringConfigurationCall.Receives.DeleteBucketIntelligentTieringConfigurationInput = param1
	if f.DeleteBucketIntelligentTieringConfigurationCall.Stub != nil {
		return f.DeleteBucketIntelligentTieringConfigurationCall.Stub(param1)
	}
	return f.DeleteBucketIntelligentTieringConfigurationCall.Returns.DeleteBucketIntelligentTieringConfigurationOutput, f.DeleteBucketIntelligentTieringConfigurationCall.Returns.Error
}
func (f *S3API) DeleteBucketIntelligentTieringConfigurationRequest(param1 *s3.DeleteBucketIntelligentTieringConfigurationInput) (*request.Request, *s3.DeleteBucketIntelligentTieringConfigurationOutput) {
	f.DeleteBucketIntelligentTieringConfigurationRequestCall.mutex.Lock()
	defer f.DeleteBucketIntelligentTieringConfigurationRequestCall.mutex.Unlock()
	f.DeleteBucketIntelligentTieringConfigurationRequestCall.CallCount++
	f.DeleteBucketIntelligentTieringConfigurationRequestCall.Receives.DeleteBucketIntelligentTieringConfigurationInput = param1
	if f.DeleteBucketIntelligentTieringConfigurationRequestCall.Stub != nil {
		return f.DeleteBucketIntelligentTieringConfigurationRequestCall.Stub(param1)
	}
	return f.DeleteBucketIntelligentTieringConfigurationRequestCall.Returns.Request, f.DeleteBucketIntelligentTieringConfigurationRequestCall.Returns.DeleteBucketIntelligentTieringConfigurationOutput
}
func (f *S3API) DeleteBucketIntelligentTieringConfigurationWithContext(param1 context.Context, param2 *s3.DeleteBucketIntelligentTieringConfigurationInput, param3 ...request.Option) (*s3.DeleteBucketIntelligentTieringConfigurationOutput, error) {
	f.DeleteBucketIntelligentTieringConfigurationWithContextCall.mutex.Lock()
	defer f.DeleteBucketIntelligentTieringConfigurationWithContextCall.mutex.Unlock()
	f.DeleteBucketIntelligentTieringConfigurationWithContextCall.CallCount++
	f.DeleteBucketIntelligentTieringConfigurationWithContextCall.Receives.Context = param1
	f.DeleteBucketIntelligentTieringConfigurationWithContextCall.Receives.DeleteBucketIntelligentTieringConfigurationInput = param2
	f.DeleteBucketIntelligentTieringConfigurationWithContextCall.Receives.OptionSlice = param3
	if f.DeleteBucketIntelligentTieringConfigurationWithContextCall.Stub != nil {
		return f.DeleteBucketIntelligentTieringConfigurationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DeleteBucketIntelligentTieringConfigurationWithContextCall.Returns.DeleteBucketIntelligentTieringConfigurationOutput, f.DeleteBucketIntelligentTieringConfigurationWithContextCall.Returns.Error
}
func (f *S3API) DeleteBucketInventoryConfiguration(param1 *s3.DeleteBucketInventoryConfigurationInput) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
	f.DeleteBucketInventoryConfigurationCall.mutex.Lock()
	defer f.DeleteBucketInventoryConfigurationCall.mutex.Unlock()
	f.DeleteBucketInventoryConfigurationCall.CallCount++
	f.DeleteBucketInventoryConfigurationCall.Receives.DeleteBucketInventoryConfigurationInput = param1
	if f.DeleteBucketInventoryConfigurationCall.Stub != nil {
		return f.DeleteBucketInventoryConfigurationCall.Stub(param1)
	}
	return f.DeleteBucketInventoryConfigurationCall.Returns.DeleteBucketInventoryConfigurationOutput, f.DeleteBucketInventoryConfigurationCall.Returns.Error
}
func (f *S3API) DeleteBucketInventoryConfigurationRequest(param1 *s3.DeleteBucketInventoryConfigurationInput) (*request.Request, *s3.DeleteBucketInventoryConfigurationOutput) {
	f.DeleteBucketInventoryConfigurationRequestCall.mutex.Lock()
	defer f.DeleteBucketInventoryConfigurationRequestCall.mutex.Unlock()
	f.DeleteBucketInventoryConfigurationRequestCall.CallCount++
	f.DeleteBucketInventoryConfigurationRequestCall.Receives.DeleteBucketInventoryConfigurationInput = param1
	if f.DeleteBucketInventoryConfigurationRequestCall.Stub != nil {
		return f.DeleteBucketInventoryConfigurationRequestCall.Stub(param1)
	}
	return f.DeleteBucketInventoryConfigurationRequestCall.Returns.Request, f.DeleteBucketInventoryConfigurationRequestCall.Returns.DeleteBucketInventoryConfigurationOutput
}
func (f *S3API) DeleteBucketInventoryConfigurationWithContext(param1 context.Context, param2 *s3.DeleteBucketInventoryConfigurationInput, param3 ...request.Option) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
	f.DeleteBucketInventoryConfigurationWithContextCall.mutex.Lock()
	defer f.DeleteBucketInventoryConfigurationWithContextCall.mutex.Unlock()
	f.DeleteBucketInventoryConfigurationWithContextCall.CallCount++
	f.DeleteBucketInventoryConfigurationWithContextCall.Receives.Context = param1
	f.DeleteBucketInventoryConfigurationWithContextCall.Receives.DeleteBucketInventoryConfigurationInput = param2
	f.DeleteBucketInventoryConfigurationWithContextCall.Receives.OptionSlice = param3
	if f.DeleteBucketInventoryConfigurationWithContextCall.Stub != nil {
		return f.DeleteBucketInventoryConfigurationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DeleteBucketInventoryConfigurationWithContextCall.Returns.DeleteBucketInventoryConfigurationOutput, f.DeleteBucketInventoryConfigurationWithContextCall.Returns.Error
}
func (f *S3API) DeleteBucketLifecycle(param1 *s3.DeleteBucketLifecycleInput) (*s3.DeleteBucketLifecycleOutput, error) {
	f.DeleteBucketLifecycleCall.mutex.Lock()
	defer f.DeleteBucketLifecycleCall.mutex.Unlock()
	f.DeleteBucketLifecycleCall.CallCount++
	f.DeleteBucketLifecycleCall.Receives.DeleteBucketLifecycleInput = param1
	if f.DeleteBucketLifecycleCall.Stub != nil {
		return f.DeleteBucketLifecycleCall.Stub(param1)
	}
	return f.DeleteBucketLifecycleCall.Returns.DeleteBucketLifecycleOutput, f.DeleteBucketLifecycleCall.Returns.Error
}
func (f *S3API) DeleteBucketLifecycleRequest(param1 *s3.DeleteBucketLifecycleInput) (*request.Request, *s3.DeleteBucketLifecycleOutput) {
	f.DeleteBucketLifecycleRequestCall.mutex.Lock()
	defer f.DeleteBucketLifecycleRequestCall.mutex.Unlock()
	f.DeleteBucketLifecycleRequestCall.CallCount++
	f.DeleteBucketLifecycleRequestCall.Receives.DeleteBucketLifecycleInput = param1
	if f.DeleteBucketLifecycleRequestCall.Stub != nil {
		return f.DeleteBucketLifecycleRequestCall.Stub(param1)
	}
	return f.DeleteBucketLifecycleRequestCall.Returns.Request, f.DeleteBucketLifecycleRequestCall.Returns.DeleteBucketLifecycleOutput
}
func (f *S3API) DeleteBucketLifecycleWithContext(param1 context.Context, param2 *s3.DeleteBucketLifecycleInput, param3 ...request.Option) (*s3.DeleteBucketLifecycleOutput, error) {
	f.DeleteBucketLifecycleWithContextCall.mutex.Lock()
	defer f.DeleteBucketLifecycleWithContextCall.mutex.Unlock()
	f.DeleteBucketLifecycleWithContextCall.CallCount++
	f.DeleteBucketLifecycleWithContextCall.Receives.Context = param1
	f.DeleteBucketLifecycleWithContextCall.Receives.DeleteBucketLifecycleInput = param2
	f.DeleteBucketLifecycleWithContextCall.Receives.OptionSlice = param3
	if f.DeleteBucketLifecycleWithContextCall.Stub != nil {
		return f.DeleteBucketLifecycleWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DeleteBucketLifecycleWithContextCall.Returns.DeleteBucketLifecycleOutput, f.DeleteBucketLifecycleWithContextCall.Returns.Error
}
func (f *S3API) DeleteBucketMetricsConfiguration(param1 *s3.DeleteBucketMetricsConfigurationInput) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
	f.DeleteBucketMetricsConfigurationCall.mutex.Lock()
	defer f.DeleteBucketMetricsConfigurationCall.mutex.Unlock()
	f.DeleteBucketMetricsConfigurationCall.CallCount++
	f.DeleteBucketMetricsConfigurationCall.Receives.DeleteBucketMetricsConfigurationInput = param1
	if f.DeleteBucketMetricsConfigurationCall.Stub != nil {
		return f.DeleteBucketMetricsConfigurationCall.Stub(param1)
	}
	return f.DeleteBucketMetricsConfigurationCall.Returns.DeleteBucketMetricsConfigurationOutput, f.DeleteBucketMetricsConfigurationCall.Returns.Error
}
func (f *S3API) DeleteBucketMetricsConfigurationRequest(param1 *s3.DeleteBucketMetricsConfigurationInput) (*request.Request, *s3.DeleteBucketMetricsConfigurationOutput) {
	f.DeleteBucketMetricsConfigurationRequestCall.mutex.Lock()
	defer f.DeleteBucketMetricsConfigurationRequestCall.mutex.Unlock()
	f.DeleteBucketMetricsConfigurationRequestCall.CallCount++
	f.DeleteBucketMetricsConfigurationRequestCall.Receives.DeleteBucketMetricsConfigurationInput = param1
	if f.DeleteBucketMetricsConfigurationRequestCall.Stub != nil {
		return f.DeleteBucketMetricsConfigurationRequestCall.Stub(param1)
	}
	return f.DeleteBucketMetricsConfigurationRequestCall.Returns.Request, f.DeleteBucketMetricsConfigurationRequestCall.Returns.DeleteBucketMetricsConfigurationOutput
}
func (f *S3API) DeleteBucketMetricsConfigurationWithContext(param1 context.Context, param2 *s3.DeleteBucketMetricsConfigurationInput, param3 ...request.Option) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
	f.DeleteBucketMetricsConfigurationWithContextCall.mutex.Lock()
	defer f.DeleteBucketMetricsConfigurationWithContextCall.mutex.Unlock()
	f.DeleteBucketMetricsConfigurationWithContextCall.CallCount++
	f.DeleteBucketMetricsConfigurationWithContextCall.Receives.Context = param1
	f.DeleteBucketMetricsConfigurationWithContextCall.Receives.DeleteBucketMetricsConfigurationInput = param2
	f.DeleteBucketMetricsConfigurationWithContextCall.Receives.OptionSlice = param3
	if f.DeleteBucketMetricsConfigurationWithContextCall.Stub != nil {
		return f.DeleteBucketMetricsConfigurationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DeleteBucketMetricsConfigurationWithContextCall.Returns.DeleteBucketMetricsConfigurationOutput, f.DeleteBucketMetricsConfigurationWithContextCall.Returns.Error
}
func (f *S3API) DeleteBucketOwnershipControls(param1 *s3.DeleteBucketOwnershipControlsInput) (*s3.DeleteBucketOwnershipControlsOutput, error) {
	f.DeleteBucketOwnershipControlsCall.mutex.Lock()
	defer f.DeleteBucketOwnershipControlsCall.mutex.Unlock()
	f.DeleteBucketOwnershipControlsCall.CallCount++
	f.DeleteBucketOwnershipControlsCall.Receives.DeleteBucketOwnershipControlsInput = param1
	if f.DeleteBucketOwnershipControlsCall.Stub != nil {
		return f.DeleteBucketOwnershipControlsCall.Stub(param1)
	}
	return f.DeleteBucketOwnershipControlsCall.Returns.DeleteBucketOwnershipControlsOutput, f.DeleteBucketOwnershipControlsCall.Returns.Error
}
func (f *S3API) DeleteBucketOwnershipControlsRequest(param1 *s3.DeleteBucketOwnershipControlsInput) (*request.Request, *s3.DeleteBucketOwnershipControlsOutput) {
	f.DeleteBucketOwnershipControlsRequestCall.mutex.Lock()
	defer f.DeleteBucketOwnershipControlsRequestCall.mutex.Unlock()
	f.DeleteBucketOwnershipControlsRequestCall.CallCount++
	f.DeleteBucketOwnershipControlsRequestCall.Receives.DeleteBucketOwnershipControlsInput = param1
	if f.DeleteBucketOwnershipControlsRequestCall.Stub != nil {
		return f.DeleteBucketOwnershipControlsRequestCall.Stub(param1)
	}
	return f.DeleteBucketOwnershipControlsRequestCall.Returns.Request, f.DeleteBucketOwnershipControlsRequestCall.Returns.DeleteBucketOwnershipControlsOutput
}
func (f *S3API) DeleteBucketOwnershipControlsWithContext(param1 context.Context, param2 *s3.DeleteBucketOwnershipControlsInput, param3 ...request.Option) (*s3.DeleteBucketOwnershipControlsOutput, error) {
	f.DeleteBucketOwnershipControlsWithContextCall.mutex.Lock()
	defer f.DeleteBucketOwnershipControlsWithContextCall.mutex.Unlock()
	f.DeleteBucketOwnershipControlsWithContextCall.CallCount++
	f.DeleteBucketOwnershipControlsWithContextCall.Receives.Context = param1
	f.DeleteBucketOwnershipControlsWithContextCall.Receives.DeleteBucketOwnershipControlsInput = param2
	f.DeleteBucketOwnershipControlsWithContextCall.Receives.OptionSlice = param3
	if f.DeleteBucketOwnershipControlsWithContextCall.Stub != nil {
		return f.DeleteBucketOwnershipControlsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DeleteBucketOwnershipControlsWithContextCall.Returns.DeleteBucketOwnershipControlsOutput, f.DeleteBucketOwnershipControlsWithContextCall.Returns.Error
}
func (f *S3API) DeleteBucketPolicy(param1 *s3.DeleteBucketPolicyInput) (*s3.DeleteBucketPolicyOutput, error) {
	f.DeleteBucketPolicyCall.mutex.Lock()
	defer f.DeleteBucketPolicyCall.mutex.Unlock()
	f.DeleteBucketPolicyCall.CallCount++
	f.DeleteBucketPolicyCall.Receives.DeleteBucketPolicyInput = param1
	if f.DeleteBucketPolicyCall.Stub != nil {
		return f.DeleteBucketPolicyCall.Stub(param1)
	}
	return f.DeleteBucketPolicyCall.Returns.DeleteBucketPolicyOutput, f.DeleteBucketPolicyCall.Returns.Error
}
func (f *S3API) DeleteBucketPolicyRequest(param1 *s3.DeleteBucketPolicyInput) (*request.Request, *s3.DeleteBucketPolicyOutput) {
	f.DeleteBucketPolicyRequestCall.mutex.Lock()
	defer f.DeleteBucketPolicyRequestCall.mutex.Unlock()
	f.DeleteBucketPolicyRequestCall.CallCount++
	f.DeleteBucketPolicyRequestCall.Receives.DeleteBucketPolicyInput = param1
	if f.DeleteBucketPolicyRequestCall.Stub != nil {
		return f.DeleteBucketPolicyRequestCall.Stub(param1)
	}
	return f.DeleteBucketPolicyRequestCall.Returns.Request, f.DeleteBucketPolicyRequestCall.Returns.DeleteBucketPolicyOutput
}
func (f *S3API) DeleteBucketPolicyWithContext(param1 context.Context, param2 *s3.DeleteBucketPolicyInput, param3 ...request.Option) (*s3.DeleteBucketPolicyOutput, error) {
	f.DeleteBucketPolicyWithContextCall.mutex.Lock()
	defer f.DeleteBucketPolicyWithContextCall.mutex.Unlock()
	f.DeleteBucketPolicyWithContextCall.CallCount++
	f.DeleteBucketPolicyWithContextCall.Receives.Context = param1
	f.DeleteBucketPolicyWithContextCall.Receives.DeleteBucketPolicyInput = param2
	f.DeleteBucketPolicyWithContextCall.Receives.OptionSlice = param3
	if f.DeleteBucketPolicyWithContextCall.Stub != nil {
		return f.DeleteBucketPolicyWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DeleteBucketPolicyWithContextCall.Returns.DeleteBucketPolicyOutput, f.DeleteBucketPolicyWithContextCall.Returns.Error
}
func (f *S3API) DeleteBucketReplication(param1 *s3.DeleteBucketReplicationInput) (*s3.DeleteBucketReplicationOutput, error) {
	f.DeleteBucketReplicationCall.mutex.Lock()
	defer f.DeleteBucketReplicationCall.mutex.Unlock()
	f.DeleteBucketReplicationCall.CallCount++
	f.DeleteBucketReplicationCall.Receives.DeleteBucketReplicationInput = param1
	if f.DeleteBucketReplicationCall.Stub != nil {
		return f.DeleteBucketReplicationCall.Stub(param1)
	}
	return f.DeleteBucketReplicationCall.Returns.DeleteBucketReplicationOutput, f.DeleteBucketReplicationCall.Returns.Error
}
func (f *S3API) DeleteBucketReplicationRequest(param1 *s3.DeleteBucketReplicationInput) (*request.Request, *s3.DeleteBucketReplicationOutput) {
	f.DeleteBucketReplicationRequestCall.mutex.Lock()
	defer f.DeleteBucketReplicationRequestCall.mutex.Unlock()
	f.DeleteBucketReplicationRequestCall.CallCount++
	f.DeleteBucketReplicationRequestCall.Receives.DeleteBucketReplicationInput = param1
	if f.DeleteBucketReplicationRequestCall.Stub != nil {
		return f.DeleteBucketReplicationRequestCall.Stub(param1)
	}
	return f.DeleteBucketReplicationRequestCall.Returns.Request, f.DeleteBucketReplicationRequestCall.Returns.DeleteBucketReplicationOutput
}
func (f *S3API) DeleteBucketReplicationWithContext(param1 context.Context, param2 *s3.DeleteBucketReplicationInput, param3 ...request.Option) (*s3.DeleteBucketReplicationOutput, error) {
	f.DeleteBucketReplicationWithContextCall.mutex.Lock()
	defer f.DeleteBucketReplicationWithContextCall.mutex.Unlock()
	f.DeleteBucketReplicationWithContextCall.CallCount++
	f.DeleteBucketReplicationWithContextCall.Receives.Context = param1
	f.DeleteBucketReplicationWithContextCall.Receives.DeleteBucketReplicationInput = param2
	f.DeleteBucketReplicationWithContextCall.Receives.OptionSlice = param3
	if f.DeleteBucketReplicationWithContextCall.Stub != nil {
		return f.DeleteBucketReplicationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DeleteBucketReplicationWithContextCall.Returns.DeleteBucketReplicationOutput, f.DeleteBucketReplicationWithContextCall.Returns.Error
}
func (f *S3API) DeleteBucketRequest(param1 *s3.DeleteBucketInput) (*request.Request, *s3.DeleteBucketOutput) {
	f.DeleteBucketRequestCall.mutex.Lock()
	defer f.DeleteBucketRequestCall.mutex.Unlock()
	f.DeleteBucketRequestCall.CallCount++
	f.DeleteBucketRequestCall.Receives.DeleteBucketInput = param1
	if f.DeleteBucketRequestCall.Stub != nil {
		return f.DeleteBucketRequestCall.Stub(param1)
	}
	return f.DeleteBucketRequestCall.Returns.Request, f.DeleteBucketRequestCall.Returns.DeleteBucketOutput
}
func (f *S3API) DeleteBucketTagging(param1 *s3.DeleteBucketTaggingInput) (*s3.DeleteBucketTaggingOutput, error) {
	f.DeleteBucketTaggingCall.mutex.Lock()
	defer f.DeleteBucketTaggingCall.mutex.Unlock()
	f.DeleteBucketTaggingCall.CallCount++
	f.DeleteBucketTaggingCall.Receives.DeleteBucketTaggingInput = param1
	if f.DeleteBucketTaggingCall.Stub != nil {
		return f.DeleteBucketTaggingCall.Stub(param1)
	}
	return f.DeleteBucketTaggingCall.Returns.DeleteBucketTaggingOutput, f.DeleteBucketTaggingCall.Returns.Error
}
func (f *S3API) DeleteBucketTaggingRequest(param1 *s3.DeleteBucketTaggingInput) (*request.Request, *s3.DeleteBucketTaggingOutput) {
	f.DeleteBucketTaggingRequestCall.mutex.Lock()
	defer f.DeleteBucketTaggingRequestCall.mutex.Unlock()
	f.DeleteBucketTaggingRequestCall.CallCount++
	f.DeleteBucketTaggingRequestCall.Receives.DeleteBucketTaggingInput = param1
	if f.DeleteBucketTaggingRequestCall.Stub != nil {
		return f.DeleteBucketTaggingRequestCall.Stub(param1)
	}
	return f.DeleteBucketTaggingRequestCall.Returns.Request, f.DeleteBucketTaggingRequestCall.Returns.DeleteBucketTaggingOutput
}
func (f *S3API) DeleteBucketTaggingWithContext(param1 context.Context, param2 *s3.DeleteBucketTaggingInput, param3 ...request.Option) (*s3.DeleteBucketTaggingOutput, error) {
	f.DeleteBucketTaggingWithContextCall.mutex.Lock()
	defer f.DeleteBucketTaggingWithContextCall.mutex.Unlock()
	f.DeleteBucketTaggingWithContextCall.CallCount++
	f.DeleteBucketTaggingWithContextCall.Receives.Context = param1
	f.DeleteBucketTaggingWithContextCall.Receives.DeleteBucketTaggingInput = param2
	f.DeleteBucketTaggingWithContextCall.Receives.OptionSlice = param3
	if f.DeleteBucketTaggingWithContextCall.Stub != nil {
		return f.DeleteBucketTaggingWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DeleteBucketTaggingWithContextCall.Returns.DeleteBucketTaggingOutput, f.DeleteBucketTaggingWithContextCall.Returns.Error
}
func (f *S3API) DeleteBucketWebsite(param1 *s3.DeleteBucketWebsiteInput) (*s3.DeleteBucketWebsiteOutput, error) {
	f.DeleteBucketWebsiteCall.mutex.Lock()
	defer f.DeleteBucketWebsiteCall.mutex.Unlock()
	f.DeleteBucketWebsiteCall.CallCount++
	f.DeleteBucketWebsiteCall.Receives.DeleteBucketWebsiteInput = param1
	if f.DeleteBucketWebsiteCall.Stub != nil {
		return f.DeleteBucketWebsiteCall.Stub(param1)
	}
	return f.DeleteBucketWebsiteCall.Returns.DeleteBucketWebsiteOutput, f.DeleteBucketWebsiteCall.Returns.Error
}
func (f *S3API) DeleteBucketWebsiteRequest(param1 *s3.DeleteBucketWebsiteInput) (*request.Request, *s3.DeleteBucketWebsiteOutput) {
	f.DeleteBucketWebsiteRequestCall.mutex.Lock()
	defer f.DeleteBucketWebsiteRequestCall.mutex.Unlock()
	f.DeleteBucketWebsiteRequestCall.CallCount++
	f.DeleteBucketWebsiteRequestCall.Receives.DeleteBucketWebsiteInput = param1
	if f.DeleteBucketWebsiteRequestCall.Stub != nil {
		return f.DeleteBucketWebsiteRequestCall.Stub(param1)
	}
	return f.DeleteBucketWebsiteRequestCall.Returns.Request, f.DeleteBucketWebsiteRequestCall.Returns.DeleteBucketWebsiteOutput
}
func (f *S3API) DeleteBucketWebsiteWithContext(param1 context.Context, param2 *s3.DeleteBucketWebsiteInput, param3 ...request.Option) (*s3.DeleteBucketWebsiteOutput, error) {
	f.DeleteBucketWebsiteWithContextCall.mutex.Lock()
	defer f.DeleteBucketWebsiteWithContextCall.mutex.Unlock()
	f.DeleteBucketWebsiteWithContextCall.CallCount++
	f.DeleteBucketWebsiteWithContextCall.Receives.Context = param1
	f.DeleteBucketWebsiteWithContextCall.Receives.DeleteBucketWebsiteInput = param2
	f.DeleteBucketWebsiteWithContextCall.Receives.OptionSlice = param3
	if f.DeleteBucketWebsiteWithContextCall.Stub != nil {
		return f.DeleteBucketWebsiteWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DeleteBucketWebsiteWithContextCall.Returns.DeleteBucketWebsiteOutput, f.DeleteBucketWebsiteWithContextCall.Returns.Error
}
func (f *S3API) DeleteBucketWithContext(param1 context.Context, param2 *s3.DeleteBucketInput, param3 ...request.Option) (*s3.DeleteBucketOutput, error) {
	f.DeleteBucketWithContextCall.mutex.Lock()
	defer f.DeleteBucketWithContextCall.mutex.Unlock()
	f.DeleteBucketWithContextCall.CallCount++
	f.DeleteBucketWithContextCall.Receives.Context = param1
	f.DeleteBucketWithContextCall.Receives.DeleteBucketInput = param2
	f.DeleteBucketWithContextCall.Receives.OptionSlice = param3
	if f.DeleteBucketWithContextCall.Stub != nil {
		return f.DeleteBucketWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DeleteBucketWithContextCall.Returns.DeleteBucketOutput, f.DeleteBucketWithContextCall.Returns.Error
}
func (f *S3API) DeleteObject(param1 *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	f.DeleteObjectCall.mutex.Lock()
	defer f.DeleteObjectCall.mutex.Unlock()
	f.DeleteObjectCall.CallCount++
	f.DeleteObjectCall.Receives.DeleteObjectInput = param1
	if f.DeleteObjectCall.Stub != nil {
		return f.DeleteObjectCall.Stub(param1)
	}
	return f.DeleteObjectCall.Returns.DeleteObjectOutput, f.DeleteObjectCall.Returns.Error
}
func (f *S3API) DeleteObjectRequest(param1 *s3.DeleteObjectInput) (*request.Request, *s3.DeleteObjectOutput) {
	f.DeleteObjectRequestCall.mutex.Lock()
	defer f.DeleteObjectRequestCall.mutex.Unlock()
	f.DeleteObjectRequestCall.CallCount++
	f.DeleteObjectRequestCall.Receives.DeleteObjectInput = param1
	if f.DeleteObjectRequestCall.Stub != nil {
		return f.DeleteObjectRequestCall.Stub(param1)
	}
	return f.DeleteObjectRequestCall.Returns.Request, f.DeleteObjectRequestCall.Returns.DeleteObjectOutput
}
func (f *S3API) DeleteObjectTagging(param1 *s3.DeleteObjectTaggingInput) (*s3.DeleteObjectTaggingOutput, error) {
	f.DeleteObjectTaggingCall.mutex.Lock()
	defer f.DeleteObjectTaggingCall.mutex.Unlock()
	f.DeleteObjectTaggingCall.CallCount++
	f.DeleteObjectTaggingCall.Receives.DeleteObjectTaggingInput = param1
	if f.DeleteObjectTaggingCall.Stub != nil {
		return f.DeleteObjectTaggingCall.Stub(param1)
	}
	return f.DeleteObjectTaggingCall.Returns.DeleteObjectTaggingOutput, f.DeleteObjectTaggingCall.Returns.Error
}
func (f *S3API) DeleteObjectTaggingRequest(param1 *s3.DeleteObjectTaggingInput) (*request.Request, *s3.DeleteObjectTaggingOutput) {
	f.DeleteObjectTaggingRequestCall.mutex.Lock()
	defer f.DeleteObjectTaggingRequestCall.mutex.Unlock()
	f.DeleteObjectTaggingRequestCall.CallCount++
	f.DeleteObjectTaggingRequestCall.Receives.DeleteObjectTaggingInput = param1
	if f.DeleteObjectTaggingRequestCall.Stub != nil {
		return f.DeleteObjectTaggingRequestCall.Stub(param1)
	}
	return f.DeleteObjectTaggingRequestCall.Returns.Request, f.DeleteObjectTaggingRequestCall.Returns.DeleteObjectTaggingOutput
}
func (f *S3API) DeleteObjectTaggingWithContext(param1 context.Context, param2 *s3.DeleteObjectTaggingInput, param3 ...request.Option) (*s3.DeleteObjectTaggingOutput, error) {
	f.DeleteObjectTaggingWithContextCall.mutex.Lock()
	defer f.DeleteObjectTaggingWithContextCall.mutex.Unlock()
	f.DeleteObjectTaggingWithContextCall.CallCount++
	f.DeleteObjectTaggingWithContextCall.Receives.Context = param1
	f.DeleteObjectTaggingWithContextCall.Receives.DeleteObjectTaggingInput = param2
	f.DeleteObjectTaggingWithContextCall.Receives.OptionSlice = param3
	if f.DeleteObjectTaggingWithContextCall.Stub != nil {
		return f.DeleteObjectTaggingWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DeleteObjectTaggingWithContextCall.Returns.DeleteObjectTaggingOutput, f.DeleteObjectTaggingWithContextCall.Returns.Error
}
func (f *S3API) DeleteObjectWithContext(param1 context.Context, param2 *s3.DeleteObjectInput, param3 ...request.Option) (*s3.DeleteObjectOutput, error) {
	f.DeleteObjectWithContextCall.mutex.Lock()
	defer f.DeleteObjectWithContextCall.mutex.Unlock()
	f.DeleteObjectWithContextCall.CallCount++
	f.DeleteObjectWithContextCall.Receives.Context = param1
	f.DeleteObjectWithContextCall.Receives.DeleteObjectInput = param2
	f.DeleteObjectWithContextCall.Receives.OptionSlice = param3
	if f.DeleteObjectWithContextCall.Stub != nil {
		return f.DeleteObjectWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DeleteObjectWithContextCall.Returns.DeleteObjectOutput, f.DeleteObjectWithContextCall.Returns.Error
}
func (f *S3API) DeleteObjects(param1 *s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error) {
	f.DeleteObjectsCall.mutex.Lock()
	defer f.DeleteObjectsCall.mutex.Unlock()
	f.DeleteObjectsCall.CallCount++
	f.DeleteObjectsCall.Receives.DeleteObjectsInput = param1
	if f.DeleteObjectsCall.Stub != nil {
		return f.DeleteObjectsCall.Stub(param1)
	}
	return f.DeleteObjectsCall.Returns.DeleteObjectsOutput, f.DeleteObjectsCall.Returns.Error
}
func (f *S3API) DeleteObjectsRequest(param1 *s3.DeleteObjectsInput) (*request.Request, *s3.DeleteObjectsOutput) {
	f.DeleteObjectsRequestCall.mutex.Lock()
	defer f.DeleteObjectsRequestCall.mutex.Unlock()
	f.DeleteObjectsRequestCall.CallCount++
	f.DeleteObjectsRequestCall.Receives.DeleteObjectsInput = param1
	if f.DeleteObjectsRequestCall.Stub != nil {
		return f.DeleteObjectsRequestCall.Stub(param1)
	}
	return f.DeleteObjectsRequestCall.Returns.Request, f.DeleteObjectsRequestCall.Returns.DeleteObjectsOutput
}
func (f *S3API) DeleteObjectsWithContext(param1 context.Context, param2 *s3.DeleteObjectsInput, param3 ...request.Option) (*s3.DeleteObjectsOutput, error) {
	f.DeleteObjectsWithContextCall.mutex.Lock()
	defer f.DeleteObjectsWithContextCall.mutex.Unlock()
	f.DeleteObjectsWithContextCall.CallCount++
	f.DeleteObjectsWithContextCall.Receives.Context = param1
	f.DeleteObjectsWithContextCall.Receives.DeleteObjectsInput = param2
	f.DeleteObjectsWithContextCall.Receives.OptionSlice = param3
	if f.DeleteObjectsWithContextCall.Stub != nil {
		return f.DeleteObjectsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DeleteObjectsWithContextCall.Returns.DeleteObjectsOutput, f.DeleteObjectsWithContextCall.Returns.Error
}
func (f *S3API) DeletePublicAccessBlock(param1 *s3.DeletePublicAccessBlockInput) (*s3.DeletePublicAccessBlockOutput, error) {
	f.DeletePublicAccessBlockCall.mutex.Lock()
	defer f.DeletePublicAccessBlockCall.mutex.Unlock()
	f.DeletePublicAccessBlockCall.CallCount++
	f.DeletePublicAccessBlockCall.Receives.DeletePublicAccessBlockInput = param1
	if f.DeletePublicAccessBlockCall.Stub != nil {
		return f.DeletePublicAccessBlockCall.Stub(param1)
	}
	return f.DeletePublicAccessBlockCall.Returns.DeletePublicAccessBlockOutput, f.DeletePublicAccessBlockCall.Returns.Error
}
func (f *S3API) DeletePublicAccessBlockRequest(param1 *s3.DeletePublicAccessBlockInput) (*request.Request, *s3.DeletePublicAccessBlockOutput) {
	f.DeletePublicAccessBlockRequestCall.mutex.Lock()
	defer f.DeletePublicAccessBlockRequestCall.mutex.Unlock()
	f.DeletePublicAccessBlockRequestCall.CallCount++
	f.DeletePublicAccessBlockRequestCall.Receives.DeletePublicAccessBlockInput = param1
	if f.DeletePublicAccessBlockRequestCall.Stub != nil {
		return f.DeletePublicAccessBlockRequestCall.Stub(param1)
	}
	return f.DeletePublicAccessBlockRequestCall.Returns.Request, f.DeletePublicAccessBlockRequestCall.Returns.DeletePublicAccessBlockOutput
}
func (f *S3API) DeletePublicAccessBlockWithContext(param1 context.Context, param2 *s3.DeletePublicAccessBlockInput, param3 ...request.Option) (*s3.DeletePublicAccessBlockOutput, error) {
	f.DeletePublicAccessBlockWithContextCall.mutex.Lock()
	defer f.DeletePublicAccessBlockWithContextCall.mutex.Unlock()
	f.DeletePublicAccessBlockWithContextCall.CallCount++
	f.DeletePublicAccessBlockWithContextCall.Receives.Context = param1
	f.DeletePublicAccessBlockWithContextCall.Receives.DeletePublicAccessBlockInput = param2
	f.DeletePublicAccessBlockWithContextCall.Receives.OptionSlice = param3
	if f.DeletePublicAccessBlockWithContextCall.Stub != nil {
		return f.DeletePublicAccessBlockWithContextCall.Stub(param1, param2, param3...)
	}
	return f.DeletePublicAccessBlockWithContextCall.Returns.DeletePublicAccessBlockOutput, f.DeletePublicAccessBlockWithContextCall.Returns.Error
}
func (f *S3API) GetBucketAccelerateConfiguration(param1 *s3.GetBucketAccelerateConfigurationInput) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	f.GetBucketAccelerateConfigurationCall.mutex.Lock()
	defer f.GetBucketAccelerateConfigurationCall.mutex.Unlock()
	f.GetBucketAccelerateConfigurationCall.CallCount++
	f.GetBucketAccelerateConfigurationCall.Receives.GetBucketAccelerateConfigurationInput = param1
	if f.GetBucketAccelerateConfigurationCall.Stub != nil {
		return f.GetBucketAccelerateConfigurationCall.Stub(param1)
	}
	return f.GetBucketAccelerateConfigurationCall.Returns.GetBucketAccelerateConfigurationOutput, f.GetBucketAccelerateConfigurationCall.Returns.Error
}
func (f *S3API) GetBucketAccelerateConfigurationRequest(param1 *s3.GetBucketAccelerateConfigurationInput) (*request.Request, *s3.GetBucketAccelerateConfigurationOutput) {
	f.GetBucketAccelerateConfigurationRequestCall.mutex.Lock()
	defer f.GetBucketAccelerateConfigurationRequestCall.mutex.Unlock()
	f.GetBucketAccelerateConfigurationRequestCall.CallCount++
	f.GetBucketAccelerateConfigurationRequestCall.Receives.GetBucketAccelerateConfigurationInput = param1
	if f.GetBucketAccelerateConfigurationRequestCall.Stub != nil {
		return f.GetBucketAccelerateConfigurationRequestCall.Stub(param1)
	}
	return f.GetBucketAccelerateConfigurationRequestCall.Returns.Request, f.GetBucketAccelerateConfigurationRequestCall.Returns.GetBucketAccelerateConfigurationOutput
}
func (f *S3API) GetBucketAccelerateConfigurationWithContext(param1 context.Context, param2 *s3.GetBucketAccelerateConfigurationInput, param3 ...request.Option) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	f.GetBucketAccelerateConfigurationWithContextCall.mutex.Lock()
	defer f.GetBucketAccelerateConfigurationWithContextCall.mutex.Unlock()
	f.GetBucketAccelerateConfigurationWithContextCall.CallCount++
	f.GetBucketAccelerateConfigurationWithContextCall.Receives.Context = param1
	f.GetBucketAccelerateConfigurationWithContextCall.Receives.GetBucketAccelerateConfigurationInput = param2
	f.GetBucketAccelerateConfigurationWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketAccelerateConfigurationWithContextCall.Stub != nil {
		return f.GetBucketAccelerateConfigurationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketAccelerateConfigurationWithContextCall.Returns.GetBucketAccelerateConfigurationOutput, f.GetBucketAccelerateConfigurationWithContextCall.Returns.Error
}
func (f *S3API) GetBucketAcl(param1 *s3.GetBucketAclInput) (*s3.GetBucketAclOutput, error) {
	f.GetBucketAclCall.mutex.Lock()
	defer f.GetBucketAclCall.mutex.Unlock()
	f.GetBucketAclCall.CallCount++
	f.GetBucketAclCall.Receives.GetBucketAclInput = param1
	if f.GetBucketAclCall.Stub != nil {
		return f.GetBucketAclCall.Stub(param1)
	}
	return f.GetBucketAclCall.Returns.GetBucketAclOutput, f.GetBucketAclCall.Returns.Error
}
func (f *S3API) GetBucketAclRequest(param1 *s3.GetBucketAclInput) (*request.Request, *s3.GetBucketAclOutput) {
	f.GetBucketAclRequestCall.mutex.Lock()
	defer f.GetBucketAclRequestCall.mutex.Unlock()
	f.GetBucketAclRequestCall.CallCount++
	f.GetBucketAclRequestCall.Receives.GetBucketAclInput = param1
	if f.GetBucketAclRequestCall.Stub != nil {
		return f.GetBucketAclRequestCall.Stub(param1)
	}
	return f.GetBucketAclRequestCall.Returns.Request, f.GetBucketAclRequestCall.Returns.GetBucketAclOutput
}
func (f *S3API) GetBucketAclWithContext(param1 context.Context, param2 *s3.GetBucketAclInput, param3 ...request.Option) (*s3.GetBucketAclOutput, error) {
	f.GetBucketAclWithContextCall.mutex.Lock()
	defer f.GetBucketAclWithContextCall.mutex.Unlock()
	f.GetBucketAclWithContextCall.CallCount++
	f.GetBucketAclWithContextCall.Receives.Context = param1
	f.GetBucketAclWithContextCall.Receives.GetBucketAclInput = param2
	f.GetBucketAclWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketAclWithContextCall.Stub != nil {
		return f.GetBucketAclWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketAclWithContextCall.Returns.GetBucketAclOutput, f.GetBucketAclWithContextCall.Returns.Error
}
func (f *S3API) GetBucketAnalyticsConfiguration(param1 *s3.GetBucketAnalyticsConfigurationInput) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	f.GetBucketAnalyticsConfigurationCall.mutex.Lock()
	defer f.GetBucketAnalyticsConfigurationCall.mutex.Unlock()
	f.GetBucketAnalyticsConfigurationCall.CallCount++
	f.GetBucketAnalyticsConfigurationCall.Receives.GetBucketAnalyticsConfigurationInput = param1
	if f.GetBucketAnalyticsConfigurationCall.Stub != nil {
		return f.GetBucketAnalyticsConfigurationCall.Stub(param1)
	}
	return f.GetBucketAnalyticsConfigurationCall.Returns.GetBucketAnalyticsConfigurationOutput, f.GetBucketAnalyticsConfigurationCall.Returns.Error
}
func (f *S3API) GetBucketAnalyticsConfigurationRequest(param1 *s3.GetBucketAnalyticsConfigurationInput) (*request.Request, *s3.GetBucketAnalyticsConfigurationOutput) {
	f.GetBucketAnalyticsConfigurationRequestCall.mutex.Lock()
	defer f.GetBucketAnalyticsConfigurationRequestCall.mutex.Unlock()
	f.GetBucketAnalyticsConfigurationRequestCall.CallCount++
	f.GetBucketAnalyticsConfigurationRequestCall.Receives.GetBucketAnalyticsConfigurationInput = param1
	if f.GetBucketAnalyticsConfigurationRequestCall.Stub != nil {
		return f.GetBucketAnalyticsConfigurationRequestCall.Stub(param1)
	}
	return f.GetBucketAnalyticsConfigurationRequestCall.Returns.Request, f.GetBucketAnalyticsConfigurationRequestCall.Returns.GetBucketAnalyticsConfigurationOutput
}
func (f *S3API) GetBucketAnalyticsConfigurationWithContext(param1 context.Context, param2 *s3.GetBucketAnalyticsConfigurationInput, param3 ...request.Option) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	f.GetBucketAnalyticsConfigurationWithContextCall.mutex.Lock()
	defer f.GetBucketAnalyticsConfigurationWithContextCall.mutex.Unlock()
	f.GetBucketAnalyticsConfigurationWithContextCall.CallCount++
	f.GetBucketAnalyticsConfigurationWithContextCall.Receives.Context = param1
	f.GetBucketAnalyticsConfigurationWithContextCall.Receives.GetBucketAnalyticsConfigurationInput = param2
	f.GetBucketAnalyticsConfigurationWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketAnalyticsConfigurationWithContextCall.Stub != nil {
		return f.GetBucketAnalyticsConfigurationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketAnalyticsConfigurationWithContextCall.Returns.GetBucketAnalyticsConfigurationOutput, f.GetBucketAnalyticsConfigurationWithContextCall.Returns.Error
}
func (f *S3API) GetBucketCors(param1 *s3.GetBucketCorsInput) (*s3.GetBucketCorsOutput, error) {
	f.GetBucketCorsCall.mutex.Lock()
	defer f.GetBucketCorsCall.mutex.Unlock()
	f.GetBucketCorsCall.CallCount++
	f.GetBucketCorsCall.Receives.GetBucketCorsInput = param1
	if f.GetBucketCorsCall.Stub != nil {
		return f.GetBucketCorsCall.Stub(param1)
	}
	return f.GetBucketCorsCall.Returns.GetBucketCorsOutput, f.GetBucketCorsCall.Returns.Error
}
func (f *S3API) GetBucketCorsRequest(param1 *s3.GetBucketCorsInput) (*request.Request, *s3.GetBucketCorsOutput) {
	f.GetBucketCorsRequestCall.mutex.Lock()
	defer f.GetBucketCorsRequestCall.mutex.Unlock()
	f.GetBucketCorsRequestCall.CallCount++
	f.GetBucketCorsRequestCall.Receives.GetBucketCorsInput = param1
	if f.GetBucketCorsRequestCall.Stub != nil {
		return f.GetBucketCorsRequestCall.Stub(param1)
	}
	return f.GetBucketCorsRequestCall.Returns.Request, f.GetBucketCorsRequestCall.Returns.GetBucketCorsOutput
}
func (f *S3API) GetBucketCorsWithContext(param1 context.Context, param2 *s3.GetBucketCorsInput, param3 ...request.Option) (*s3.GetBucketCorsOutput, error) {
	f.GetBucketCorsWithContextCall.mutex.Lock()
	defer f.GetBucketCorsWithContextCall.mutex.Unlock()
	f.GetBucketCorsWithContextCall.CallCount++
	f.GetBucketCorsWithContextCall.Receives.Context = param1
	f.GetBucketCorsWithContextCall.Receives.GetBucketCorsInput = param2
	f.GetBucketCorsWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketCorsWithContextCall.Stub != nil {
		return f.GetBucketCorsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketCorsWithContextCall.Returns.GetBucketCorsOutput, f.GetBucketCorsWithContextCall.Returns.Error
}
func (f *S3API) GetBucketEncryption(param1 *s3.GetBucketEncryptionInput) (*s3.GetBucketEncryptionOutput, error) {
	f.GetBucketEncryptionCall.mutex.Lock()
	defer f.GetBucketEncryptionCall.mutex.Unlock()
	f.GetBucketEncryptionCall.CallCount++
	f.GetBucketEncryptionCall.Receives.GetBucketEncryptionInput = param1
	if f.GetBucketEncryptionCall.Stub != nil {
		return f.GetBucketEncryptionCall.Stub(param1)
	}
	return f.GetBucketEncryptionCall.Returns.GetBucketEncryptionOutput, f.GetBucketEncryptionCall.Returns.Error
}
func (f *S3API) GetBucketEncryptionRequest(param1 *s3.GetBucketEncryptionInput) (*request.Request, *s3.GetBucketEncryptionOutput) {
	f.GetBucketEncryptionRequestCall.mutex.Lock()
	defer f.GetBucketEncryptionRequestCall.mutex.Unlock()
	f.GetBucketEncryptionRequestCall.CallCount++
	f.GetBucketEncryptionRequestCall.Receives.GetBucketEncryptionInput = param1
	if f.GetBucketEncryptionRequestCall.Stub != nil {
		return f.GetBucketEncryptionRequestCall.Stub(param1)
	}
	return f.GetBucketEncryptionRequestCall.Returns.Request, f.GetBucketEncryptionRequestCall.Returns.GetBucketEncryptionOutput
}
func (f *S3API) GetBucketEncryptionWithContext(param1 context.Context, param2 *s3.GetBucketEncryptionInput, param3 ...request.Option) (*s3.GetBucketEncryptionOutput, error) {
	f.GetBucketEncryptionWithContextCall.mutex.Lock()
	defer f.GetBucketEncryptionWithContextCall.mutex.Unlock()
	f.GetBucketEncryptionWithContextCall.CallCount++
	f.GetBucketEncryptionWithContextCall.Receives.Context = param1
	f.GetBucketEncryptionWithContextCall.Receives.GetBucketEncryptionInput = param2
	f.GetBucketEncryptionWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketEncryptionWithContextCall.Stub != nil {
		return f.GetBucketEncryptionWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketEncryptionWithContextCall.Returns.GetBucketEncryptionOutput, f.GetBucketEncryptionWithContextCall.Returns.Error
}
func (f *S3API) GetBucketIntelligentTieringConfiguration(param1 *s3.GetBucketIntelligentTieringConfigurationInput) (*s3.GetBucketIntelligentTieringConfigurationOutput, error) {
	f.GetBucketIntelligentTieringConfigurationCall.mutex.Lock()
	defer f.GetBucketIntelligentTieringConfigurationCall.mutex.Unlock()
	f.GetBucketIntelligentTieringConfigurationCall.CallCount++
	f.GetBucketIntelligentTieringConfigurationCall.Receives.GetBucketIntelligentTieringConfigurationInput = param1
	if f.GetBucketIntelligentTieringConfigurationCall.Stub != nil {
		return f.GetBucketIntelligentTieringConfigurationCall.Stub(param1)
	}
	return f.GetBucketIntelligentTieringConfigurationCall.Returns.GetBucketIntelligentTieringConfigurationOutput, f.GetBucketIntelligentTieringConfigurationCall.Returns.Error
}
func (f *S3API) GetBucketIntelligentTieringConfigurationRequest(param1 *s3.GetBucketIntelligentTieringConfigurationInput) (*request.Request, *s3.GetBucketIntelligentTieringConfigurationOutput) {
	f.GetBucketIntelligentTieringConfigurationRequestCall.mutex.Lock()
	defer f.GetBucketIntelligentTieringConfigurationRequestCall.mutex.Unlock()
	f.GetBucketIntelligentTieringConfigurationRequestCall.CallCount++
	f.GetBucketIntelligentTieringConfigurationRequestCall.Receives.GetBucketIntelligentTieringConfigurationInput = param1
	if f.GetBucketIntelligentTieringConfigurationRequestCall.Stub != nil {
		return f.GetBucketIntelligentTieringConfigurationRequestCall.Stub(param1)
	}
	return f.GetBucketIntelligentTieringConfigurationRequestCall.Returns.Request, f.GetBucketIntelligentTieringConfigurationRequestCall.Returns.GetBucketIntelligentTieringConfigurationOutput
}
func (f *S3API) GetBucketIntelligentTieringConfigurationWithContext(param1 context.Context, param2 *s3.GetBucketIntelligentTieringConfigurationInput, param3 ...request.Option) (*s3.GetBucketIntelligentTieringConfigurationOutput, error) {
	f.GetBucketIntelligentTieringConfigurationWithContextCall.mutex.Lock()
	defer f.GetBucketIntelligentTieringConfigurationWithContextCall.mutex.Unlock()
	f.GetBucketIntelligentTieringConfigurationWithContextCall.CallCount++
	f.GetBucketIntelligentTieringConfigurationWithContextCall.Receives.Context = param1
	f.GetBucketIntelligentTieringConfigurationWithContextCall.Receives.GetBucketIntelligentTieringConfigurationInput = param2
	f.GetBucketIntelligentTieringConfigurationWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketIntelligentTieringConfigurationWithContextCall.Stub != nil {
		return f.GetBucketIntelligentTieringConfigurationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketIntelligentTieringConfigurationWithContextCall.Returns.GetBucketIntelligentTieringConfigurationOutput, f.GetBucketIntelligentTieringConfigurationWithContextCall.Returns.Error
}
func (f *S3API) GetBucketInventoryConfiguration(param1 *s3.GetBucketInventoryConfigurationInput) (*s3.GetBucketInventoryConfigurationOutput, error) {
	f.GetBucketInventoryConfigurationCall.mutex.Lock()
	defer f.GetBucketInventoryConfigurationCall.mutex.Unlock()
	f.GetBucketInventoryConfigurationCall.CallCount++
	f.GetBucketInventoryConfigurationCall.Receives.GetBucketInventoryConfigurationInput = param1
	if f.GetBucketInventoryConfigurationCall.Stub != nil {
		return f.GetBucketInventoryConfigurationCall.Stub(param1)
	}
	return f.GetBucketInventoryConfigurationCall.Returns.GetBucketInventoryConfigurationOutput, f.GetBucketInventoryConfigurationCall.Returns.Error
}
func (f *S3API) GetBucketInventoryConfigurationRequest(param1 *s3.GetBucketInventoryConfigurationInput) (*request.Request, *s3.GetBucketInventoryConfigurationOutput) {
	f.GetBucketInventoryConfigurationRequestCall.mutex.Lock()
	defer f.GetBucketInventoryConfigurationRequestCall.mutex.Unlock()
	f.GetBucketInventoryConfigurationRequestCall.CallCount++
	f.GetBucketInventoryConfigurationRequestCall.Receives.GetBucketInventoryConfigurationInput = param1
	if f.GetBucketInventoryConfigurationRequestCall.Stub != nil {
		return f.GetBucketInventoryConfigurationRequestCall.Stub(param1)
	}
	return f.GetBucketInventoryConfigurationRequestCall.Returns.Request, f.GetBucketInventoryConfigurationRequestCall.Returns.GetBucketInventoryConfigurationOutput
}
func (f *S3API) GetBucketInventoryConfigurationWithContext(param1 context.Context, param2 *s3.GetBucketInventoryConfigurationInput, param3 ...request.Option) (*s3.GetBucketInventoryConfigurationOutput, error) {
	f.GetBucketInventoryConfigurationWithContextCall.mutex.Lock()
	defer f.GetBucketInventoryConfigurationWithContextCall.mutex.Unlock()
	f.GetBucketInventoryConfigurationWithContextCall.CallCount++
	f.GetBucketInventoryConfigurationWithContextCall.Receives.Context = param1
	f.GetBucketInventoryConfigurationWithContextCall.Receives.GetBucketInventoryConfigurationInput = param2
	f.GetBucketInventoryConfigurationWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketInventoryConfigurationWithContextCall.Stub != nil {
		return f.GetBucketInventoryConfigurationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketInventoryConfigurationWithContextCall.Returns.GetBucketInventoryConfigurationOutput, f.GetBucketInventoryConfigurationWithContextCall.Returns.Error
}
func (f *S3API) GetBucketLifecycle(param1 *s3.GetBucketLifecycleInput) (*s3.GetBucketLifecycleOutput, error) {
	f.GetBucketLifecycleCall.mutex.Lock()
	defer f.GetBucketLifecycleCall.mutex.Unlock()
	f.GetBucketLifecycleCall.CallCount++
	f.GetBucketLifecycleCall.Receives.GetBucketLifecycleInput = param1
	if f.GetBucketLifecycleCall.Stub != nil {
		return f.GetBucketLifecycleCall.Stub(param1)
	}
	return f.GetBucketLifecycleCall.Returns.GetBucketLifecycleOutput, f.GetBucketLifecycleCall.Returns.Error
}
func (f *S3API) GetBucketLifecycleConfiguration(param1 *s3.GetBucketLifecycleConfigurationInput) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	f.GetBucketLifecycleConfigurationCall.mutex.Lock()
	defer f.GetBucketLifecycleConfigurationCall.mutex.Unlock()
	f.GetBucketLifecycleConfigurationCall.CallCount++
	f.GetBucketLifecycleConfigurationCall.Receives.GetBucketLifecycleConfigurationInput = param1
	if f.GetBucketLifecycleConfigurationCall.Stub != nil {
		return f.GetBucketLifecycleConfigurationCall.Stub(param1)
	}
	return f.GetBucketLifecycleConfigurationCall.Returns.GetBucketLifecycleConfigurationOutput, f.GetBucketLifecycleConfigurationCall.Returns.Error
}
func (f *S3API) GetBucketLifecycleConfigurationRequest(param1 *s3.GetBucketLifecycleConfigurationInput) (*request.Request, *s3.GetBucketLifecycleConfigurationOutput) {
	f.GetBucketLifecycleConfigurationRequestCall.mutex.Lock()
	defer f.GetBucketLifecycleConfigurationRequestCall.mutex.Unlock()
	f.GetBucketLifecycleConfigurationRequestCall.CallCount++
	f.GetBucketLifecycleConfigurationRequestCall.Receives.GetBucketLifecycleConfigurationInput = param1
	if f.GetBucketLifecycleConfigurationRequestCall.Stub != nil {
		return f.GetBucketLifecycleConfigurationRequestCall.Stub(param1)
	}
	return f.GetBucketLifecycleConfigurationRequestCall.Returns.Request, f.GetBucketLifecycleConfigurationRequestCall.Returns.GetBucketLifecycleConfigurationOutput
}
func (f *S3API) GetBucketLifecycleConfigurationWithContext(param1 context.Context, param2 *s3.GetBucketLifecycleConfigurationInput, param3 ...request.Option) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	f.GetBucketLifecycleConfigurationWithContextCall.mutex.Lock()
	defer f.GetBucketLifecycleConfigurationWithContextCall.mutex.Unlock()
	f.GetBucketLifecycleConfigurationWithContextCall.CallCount++
	f.GetBucketLifecycleConfigurationWithContextCall.Receives.Context = param1
	f.GetBucketLifecycleConfigurationWithContextCall.Receives.GetBucketLifecycleConfigurationInput = param2
	f.GetBucketLifecycleConfigurationWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketLifecycleConfigurationWithContextCall.Stub != nil {
		return f.GetBucketLifecycleConfigurationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketLifecycleConfigurationWithContextCall.Returns.GetBucketLifecycleConfigurationOutput, f.GetBucketLifecycleConfigurationWithContextCall.Returns.Error
}
func (f *S3API) GetBucketLifecycleRequest(param1 *s3.GetBucketLifecycleInput) (*request.Request, *s3.GetBucketLifecycleOutput) {
	f.GetBucketLifecycleRequestCall.mutex.Lock()
	defer f.GetBucketLifecycleRequestCall.mutex.Unlock()
	f.GetBucketLifecycleRequestCall.CallCount++
	f.GetBucketLifecycleRequestCall.Receives.GetBucketLifecycleInput = param1
	if f.GetBucketLifecycleRequestCall.Stub != nil {
		return f.GetBucketLifecycleRequestCall.Stub(param1)
	}
	return f.GetBucketLifecycleRequestCall.Returns.Request, f.GetBucketLifecycleRequestCall.Returns.GetBucketLifecycleOutput
}
func (f *S3API) GetBucketLifecycleWithContext(param1 context.Context, param2 *s3.GetBucketLifecycleInput, param3 ...request.Option) (*s3.GetBucketLifecycleOutput, error) {
	f.GetBucketLifecycleWithContextCall.mutex.Lock()
	defer f.GetBucketLifecycleWithContextCall.mutex.Unlock()
	f.GetBucketLifecycleWithContextCall.CallCount++
	f.GetBucketLifecycleWithContextCall.Receives.Context = param1
	f.GetBucketLifecycleWithContextCall.Receives.GetBucketLifecycleInput = param2
	f.GetBucketLifecycleWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketLifecycleWithContextCall.Stub != nil {
		return f.GetBucketLifecycleWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketLifecycleWithContextCall.Returns.GetBucketLifecycleOutput, f.GetBucketLifecycleWithContextCall.Returns.Error
}
func (f *S3API) GetBucketLocation(param1 *s3.GetBucketLocationInput) (*s3.GetBucketLocationOutput, error) {
	f.GetBucketLocationCall.mutex.Lock()
	defer f.GetBucketLocationCall.mutex.Unlock()
	f.GetBucketLocationCall.CallCount++
	f.GetBucketLocationCall.Receives.GetBucketLocationInput = param1
	if f.GetBucketLocationCall.Stub != nil {
		return f.GetBucketLocationCall.Stub(param1)
	}
	return f.GetBucketLocationCall.Returns.GetBucketLocationOutput, f.GetBucketLocationCall.Returns.Error
}
func (f *S3API) GetBucketLocationRequest(param1 *s3.GetBucketLocationInput) (*request.Request, *s3.GetBucketLocationOutput) {
	f.GetBucketLocationRequestCall.mutex.Lock()
	defer f.GetBucketLocationRequestCall.mutex.Unlock()
	f.GetBucketLocationRequestCall.CallCount++
	f.GetBucketLocationRequestCall.Receives.GetBucketLocationInput = param1
	if f.GetBucketLocationRequestCall.Stub != nil {
		return f.GetBucketLocationRequestCall.Stub(param1)
	}
	return f.GetBucketLocationRequestCall.Returns.Request, f.GetBucketLocationRequestCall.Returns.GetBucketLocationOutput
}
func (f *S3API) GetBucketLocationWithContext(param1 context.Context, param2 *s3.GetBucketLocationInput, param3 ...request.Option) (*s3.GetBucketLocationOutput, error) {
	f.GetBucketLocationWithContextCall.mutex.Lock()
	defer f.GetBucketLocationWithContextCall.mutex.Unlock()
	f.GetBucketLocationWithContextCall.CallCount++
	f.GetBucketLocationWithContextCall.Receives.Context = param1
	f.GetBucketLocationWithContextCall.Receives.GetBucketLocationInput = param2
	f.GetBucketLocationWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketLocationWithContextCall.Stub != nil {
		return f.GetBucketLocationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketLocationWithContextCall.Returns.GetBucketLocationOutput, f.GetBucketLocationWithContextCall.Returns.Error
}
func (f *S3API) GetBucketLogging(param1 *s3.GetBucketLoggingInput) (*s3.GetBucketLoggingOutput, error) {
	f.GetBucketLoggingCall.mutex.Lock()
	defer f.GetBucketLoggingCall.mutex.Unlock()
	f.GetBucketLoggingCall.CallCount++
	f.GetBucketLoggingCall.Receives.GetBucketLoggingInput = param1
	if f.GetBucketLoggingCall.Stub != nil {
		return f.GetBucketLoggingCall.Stub(param1)
	}
	return f.GetBucketLoggingCall.Returns.GetBucketLoggingOutput, f.GetBucketLoggingCall.Returns.Error
}
func (f *S3API) GetBucketLoggingRequest(param1 *s3.GetBucketLoggingInput) (*request.Request, *s3.GetBucketLoggingOutput) {
	f.GetBucketLoggingRequestCall.mutex.Lock()
	defer f.GetBucketLoggingRequestCall.mutex.Unlock()
	f.GetBucketLoggingRequestCall.CallCount++
	f.GetBucketLoggingRequestCall.Receives.GetBucketLoggingInput = param1
	if f.GetBucketLoggingRequestCall.Stub != nil {
		return f.GetBucketLoggingRequestCall.Stub(param1)
	}
	return f.GetBucketLoggingRequestCall.Returns.Request, f.GetBucketLoggingRequestCall.Returns.GetBucketLoggingOutput
}
func (f *S3API) GetBucketLoggingWithContext(param1 context.Context, param2 *s3.GetBucketLoggingInput, param3 ...request.Option) (*s3.GetBucketLoggingOutput, error) {
	f.GetBucketLoggingWithContextCall.mutex.Lock()
	defer f.GetBucketLoggingWithContextCall.mutex.Unlock()
	f.GetBucketLoggingWithContextCall.CallCount++
	f.GetBucketLoggingWithContextCall.Receives.Context = param1
	f.GetBucketLoggingWithContextCall.Receives.GetBucketLoggingInput = param2
	f.GetBucketLoggingWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketLoggingWithContextCall.Stub != nil {
		return f.GetBucketLoggingWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketLoggingWithContextCall.Returns.GetBucketLoggingOutput, f.GetBucketLoggingWithContextCall.Returns.Error
}
func (f *S3API) GetBucketMetricsConfiguration(param1 *s3.GetBucketMetricsConfigurationInput) (*s3.GetBucketMetricsConfigurationOutput, error) {
	f.GetBucketMetricsConfigurationCall.mutex.Lock()
	defer f.GetBucketMetricsConfigurationCall.mutex.Unlock()
	f.GetBucketMetricsConfigurationCall.CallCount++
	f.GetBucketMetricsConfigurationCall.Receives.GetBucketMetricsConfigurationInput = param1
	if f.GetBucketMetricsConfigurationCall.Stub != nil {
		return f.GetBucketMetricsConfigurationCall.Stub(param1)
	}
	return f.GetBucketMetricsConfigurationCall.Returns.GetBucketMetricsConfigurationOutput, f.GetBucketMetricsConfigurationCall.Returns.Error
}
func (f *S3API) GetBucketMetricsConfigurationRequest(param1 *s3.GetBucketMetricsConfigurationInput) (*request.Request, *s3.GetBucketMetricsConfigurationOutput) {
	f.GetBucketMetricsConfigurationRequestCall.mutex.Lock()
	defer f.GetBucketMetricsConfigurationRequestCall.mutex.Unlock()
	f.GetBucketMetricsConfigurationRequestCall.CallCount++
	f.GetBucketMetricsConfigurationRequestCall.Receives.GetBucketMetricsConfigurationInput = param1
	if f.GetBucketMetricsConfigurationRequestCall.Stub != nil {
		return f.GetBucketMetricsConfigurationRequestCall.Stub(param1)
	}
	return f.GetBucketMetricsConfigurationRequestCall.Returns.Request, f.GetBucketMetricsConfigurationRequestCall.Returns.GetBucketMetricsConfigurationOutput
}
func (f *S3API) GetBucketMetricsConfigurationWithContext(param1 context.Context, param2 *s3.GetBucketMetricsConfigurationInput, param3 ...request.Option) (*s3.GetBucketMetricsConfigurationOutput, error) {
	f.GetBucketMetricsConfigurationWithContextCall.mutex.Lock()
	defer f.GetBucketMetricsConfigurationWithContextCall.mutex.Unlock()
	f.GetBucketMetricsConfigurationWithContextCall.CallCount++
	f.GetBucketMetricsConfigurationWithContextCall.Receives.Context = param1
	f.GetBucketMetricsConfigurationWithContextCall.Receives.GetBucketMetricsConfigurationInput = param2
	f.GetBucketMetricsConfigurationWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketMetricsConfigurationWithContextCall.Stub != nil {
		return f.GetBucketMetricsConfigurationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketMetricsConfigurationWithContextCall.Returns.GetBucketMetricsConfigurationOutput, f.GetBucketMetricsConfigurationWithContextCall.Returns.Error
}
func (f *S3API) GetBucketNotification(param1 *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfigurationDeprecated, error) {
	f.GetBucketNotificationCall.mutex.Lock()
	defer f.GetBucketNotificationCall.mutex.Unlock()
	f.GetBucketNotificationCall.CallCount++
	f.GetBucketNotificationCall.Receives.GetBucketNotificationConfigurationRequest = param1
	if f.GetBucketNotificationCall.Stub != nil {
		return f.GetBucketNotificationCall.Stub(param1)
	}
	return f.GetBucketNotificationCall.Returns.NotificationConfigurationDeprecated, f.GetBucketNotificationCall.Returns.Error
}
func (f *S3API) GetBucketNotificationConfiguration(param1 *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfiguration, error) {
	f.GetBucketNotificationConfigurationCall.mutex.Lock()
	defer f.GetBucketNotificationConfigurationCall.mutex.Unlock()
	f.GetBucketNotificationConfigurationCall.CallCount++
	f.GetBucketNotificationConfigurationCall.Receives.GetBucketNotificationConfigurationRequest = param1
	if f.GetBucketNotificationConfigurationCall.Stub != nil {
		return f.GetBucketNotificationConfigurationCall.Stub(param1)
	}
	return f.GetBucketNotificationConfigurationCall.Returns.NotificationConfiguration, f.GetBucketNotificationConfigurationCall.Returns.Error
}
func (f *S3API) GetBucketNotificationConfigurationRequest(param1 *s3.GetBucketNotificationConfigurationRequest) (*request.Request, *s3.NotificationConfiguration) {
	f.GetBucketNotificationConfigurationRequestCall.mutex.Lock()
	defer f.GetBucketNotificationConfigurationRequestCall.mutex.Unlock()
	f.GetBucketNotificationConfigurationRequestCall.CallCount++
	f.GetBucketNotificationConfigurationRequestCall.Receives.GetBucketNotificationConfigurationRequest = param1
	if f.GetBucketNotificationConfigurationRequestCall.Stub != nil {
		return f.GetBucketNotificationConfigurationRequestCall.Stub(param1)
	}
	return f.GetBucketNotificationConfigurationRequestCall.Returns.Request, f.GetBucketNotificationConfigurationRequestCall.Returns.NotificationConfiguration
}
func (f *S3API) GetBucketNotificationConfigurationWithContext(param1 context.Context, param2 *s3.GetBucketNotificationConfigurationRequest, param3 ...request.Option) (*s3.NotificationConfiguration, error) {
	f.GetBucketNotificationConfigurationWithContextCall.mutex.Lock()
	defer f.GetBucketNotificationConfigurationWithContextCall.mutex.Unlock()
	f.GetBucketNotificationConfigurationWithContextCall.CallCount++
	f.GetBucketNotificationConfigurationWithContextCall.Receives.Context = param1
	f.GetBucketNotificationConfigurationWithContextCall.Receives.GetBucketNotificationConfigurationRequest = param2
	f.GetBucketNotificationConfigurationWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketNotificationConfigurationWithContextCall.Stub != nil {
		return f.GetBucketNotificationConfigurationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketNotificationConfigurationWithContextCall.Returns.NotificationConfiguration, f.GetBucketNotificationConfigurationWithContextCall.Returns.Error
}
func (f *S3API) GetBucketNotificationRequest(param1 *s3.GetBucketNotificationConfigurationRequest) (*request.Request, *s3.NotificationConfigurationDeprecated) {
	f.GetBucketNotificationRequestCall.mutex.Lock()
	defer f.GetBucketNotificationRequestCall.mutex.Unlock()
	f.GetBucketNotificationRequestCall.CallCount++
	f.GetBucketNotificationRequestCall.Receives.GetBucketNotificationConfigurationRequest = param1
	if f.GetBucketNotificationRequestCall.Stub != nil {
		return f.GetBucketNotificationRequestCall.Stub(param1)
	}
	return f.GetBucketNotificationRequestCall.Returns.Request, f.GetBucketNotificationRequestCall.Returns.NotificationConfigurationDeprecated
}
func (f *S3API) GetBucketNotificationWithContext(param1 context.Context, param2 *s3.GetBucketNotificationConfigurationRequest, param3 ...request.Option) (*s3.NotificationConfigurationDeprecated, error) {
	f.GetBucketNotificationWithContextCall.mutex.Lock()
	defer f.GetBucketNotificationWithContextCall.mutex.Unlock()
	f.GetBucketNotificationWithContextCall.CallCount++
	f.GetBucketNotificationWithContextCall.Receives.Context = param1
	f.GetBucketNotificationWithContextCall.Receives.GetBucketNotificationConfigurationRequest = param2
	f.GetBucketNotificationWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketNotificationWithContextCall.Stub != nil {
		return f.GetBucketNotificationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketNotificationWithContextCall.Returns.NotificationConfigurationDeprecated, f.GetBucketNotificationWithContextCall.Returns.Error
}
func (f *S3API) GetBucketOwnershipControls(param1 *s3.GetBucketOwnershipControlsInput) (*s3.GetBucketOwnershipControlsOutput, error) {
	f.GetBucketOwnershipControlsCall.mutex.Lock()
	defer f.GetBucketOwnershipControlsCall.mutex.Unlock()
	f.GetBucketOwnershipControlsCall.CallCount++
	f.GetBucketOwnershipControlsCall.Receives.GetBucketOwnershipControlsInput = param1
	if f.GetBucketOwnershipControlsCall.Stub != nil {
		return f.GetBucketOwnershipControlsCall.Stub(param1)
	}
	return f.GetBucketOwnershipControlsCall.Returns.GetBucketOwnershipControlsOutput, f.GetBucketOwnershipControlsCall.Returns.Error
}
func (f *S3API) GetBucketOwnershipControlsRequest(param1 *s3.GetBucketOwnershipControlsInput) (*request.Request, *s3.GetBucketOwnershipControlsOutput) {
	f.GetBucketOwnershipControlsRequestCall.mutex.Lock()
	defer f.GetBucketOwnershipControlsRequestCall.mutex.Unlock()
	f.GetBucketOwnershipControlsRequestCall.CallCount++
	f.GetBucketOwnershipControlsRequestCall.Receives.GetBucketOwnershipControlsInput = param1
	if f.GetBucketOwnershipControlsRequestCall.Stub != nil {
		return f.GetBucketOwnershipControlsRequestCall.Stub(param1)
	}
	return f.GetBucketOwnershipControlsRequestCall.Returns.Request, f.GetBucketOwnershipControlsRequestCall.Returns.GetBucketOwnershipControlsOutput
}
func (f *S3API) GetBucketOwnershipControlsWithContext(param1 context.Context, param2 *s3.GetBucketOwnershipControlsInput, param3 ...request.Option) (*s3.GetBucketOwnershipControlsOutput, error) {
	f.GetBucketOwnershipControlsWithContextCall.mutex.Lock()
	defer f.GetBucketOwnershipControlsWithContextCall.mutex.Unlock()
	f.GetBucketOwnershipControlsWithContextCall.CallCount++
	f.GetBucketOwnershipControlsWithContextCall.Receives.Context = param1
	f.GetBucketOwnershipControlsWithContextCall.Receives.GetBucketOwnershipControlsInput = param2
	f.GetBucketOwnershipControlsWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketOwnershipControlsWithContextCall.Stub != nil {
		return f.GetBucketOwnershipControlsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketOwnershipControlsWithContextCall.Returns.GetBucketOwnershipControlsOutput, f.GetBucketOwnershipControlsWithContextCall.Returns.Error
}
func (f *S3API) GetBucketPolicy(param1 *s3.GetBucketPolicyInput) (*s3.GetBucketPolicyOutput, error) {
	f.GetBucketPolicyCall.mutex.Lock()
	defer f.GetBucketPolicyCall.mutex.Unlock()
	f.GetBucketPolicyCall.CallCount++
	f.GetBucketPolicyCall.Receives.GetBucketPolicyInput = param1
	if f.GetBucketPolicyCall.Stub != nil {
		return f.GetBucketPolicyCall.Stub(param1)
	}
	return f.GetBucketPolicyCall.Returns.GetBucketPolicyOutput, f.GetBucketPolicyCall.Returns.Error
}
func (f *S3API) GetBucketPolicyRequest(param1 *s3.GetBucketPolicyInput) (*request.Request, *s3.GetBucketPolicyOutput) {
	f.GetBucketPolicyRequestCall.mutex.Lock()
	defer f.GetBucketPolicyRequestCall.mutex.Unlock()
	f.GetBucketPolicyRequestCall.CallCount++
	f.GetBucketPolicyRequestCall.Receives.GetBucketPolicyInput = param1
	if f.GetBucketPolicyRequestCall.Stub != nil {
		return f.GetBucketPolicyRequestCall.Stub(param1)
	}
	return f.GetBucketPolicyRequestCall.Returns.Request, f.GetBucketPolicyRequestCall.Returns.GetBucketPolicyOutput
}
func (f *S3API) GetBucketPolicyStatus(param1 *s3.GetBucketPolicyStatusInput) (*s3.GetBucketPolicyStatusOutput, error) {
	f.GetBucketPolicyStatusCall.mutex.Lock()
	defer f.GetBucketPolicyStatusCall.mutex.Unlock()
	f.GetBucketPolicyStatusCall.CallCount++
	f.GetBucketPolicyStatusCall.Receives.GetBucketPolicyStatusInput = param1
	if f.GetBucketPolicyStatusCall.Stub != nil {
		return f.GetBucketPolicyStatusCall.Stub(param1)
	}
	return f.GetBucketPolicyStatusCall.Returns.GetBucketPolicyStatusOutput, f.GetBucketPolicyStatusCall.Returns.Error
}
func (f *S3API) GetBucketPolicyStatusRequest(param1 *s3.GetBucketPolicyStatusInput) (*request.Request, *s3.GetBucketPolicyStatusOutput) {
	f.GetBucketPolicyStatusRequestCall.mutex.Lock()
	defer f.GetBucketPolicyStatusRequestCall.mutex.Unlock()
	f.GetBucketPolicyStatusRequestCall.CallCount++
	f.GetBucketPolicyStatusRequestCall.Receives.GetBucketPolicyStatusInput = param1
	if f.GetBucketPolicyStatusRequestCall.Stub != nil {
		return f.GetBucketPolicyStatusRequestCall.Stub(param1)
	}
	return f.GetBucketPolicyStatusRequestCall.Returns.Request, f.GetBucketPolicyStatusRequestCall.Returns.GetBucketPolicyStatusOutput
}
func (f *S3API) GetBucketPolicyStatusWithContext(param1 context.Context, param2 *s3.GetBucketPolicyStatusInput, param3 ...request.Option) (*s3.GetBucketPolicyStatusOutput, error) {
	f.GetBucketPolicyStatusWithContextCall.mutex.Lock()
	defer f.GetBucketPolicyStatusWithContextCall.mutex.Unlock()
	f.GetBucketPolicyStatusWithContextCall.CallCount++
	f.GetBucketPolicyStatusWithContextCall.Receives.Context = param1
	f.GetBucketPolicyStatusWithContextCall.Receives.GetBucketPolicyStatusInput = param2
	f.GetBucketPolicyStatusWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketPolicyStatusWithContextCall.Stub != nil {
		return f.GetBucketPolicyStatusWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketPolicyStatusWithContextCall.Returns.GetBucketPolicyStatusOutput, f.GetBucketPolicyStatusWithContextCall.Returns.Error
}
func (f *S3API) GetBucketPolicyWithContext(param1 context.Context, param2 *s3.GetBucketPolicyInput, param3 ...request.Option) (*s3.GetBucketPolicyOutput, error) {
	f.GetBucketPolicyWithContextCall.mutex.Lock()
	defer f.GetBucketPolicyWithContextCall.mutex.Unlock()
	f.GetBucketPolicyWithContextCall.CallCount++
	f.GetBucketPolicyWithContextCall.Receives.Context = param1
	f.GetBucketPolicyWithContextCall.Receives.GetBucketPolicyInput = param2
	f.GetBucketPolicyWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketPolicyWithContextCall.Stub != nil {
		return f.GetBucketPolicyWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketPolicyWithContextCall.Returns.GetBucketPolicyOutput, f.GetBucketPolicyWithContextCall.Returns.Error
}
func (f *S3API) GetBucketReplication(param1 *s3.GetBucketReplicationInput) (*s3.GetBucketReplicationOutput, error) {
	f.GetBucketReplicationCall.mutex.Lock()
	defer f.GetBucketReplicationCall.mutex.Unlock()
	f.GetBucketReplicationCall.CallCount++
	f.GetBucketReplicationCall.Receives.GetBucketReplicationInput = param1
	if f.GetBucketReplicationCall.Stub != nil {
		return f.GetBucketReplicationCall.Stub(param1)
	}
	return f.GetBucketReplicationCall.Returns.GetBucketReplicationOutput, f.GetBucketReplicationCall.Returns.Error
}
func (f *S3API) GetBucketReplicationRequest(param1 *s3.GetBucketReplicationInput) (*request.Request, *s3.GetBucketReplicationOutput) {
	f.GetBucketReplicationRequestCall.mutex.Lock()
	defer f.GetBucketReplicationRequestCall.mutex.Unlock()
	f.GetBucketReplicationRequestCall.CallCount++
	f.GetBucketReplicationRequestCall.Receives.GetBucketReplicationInput = param1
	if f.GetBucketReplicationRequestCall.Stub != nil {
		return f.GetBucketReplicationRequestCall.Stub(param1)
	}
	return f.GetBucketReplicationRequestCall.Returns.Request, f.GetBucketReplicationRequestCall.Returns.GetBucketReplicationOutput
}
func (f *S3API) GetBucketReplicationWithContext(param1 context.Context, param2 *s3.GetBucketReplicationInput, param3 ...request.Option) (*s3.GetBucketReplicationOutput, error) {
	f.GetBucketReplicationWithContextCall.mutex.Lock()
	defer f.GetBucketReplicationWithContextCall.mutex.Unlock()
	f.GetBucketReplicationWithContextCall.CallCount++
	f.GetBucketReplicationWithContextCall.Receives.Context = param1
	f.GetBucketReplicationWithContextCall.Receives.GetBucketReplicationInput = param2
	f.GetBucketReplicationWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketReplicationWithContextCall.Stub != nil {
		return f.GetBucketReplicationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketReplicationWithContextCall.Returns.GetBucketReplicationOutput, f.GetBucketReplicationWithContextCall.Returns.Error
}
func (f *S3API) GetBucketRequestPayment(param1 *s3.GetBucketRequestPaymentInput) (*s3.GetBucketRequestPaymentOutput, error) {
	f.GetBucketRequestPaymentCall.mutex.Lock()
	defer f.GetBucketRequestPaymentCall.mutex.Unlock()
	f.GetBucketRequestPaymentCall.CallCount++
	f.GetBucketRequestPaymentCall.Receives.GetBucketRequestPaymentInput = param1
	if f.GetBucketRequestPaymentCall.Stub != nil {
		return f.GetBucketRequestPaymentCall.Stub(param1)
	}
	return f.GetBucketRequestPaymentCall.Returns.GetBucketRequestPaymentOutput, f.GetBucketRequestPaymentCall.Returns.Error
}
func (f *S3API) GetBucketRequestPaymentRequest(param1 *s3.GetBucketRequestPaymentInput) (*request.Request, *s3.GetBucketRequestPaymentOutput) {
	f.GetBucketRequestPaymentRequestCall.mutex.Lock()
	defer f.GetBucketRequestPaymentRequestCall.mutex.Unlock()
	f.GetBucketRequestPaymentRequestCall.CallCount++
	f.GetBucketRequestPaymentRequestCall.Receives.GetBucketRequestPaymentInput = param1
	if f.GetBucketRequestPaymentRequestCall.Stub != nil {
		return f.GetBucketRequestPaymentRequestCall.Stub(param1)
	}
	return f.GetBucketRequestPaymentRequestCall.Returns.Request, f.GetBucketRequestPaymentRequestCall.Returns.GetBucketRequestPaymentOutput
}
func (f *S3API) GetBucketRequestPaymentWithContext(param1 context.Context, param2 *s3.GetBucketRequestPaymentInput, param3 ...request.Option) (*s3.GetBucketRequestPaymentOutput, error) {
	f.GetBucketRequestPaymentWithContextCall.mutex.Lock()
	defer f.GetBucketRequestPaymentWithContextCall.mutex.Unlock()
	f.GetBucketRequestPaymentWithContextCall.CallCount++
	f.GetBucketRequestPaymentWithContextCall.Receives.Context = param1
	f.GetBucketRequestPaymentWithContextCall.Receives.GetBucketRequestPaymentInput = param2
	f.GetBucketRequestPaymentWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketRequestPaymentWithContextCall.Stub != nil {
		return f.GetBucketRequestPaymentWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketRequestPaymentWithContextCall.Returns.GetBucketRequestPaymentOutput, f.GetBucketRequestPaymentWithContextCall.Returns.Error
}
func (f *S3API) GetBucketTagging(param1 *s3.GetBucketTaggingInput) (*s3.GetBucketTaggingOutput, error) {
	f.GetBucketTaggingCall.mutex.Lock()
	defer f.GetBucketTaggingCall.mutex.Unlock()
	f.GetBucketTaggingCall.CallCount++
	f.GetBucketTaggingCall.Receives.GetBucketTaggingInput = param1
	if f.GetBucketTaggingCall.Stub != nil {
		return f.GetBucketTaggingCall.Stub(param1)
	}
	return f.GetBucketTaggingCall.Returns.GetBucketTaggingOutput, f.GetBucketTaggingCall.Returns.Error
}
func (f *S3API) GetBucketTaggingRequest(param1 *s3.GetBucketTaggingInput) (*request.Request, *s3.GetBucketTaggingOutput) {
	f.GetBucketTaggingRequestCall.mutex.Lock()
	defer f.GetBucketTaggingRequestCall.mutex.Unlock()
	f.GetBucketTaggingRequestCall.CallCount++
	f.GetBucketTaggingRequestCall.Receives.GetBucketTaggingInput = param1
	if f.GetBucketTaggingRequestCall.Stub != nil {
		return f.GetBucketTaggingRequestCall.Stub(param1)
	}
	return f.GetBucketTaggingRequestCall.Returns.Request, f.GetBucketTaggingRequestCall.Returns.GetBucketTaggingOutput
}
func (f *S3API) GetBucketTaggingWithContext(param1 context.Context, param2 *s3.GetBucketTaggingInput, param3 ...request.Option) (*s3.GetBucketTaggingOutput, error) {
	f.GetBucketTaggingWithContextCall.mutex.Lock()
	defer f.GetBucketTaggingWithContextCall.mutex.Unlock()
	f.GetBucketTaggingWithContextCall.CallCount++
	f.GetBucketTaggingWithContextCall.Receives.Context = param1
	f.GetBucketTaggingWithContextCall.Receives.GetBucketTaggingInput = param2
	f.GetBucketTaggingWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketTaggingWithContextCall.Stub != nil {
		return f.GetBucketTaggingWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketTaggingWithContextCall.Returns.GetBucketTaggingOutput, f.GetBucketTaggingWithContextCall.Returns.Error
}
func (f *S3API) GetBucketVersioning(param1 *s3.GetBucketVersioningInput) (*s3.GetBucketVersioningOutput, error) {
	f.GetBucketVersioningCall.mutex.Lock()
	defer f.GetBucketVersioningCall.mutex.Unlock()
	f.GetBucketVersioningCall.CallCount++
	f.GetBucketVersioningCall.Receives.GetBucketVersioningInput = param1
	if f.GetBucketVersioningCall.Stub != nil {
		return f.GetBucketVersioningCall.Stub(param1)
	}
	return f.GetBucketVersioningCall.Returns.GetBucketVersioningOutput, f.GetBucketVersioningCall.Returns.Error
}
func (f *S3API) GetBucketVersioningRequest(param1 *s3.GetBucketVersioningInput) (*request.Request, *s3.GetBucketVersioningOutput) {
	f.GetBucketVersioningRequestCall.mutex.Lock()
	defer f.GetBucketVersioningRequestCall.mutex.Unlock()
	f.GetBucketVersioningRequestCall.CallCount++
	f.GetBucketVersioningRequestCall.Receives.GetBucketVersioningInput = param1
	if f.GetBucketVersioningRequestCall.Stub != nil {
		return f.GetBucketVersioningRequestCall.Stub(param1)
	}
	return f.GetBucketVersioningRequestCall.Returns.Request, f.GetBucketVersioningRequestCall.Returns.GetBucketVersioningOutput
}
func (f *S3API) GetBucketVersioningWithContext(param1 context.Context, param2 *s3.GetBucketVersioningInput, param3 ...request.Option) (*s3.GetBucketVersioningOutput, error) {
	f.GetBucketVersioningWithContextCall.mutex.Lock()
	defer f.GetBucketVersioningWithContextCall.mutex.Unlock()
	f.GetBucketVersioningWithContextCall.CallCount++
	f.GetBucketVersioningWithContextCall.Receives.Context = param1
	f.GetBucketVersioningWithContextCall.Receives.GetBucketVersioningInput = param2
	f.GetBucketVersioningWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketVersioningWithContextCall.Stub != nil {
		return f.GetBucketVersioningWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketVersioningWithContextCall.Returns.GetBucketVersioningOutput, f.GetBucketVersioningWithContextCall.Returns.Error
}
func (f *S3API) GetBucketWebsite(param1 *s3.GetBucketWebsiteInput) (*s3.GetBucketWebsiteOutput, error) {
	f.GetBucketWebsiteCall.mutex.Lock()
	defer f.GetBucketWebsiteCall.mutex.Unlock()
	f.GetBucketWebsiteCall.CallCount++
	f.GetBucketWebsiteCall.Receives.GetBucketWebsiteInput = param1
	if f.GetBucketWebsiteCall.Stub != nil {
		return f.GetBucketWebsiteCall.Stub(param1)
	}
	return f.GetBucketWebsiteCall.Returns.GetBucketWebsiteOutput, f.GetBucketWebsiteCall.Returns.Error
}
func (f *S3API) GetBucketWebsiteRequest(param1 *s3.GetBucketWebsiteInput) (*request.Request, *s3.GetBucketWebsiteOutput) {
	f.GetBucketWebsiteRequestCall.mutex.Lock()
	defer f.GetBucketWebsiteRequestCall.mutex.Unlock()
	f.GetBucketWebsiteRequestCall.CallCount++
	f.GetBucketWebsiteRequestCall.Receives.GetBucketWebsiteInput = param1
	if f.GetBucketWebsiteRequestCall.Stub != nil {
		return f.GetBucketWebsiteRequestCall.Stub(param1)
	}
	return f.GetBucketWebsiteRequestCall.Returns.Request, f.GetBucketWebsiteRequestCall.Returns.GetBucketWebsiteOutput
}
func (f *S3API) GetBucketWebsiteWithContext(param1 context.Context, param2 *s3.GetBucketWebsiteInput, param3 ...request.Option) (*s3.GetBucketWebsiteOutput, error) {
	f.GetBucketWebsiteWithContextCall.mutex.Lock()
	defer f.GetBucketWebsiteWithContextCall.mutex.Unlock()
	f.GetBucketWebsiteWithContextCall.CallCount++
	f.GetBucketWebsiteWithContextCall.Receives.Context = param1
	f.GetBucketWebsiteWithContextCall.Receives.GetBucketWebsiteInput = param2
	f.GetBucketWebsiteWithContextCall.Receives.OptionSlice = param3
	if f.GetBucketWebsiteWithContextCall.Stub != nil {
		return f.GetBucketWebsiteWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetBucketWebsiteWithContextCall.Returns.GetBucketWebsiteOutput, f.GetBucketWebsiteWithContextCall.Returns.Error
}
func (f *S3API) GetObject(param1 *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	f.GetObjectCall.mutex.Lock()
	defer f.GetObjectCall.mutex.Unlock()
	f.GetObjectCall.CallCount++
	f.GetObjectCall.Receives.GetObjectInput = param1
	if f.GetObjectCall.Stub != nil {
		return f.GetObjectCall.Stub(param1)
	}
	return f.GetObjectCall.Returns.GetObjectOutput, f.GetObjectCall.Returns.Error
}
func (f *S3API) GetObjectAcl(param1 *s3.GetObjectAclInput) (*s3.GetObjectAclOutput, error) {
	f.GetObjectAclCall.mutex.Lock()
	defer f.GetObjectAclCall.mutex.Unlock()
	f.GetObjectAclCall.CallCount++
	f.GetObjectAclCall.Receives.GetObjectAclInput = param1
	if f.GetObjectAclCall.Stub != nil {
		return f.GetObjectAclCall.Stub(param1)
	}
	return f.GetObjectAclCall.Returns.GetObjectAclOutput, f.GetObjectAclCall.Returns.Error
}
func (f *S3API) GetObjectAclRequest(param1 *s3.GetObjectAclInput) (*request.Request, *s3.GetObjectAclOutput) {
	f.GetObjectAclRequestCall.mutex.Lock()
	defer f.GetObjectAclRequestCall.mutex.Unlock()
	f.GetObjectAclRequestCall.CallCount++
	f.GetObjectAclRequestCall.Receives.GetObjectAclInput = param1
	if f.GetObjectAclRequestCall.Stub != nil {
		return f.GetObjectAclRequestCall.Stub(param1)
	}
	return f.GetObjectAclRequestCall.Returns.Request, f.GetObjectAclRequestCall.Returns.GetObjectAclOutput
}
func (f *S3API) GetObjectAclWithContext(param1 context.Context, param2 *s3.GetObjectAclInput, param3 ...request.Option) (*s3.GetObjectAclOutput, error) {
	f.GetObjectAclWithContextCall.mutex.Lock()
	defer f.GetObjectAclWithContextCall.mutex.Unlock()
	f.GetObjectAclWithContextCall.CallCount++
	f.GetObjectAclWithContextCall.Receives.Context = param1
	f.GetObjectAclWithContextCall.Receives.GetObjectAclInput = param2
	f.GetObjectAclWithContextCall.Receives.OptionSlice = param3
	if f.GetObjectAclWithContextCall.Stub != nil {
		return f.GetObjectAclWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetObjectAclWithContextCall.Returns.GetObjectAclOutput, f.GetObjectAclWithContextCall.Returns.Error
}
func (f *S3API) GetObjectAttributes(param1 *s3.GetObjectAttributesInput) (*s3.GetObjectAttributesOutput, error) {
	f.GetObjectAttributesCall.mutex.Lock()
	defer f.GetObjectAttributesCall.mutex.Unlock()
	f.GetObjectAttributesCall.CallCount++
	f.GetObjectAttributesCall.Receives.GetObjectAttributesInput = param1
	if f.GetObjectAttributesCall.Stub != nil {
		return f.GetObjectAttributesCall.Stub(param1)
	}
	return f.GetObjectAttributesCall.Returns.GetObjectAttributesOutput, f.GetObjectAttributesCall.Returns.Error
}
func (f *S3API) GetObjectAttributesRequest(param1 *s3.GetObjectAttributesInput) (*request.Request, *s3.GetObjectAttributesOutput) {
	f.GetObjectAttributesRequestCall.mutex.Lock()
	defer f.GetObjectAttributesRequestCall.mutex.Unlock()
	f.GetObjectAttributesRequestCall.CallCount++
	f.GetObjectAttributesRequestCall.Receives.GetObjectAttributesInput = param1
	if f.GetObjectAttributesRequestCall.Stub != nil {
		return f.GetObjectAttributesRequestCall.Stub(param1)
	}
	return f.GetObjectAttributesRequestCall.Returns.Request, f.GetObjectAttributesRequestCall.Returns.GetObjectAttributesOutput
}
func (f *S3API) GetObjectAttributesWithContext(param1 context.Context, param2 *s3.GetObjectAttributesInput, param3 ...request.Option) (*s3.GetObjectAttributesOutput, error) {
	f.GetObjectAttributesWithContextCall.mutex.Lock()
	defer f.GetObjectAttributesWithContextCall.mutex.Unlock()
	f.GetObjectAttributesWithContextCall.CallCount++
	f.GetObjectAttributesWithContextCall.Receives.Context = param1
	f.GetObjectAttributesWithContextCall.Receives.GetObjectAttributesInput = param2
	f.GetObjectAttributesWithContextCall.Receives.OptionSlice = param3
	if f.GetObjectAttributesWithContextCall.Stub != nil {
		return f.GetObjectAttributesWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetObjectAttributesWithContextCall.Returns.GetObjectAttributesOutput, f.GetObjectAttributesWithContextCall.Returns.Error
}
func (f *S3API) GetObjectLegalHold(param1 *s3.GetObjectLegalHoldInput) (*s3.GetObjectLegalHoldOutput, error) {
	f.GetObjectLegalHoldCall.mutex.Lock()
	defer f.GetObjectLegalHoldCall.mutex.Unlock()
	f.GetObjectLegalHoldCall.CallCount++
	f.GetObjectLegalHoldCall.Receives.GetObjectLegalHoldInput = param1
	if f.GetObjectLegalHoldCall.Stub != nil {
		return f.GetObjectLegalHoldCall.Stub(param1)
	}
	return f.GetObjectLegalHoldCall.Returns.GetObjectLegalHoldOutput, f.GetObjectLegalHoldCall.Returns.Error
}
func (f *S3API) GetObjectLegalHoldRequest(param1 *s3.GetObjectLegalHoldInput) (*request.Request, *s3.GetObjectLegalHoldOutput) {
	f.GetObjectLegalHoldRequestCall.mutex.Lock()
	defer f.GetObjectLegalHoldRequestCall.mutex.Unlock()
	f.GetObjectLegalHoldRequestCall.CallCount++
	f.GetObjectLegalHoldRequestCall.Receives.GetObjectLegalHoldInput = param1
	if f.GetObjectLegalHoldRequestCall.Stub != nil {
		return f.GetObjectLegalHoldRequestCall.Stub(param1)
	}
	return f.GetObjectLegalHoldRequestCall.Returns.Request, f.GetObjectLegalHoldRequestCall.Returns.GetObjectLegalHoldOutput
}
func (f *S3API) GetObjectLegalHoldWithContext(param1 context.Context, param2 *s3.GetObjectLegalHoldInput, param3 ...request.Option) (*s3.GetObjectLegalHoldOutput, error) {
	f.GetObjectLegalHoldWithContextCall.mutex.Lock()
	defer f.GetObjectLegalHoldWithContextCall.mutex.Unlock()
	f.GetObjectLegalHoldWithContextCall.CallCount++
	f.GetObjectLegalHoldWithContextCall.Receives.Context = param1
	f.GetObjectLegalHoldWithContextCall.Receives.GetObjectLegalHoldInput = param2
	f.GetObjectLegalHoldWithContextCall.Receives.OptionSlice = param3
	if f.GetObjectLegalHoldWithContextCall.Stub != nil {
		return f.GetObjectLegalHoldWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetObjectLegalHoldWithContextCall.Returns.GetObjectLegalHoldOutput, f.GetObjectLegalHoldWithContextCall.Returns.Error
}
func (f *S3API) GetObjectLockConfiguration(param1 *s3.GetObjectLockConfigurationInput) (*s3.GetObjectLockConfigurationOutput, error) {
	f.GetObjectLockConfigurationCall.mutex.Lock()
	defer f.GetObjectLockConfigurationCall.mutex.Unlock()
	f.GetObjectLockConfigurationCall.CallCount++
	f.GetObjectLockConfigurationCall.Receives.GetObjectLockConfigurationInput = param1
	if f.GetObjectLockConfigurationCall.Stub != nil {
		return f.GetObjectLockConfigurationCall.Stub(param1)
	}
	return f.GetObjectLockConfigurationCall.Returns.GetObjectLockConfigurationOutput, f.GetObjectLockConfigurationCall.Returns.Error
}
func (f *S3API) GetObjectLockConfigurationRequest(param1 *s3.GetObjectLockConfigurationInput) (*request.Request, *s3.GetObjectLockConfigurationOutput) {
	f.GetObjectLockConfigurationRequestCall.mutex.Lock()
	defer f.GetObjectLockConfigurationRequestCall.mutex.Unlock()
	f.GetObjectLockConfigurationRequestCall.CallCount++
	f.GetObjectLockConfigurationRequestCall.Receives.GetObjectLockConfigurationInput = param1
	if f.GetObjectLockConfigurationRequestCall.Stub != nil {
		return f.GetObjectLockConfigurationRequestCall.Stub(param1)
	}
	return f.GetObjectLockConfigurationRequestCall.Returns.Request, f.GetObjectLockConfigurationRequestCall.Returns.GetObjectLockConfigurationOutput
}
func (f *S3API) GetObjectLockConfigurationWithContext(param1 context.Context, param2 *s3.GetObjectLockConfigurationInput, param3 ...request.Option) (*s3.GetObjectLockConfigurationOutput, error) {
	f.GetObjectLockConfigurationWithContextCall.mutex.Lock()
	defer f.GetObjectLockConfigurationWithContextCall.mutex.Unlock()
	f.GetObjectLockConfigurationWithContextCall.CallCount++
	f.GetObjectLockConfigurationWithContextCall.Receives.Context = param1
	f.GetObjectLockConfigurationWithContextCall.Receives.GetObjectLockConfigurationInput = param2
	f.GetObjectLockConfigurationWithContextCall.Receives.OptionSlice = param3
	if f.GetObjectLockConfigurationWithContextCall.Stub != nil {
		return f.GetObjectLockConfigurationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetObjectLockConfigurationWithContextCall.Returns.GetObjectLockConfigurationOutput, f.GetObjectLockConfigurationWithContextCall.Returns.Error
}
func (f *S3API) GetObjectRequest(param1 *s3.GetObjectInput) (*request.Request, *s3.GetObjectOutput) {
	f.GetObjectRequestCall.mutex.Lock()
	defer f.GetObjectRequestCall.mutex.Unlock()
	f.GetObjectRequestCall.CallCount++
	f.GetObjectRequestCall.Receives.GetObjectInput = param1
	if f.GetObjectRequestCall.Stub != nil {
		return f.GetObjectRequestCall.Stub(param1)
	}
	return f.GetObjectRequestCall.Returns.Request, f.GetObjectRequestCall.Returns.GetObjectOutput
}
func (f *S3API) GetObjectRetention(param1 *s3.GetObjectRetentionInput) (*s3.GetObjectRetentionOutput, error) {
	f.GetObjectRetentionCall.mutex.Lock()
	defer f.GetObjectRetentionCall.mutex.Unlock()
	f.GetObjectRetentionCall.CallCount++
	f.GetObjectRetentionCall.Receives.GetObjectRetentionInput = param1
	if f.GetObjectRetentionCall.Stub != nil {
		return f.GetObjectRetentionCall.Stub(param1)
	}
	return f.GetObjectRetentionCall.Returns.GetObjectRetentionOutput, f.GetObjectRetentionCall.Returns.Error
}
func (f *S3API) GetObjectRetentionRequest(param1 *s3.GetObjectRetentionInput) (*request.Request, *s3.GetObjectRetentionOutput) {
	f.GetObjectRetentionRequestCall.mutex.Lock()
	defer f.GetObjectRetentionRequestCall.mutex.Unlock()
	f.GetObjectRetentionRequestCall.CallCount++
	f.GetObjectRetentionRequestCall.Receives.GetObjectRetentionInput = param1
	if f.GetObjectRetentionRequestCall.Stub != nil {
		return f.GetObjectRetentionRequestCall.Stub(param1)
	}
	return f.GetObjectRetentionRequestCall.Returns.Request, f.GetObjectRetentionRequestCall.Returns.GetObjectRetentionOutput
}
func (f *S3API) GetObjectRetentionWithContext(param1 context.Context, param2 *s3.GetObjectRetentionInput, param3 ...request.Option) (*s3.GetObjectRetentionOutput, error) {
	f.GetObjectRetentionWithContextCall.mutex.Lock()
	defer f.GetObjectRetentionWithContextCall.mutex.Unlock()
	f.GetObjectRetentionWithContextCall.CallCount++
	f.GetObjectRetentionWithContextCall.Receives.Context = param1
	f.GetObjectRetentionWithContextCall.Receives.GetObjectRetentionInput = param2
	f.GetObjectRetentionWithContextCall.Receives.OptionSlice = param3
	if f.GetObjectRetentionWithContextCall.Stub != nil {
		return f.GetObjectRetentionWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetObjectRetentionWithContextCall.Returns.GetObjectRetentionOutput, f.GetObjectRetentionWithContextCall.Returns.Error
}
func (f *S3API) GetObjectTagging(param1 *s3.GetObjectTaggingInput) (*s3.GetObjectTaggingOutput, error) {
	f.GetObjectTaggingCall.mutex.Lock()
	defer f.GetObjectTaggingCall.mutex.Unlock()
	f.GetObjectTaggingCall.CallCount++
	f.GetObjectTaggingCall.Receives.GetObjectTaggingInput = param1
	if f.GetObjectTaggingCall.Stub != nil {
		return f.GetObjectTaggingCall.Stub(param1)
	}
	return f.GetObjectTaggingCall.Returns.GetObjectTaggingOutput, f.GetObjectTaggingCall.Returns.Error
}
func (f *S3API) GetObjectTaggingRequest(param1 *s3.GetObjectTaggingInput) (*request.Request, *s3.GetObjectTaggingOutput) {
	f.GetObjectTaggingRequestCall.mutex.Lock()
	defer f.GetObjectTaggingRequestCall.mutex.Unlock()
	f.GetObjectTaggingRequestCall.CallCount++
	f.GetObjectTaggingRequestCall.Receives.GetObjectTaggingInput = param1
	if f.GetObjectTaggingRequestCall.Stub != nil {
		return f.GetObjectTaggingRequestCall.Stub(param1)
	}
	return f.GetObjectTaggingRequestCall.Returns.Request, f.GetObjectTaggingRequestCall.Returns.GetObjectTaggingOutput
}
func (f *S3API) GetObjectTaggingWithContext(param1 context.Context, param2 *s3.GetObjectTaggingInput, param3 ...request.Option) (*s3.GetObjectTaggingOutput, error) {
	f.GetObjectTaggingWithContextCall.mutex.Lock()
	defer f.GetObjectTaggingWithContextCall.mutex.Unlock()
	f.GetObjectTaggingWithContextCall.CallCount++
	f.GetObjectTaggingWithContextCall.Receives.Context = param1
	f.GetObjectTaggingWithContextCall.Receives.GetObjectTaggingInput = param2
	f.GetObjectTaggingWithContextCall.Receives.OptionSlice = param3
	if f.GetObjectTaggingWithContextCall.Stub != nil {
		return f.GetObjectTaggingWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetObjectTaggingWithContextCall.Returns.GetObjectTaggingOutput, f.GetObjectTaggingWithContextCall.Returns.Error
}
func (f *S3API) GetObjectTorrent(param1 *s3.GetObjectTorrentInput) (*s3.GetObjectTorrentOutput, error) {
	f.GetObjectTorrentCall.mutex.Lock()
	defer f.GetObjectTorrentCall.mutex.Unlock()
	f.GetObjectTorrentCall.CallCount++
	f.GetObjectTorrentCall.Receives.GetObjectTorrentInput = param1
	if f.GetObjectTorrentCall.Stub != nil {
		return f.GetObjectTorrentCall.Stub(param1)
	}
	return f.GetObjectTorrentCall.Returns.GetObjectTorrentOutput, f.GetObjectTorrentCall.Returns.Error
}
func (f *S3API) GetObjectTorrentRequest(param1 *s3.GetObjectTorrentInput) (*request.Request, *s3.GetObjectTorrentOutput) {
	f.GetObjectTorrentRequestCall.mutex.Lock()
	defer f.GetObjectTorrentRequestCall.mutex.Unlock()
	f.GetObjectTorrentRequestCall.CallCount++
	f.GetObjectTorrentRequestCall.Receives.GetObjectTorrentInput = param1
	if f.GetObjectTorrentRequestCall.Stub != nil {
		return f.GetObjectTorrentRequestCall.Stub(param1)
	}
	return f.GetObjectTorrentRequestCall.Returns.Request, f.GetObjectTorrentRequestCall.Returns.GetObjectTorrentOutput
}
func (f *S3API) GetObjectTorrentWithContext(param1 context.Context, param2 *s3.GetObjectTorrentInput, param3 ...request.Option) (*s3.GetObjectTorrentOutput, error) {
	f.GetObjectTorrentWithContextCall.mutex.Lock()
	defer f.GetObjectTorrentWithContextCall.mutex.Unlock()
	f.GetObjectTorrentWithContextCall.CallCount++
	f.GetObjectTorrentWithContextCall.Receives.Context = param1
	f.GetObjectTorrentWithContextCall.Receives.GetObjectTorrentInput = param2
	f.GetObjectTorrentWithContextCall.Receives.OptionSlice = param3
	if f.GetObjectTorrentWithContextCall.Stub != nil {
		return f.GetObjectTorrentWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetObjectTorrentWithContextCall.Returns.GetObjectTorrentOutput, f.GetObjectTorrentWithContextCall.Returns.Error
}
func (f *S3API) GetObjectWithContext(param1 context.Context, param2 *s3.GetObjectInput, param3 ...request.Option) (*s3.GetObjectOutput, error) {
	f.GetObjectWithContextCall.mutex.Lock()
	defer f.GetObjectWithContextCall.mutex.Unlock()
	f.GetObjectWithContextCall.CallCount++
	f.GetObjectWithContextCall.Receives.Context = param1
	f.GetObjectWithContextCall.Receives.GetObjectInput = param2
	f.GetObjectWithContextCall.Receives.OptionSlice = param3
	if f.GetObjectWithContextCall.Stub != nil {
		return f.GetObjectWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetObjectWithContextCall.Returns.GetObjectOutput, f.GetObjectWithContextCall.Returns.Error
}
func (f *S3API) GetPublicAccessBlock(param1 *s3.GetPublicAccessBlockInput) (*s3.GetPublicAccessBlockOutput, error) {
	f.GetPublicAccessBlockCall.mutex.Lock()
	defer f.GetPublicAccessBlockCall.mutex.Unlock()
	f.GetPublicAccessBlockCall.CallCount++
	f.GetPublicAccessBlockCall.Receives.GetPublicAccessBlockInput = param1
	if f.GetPublicAccessBlockCall.Stub != nil {
		return f.GetPublicAccessBlockCall.Stub(param1)
	}
	return f.GetPublicAccessBlockCall.Returns.GetPublicAccessBlockOutput, f.GetPublicAccessBlockCall.Returns.Error
}
func (f *S3API) GetPublicAccessBlockRequest(param1 *s3.GetPublicAccessBlockInput) (*request.Request, *s3.GetPublicAccessBlockOutput) {
	f.GetPublicAccessBlockRequestCall.mutex.Lock()
	defer f.GetPublicAccessBlockRequestCall.mutex.Unlock()
	f.GetPublicAccessBlockRequestCall.CallCount++
	f.GetPublicAccessBlockRequestCall.Receives.GetPublicAccessBlockInput = param1
	if f.GetPublicAccessBlockRequestCall.Stub != nil {
		return f.GetPublicAccessBlockRequestCall.Stub(param1)
	}
	return f.GetPublicAccessBlockRequestCall.Returns.Request, f.GetPublicAccessBlockRequestCall.Returns.GetPublicAccessBlockOutput
}
func (f *S3API) GetPublicAccessBlockWithContext(param1 context.Context, param2 *s3.GetPublicAccessBlockInput, param3 ...request.Option) (*s3.GetPublicAccessBlockOutput, error) {
	f.GetPublicAccessBlockWithContextCall.mutex.Lock()
	defer f.GetPublicAccessBlockWithContextCall.mutex.Unlock()
	f.GetPublicAccessBlockWithContextCall.CallCount++
	f.GetPublicAccessBlockWithContextCall.Receives.Context = param1
	f.GetPublicAccessBlockWithContextCall.Receives.GetPublicAccessBlockInput = param2
	f.GetPublicAccessBlockWithContextCall.Receives.OptionSlice = param3
	if f.GetPublicAccessBlockWithContextCall.Stub != nil {
		return f.GetPublicAccessBlockWithContextCall.Stub(param1, param2, param3...)
	}
	return f.GetPublicAccessBlockWithContextCall.Returns.GetPublicAccessBlockOutput, f.GetPublicAccessBlockWithContextCall.Returns.Error
}
func (f *S3API) HeadBucket(param1 *s3.HeadBucketInput) (*s3.HeadBucketOutput, error) {
	f.HeadBucketCall.mutex.Lock()
	defer f.HeadBucketCall.mutex.Unlock()
	f.HeadBucketCall.CallCount++
	f.HeadBucketCall.Receives.HeadBucketInput = param1
	if f.HeadBucketCall.Stub != nil {
		return f.HeadBucketCall.Stub(param1)
	}
	return f.HeadBucketCall.Returns.HeadBucketOutput, f.HeadBucketCall.Returns.Error
}
func (f *S3API) HeadBucketRequest(param1 *s3.HeadBucketInput) (*request.Request, *s3.HeadBucketOutput) {
	f.HeadBucketRequestCall.mutex.Lock()
	defer f.HeadBucketRequestCall.mutex.Unlock()
	f.HeadBucketRequestCall.CallCount++
	f.HeadBucketRequestCall.Receives.HeadBucketInput = param1
	if f.HeadBucketRequestCall.Stub != nil {
		return f.HeadBucketRequestCall.Stub(param1)
	}
	return f.HeadBucketRequestCall.Returns.Request, f.HeadBucketRequestCall.Returns.HeadBucketOutput
}
func (f *S3API) HeadBucketWithContext(param1 context.Context, param2 *s3.HeadBucketInput, param3 ...request.Option) (*s3.HeadBucketOutput, error) {
	f.HeadBucketWithContextCall.mutex.Lock()
	defer f.HeadBucketWithContextCall.mutex.Unlock()
	f.HeadBucketWithContextCall.CallCount++
	f.HeadBucketWithContextCall.Receives.Context = param1
	f.HeadBucketWithContextCall.Receives.HeadBucketInput = param2
	f.HeadBucketWithContextCall.Receives.OptionSlice = param3
	if f.HeadBucketWithContextCall.Stub != nil {
		return f.HeadBucketWithContextCall.Stub(param1, param2, param3...)
	}
	return f.HeadBucketWithContextCall.Returns.HeadBucketOutput, f.HeadBucketWithContextCall.Returns.Error
}
func (f *S3API) HeadObject(param1 *s3.HeadObjectInput) (*s3.HeadObjectOutput, error) {
	f.HeadObjectCall.mutex.Lock()
	defer f.HeadObjectCall.mutex.Unlock()
	f.HeadObjectCall.CallCount++
	f.HeadObjectCall.Receives.HeadObjectInput = param1
	if f.HeadObjectCall.Stub != nil {
		return f.HeadObjectCall.Stub(param1)
	}
	return f.HeadObjectCall.Returns.HeadObjectOutput, f.HeadObjectCall.Returns.Error
}
func (f *S3API) HeadObjectRequest(param1 *s3.HeadObjectInput) (*request.Request, *s3.HeadObjectOutput) {
	f.HeadObjectRequestCall.mutex.Lock()
	defer f.HeadObjectRequestCall.mutex.Unlock()
	f.HeadObjectRequestCall.CallCount++
	f.HeadObjectRequestCall.Receives.HeadObjectInput = param1
	if f.HeadObjectRequestCall.Stub != nil {
		return f.HeadObjectRequestCall.Stub(param1)
	}
	return f.HeadObjectRequestCall.Returns.Request, f.HeadObjectRequestCall.Returns.HeadObjectOutput
}
func (f *S3API) HeadObjectWithContext(param1 context.Context, param2 *s3.HeadObjectInput, param3 ...request.Option) (*s3.HeadObjectOutput, error) {
	f.HeadObjectWithContextCall.mutex.Lock()
	defer f.HeadObjectWithContextCall.mutex.Unlock()
	f.HeadObjectWithContextCall.CallCount++
	f.HeadObjectWithContextCall.Receives.Context = param1
	f.HeadObjectWithContextCall.Receives.HeadObjectInput = param2
	f.HeadObjectWithContextCall.Receives.OptionSlice = param3
	if f.HeadObjectWithContextCall.Stub != nil {
		return f.HeadObjectWithContextCall.Stub(param1, param2, param3...)
	}
	return f.HeadObjectWithContextCall.Returns.HeadObjectOutput, f.HeadObjectWithContextCall.Returns.Error
}
func (f *S3API) ListBucketAnalyticsConfigurations(param1 *s3.ListBucketAnalyticsConfigurationsInput) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	f.ListBucketAnalyticsConfigurationsCall.mutex.Lock()
	defer f.ListBucketAnalyticsConfigurationsCall.mutex.Unlock()
	f.ListBucketAnalyticsConfigurationsCall.CallCount++
	f.ListBucketAnalyticsConfigurationsCall.Receives.ListBucketAnalyticsConfigurationsInput = param1
	if f.ListBucketAnalyticsConfigurationsCall.Stub != nil {
		return f.ListBucketAnalyticsConfigurationsCall.Stub(param1)
	}
	return f.ListBucketAnalyticsConfigurationsCall.Returns.ListBucketAnalyticsConfigurationsOutput, f.ListBucketAnalyticsConfigurationsCall.Returns.Error
}
func (f *S3API) ListBucketAnalyticsConfigurationsRequest(param1 *s3.ListBucketAnalyticsConfigurationsInput) (*request.Request, *s3.ListBucketAnalyticsConfigurationsOutput) {
	f.ListBucketAnalyticsConfigurationsRequestCall.mutex.Lock()
	defer f.ListBucketAnalyticsConfigurationsRequestCall.mutex.Unlock()
	f.ListBucketAnalyticsConfigurationsRequestCall.CallCount++
	f.ListBucketAnalyticsConfigurationsRequestCall.Receives.ListBucketAnalyticsConfigurationsInput = param1
	if f.ListBucketAnalyticsConfigurationsRequestCall.Stub != nil {
		return f.ListBucketAnalyticsConfigurationsRequestCall.Stub(param1)
	}
	return f.ListBucketAnalyticsConfigurationsRequestCall.Returns.Request, f.ListBucketAnalyticsConfigurationsRequestCall.Returns.ListBucketAnalyticsConfigurationsOutput
}
func (f *S3API) ListBucketAnalyticsConfigurationsWithContext(param1 context.Context, param2 *s3.ListBucketAnalyticsConfigurationsInput, param3 ...request.Option) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	f.ListBucketAnalyticsConfigurationsWithContextCall.mutex.Lock()
	defer f.ListBucketAnalyticsConfigurationsWithContextCall.mutex.Unlock()
	f.ListBucketAnalyticsConfigurationsWithContextCall.CallCount++
	f.ListBucketAnalyticsConfigurationsWithContextCall.Receives.Context = param1
	f.ListBucketAnalyticsConfigurationsWithContextCall.Receives.ListBucketAnalyticsConfigurationsInput = param2
	f.ListBucketAnalyticsConfigurationsWithContextCall.Receives.OptionSlice = param3
	if f.ListBucketAnalyticsConfigurationsWithContextCall.Stub != nil {
		return f.ListBucketAnalyticsConfigurationsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.ListBucketAnalyticsConfigurationsWithContextCall.Returns.ListBucketAnalyticsConfigurationsOutput, f.ListBucketAnalyticsConfigurationsWithContextCall.Returns.Error
}
func (f *S3API) ListBucketIntelligentTieringConfigurations(param1 *s3.ListBucketIntelligentTieringConfigurationsInput) (*s3.ListBucketIntelligentTieringConfigurationsOutput, error) {
	f.ListBucketIntelligentTieringConfigurationsCall.mutex.Lock()
	defer f.ListBucketIntelligentTieringConfigurationsCall.mutex.Unlock()
	f.ListBucketIntelligentTieringConfigurationsCall.CallCount++
	f.ListBucketIntelligentTieringConfigurationsCall.Receives.ListBucketIntelligentTieringConfigurationsInput = param1
	if f.ListBucketIntelligentTieringConfigurationsCall.Stub != nil {
		return f.ListBucketIntelligentTieringConfigurationsCall.Stub(param1)
	}
	return f.ListBucketIntelligentTieringConfigurationsCall.Returns.ListBucketIntelligentTieringConfigurationsOutput, f.ListBucketIntelligentTieringConfigurationsCall.Returns.Error
}
func (f *S3API) ListBucketIntelligentTieringConfigurationsRequest(param1 *s3.ListBucketIntelligentTieringConfigurationsInput) (*request.Request, *s3.ListBucketIntelligentTieringConfigurationsOutput) {
	f.ListBucketIntelligentTieringConfigurationsRequestCall.mutex.Lock()
	defer f.ListBucketIntelligentTieringConfigurationsRequestCall.mutex.Unlock()
	f.ListBucketIntelligentTieringConfigurationsRequestCall.CallCount++
	f.ListBucketIntelligentTieringConfigurationsRequestCall.Receives.ListBucketIntelligentTieringConfigurationsInput = param1
	if f.ListBucketIntelligentTieringConfigurationsRequestCall.Stub != nil {
		return f.ListBucketIntelligentTieringConfigurationsRequestCall.Stub(param1)
	}
	return f.ListBucketIntelligentTieringConfigurationsRequestCall.Returns.Request, f.ListBucketIntelligentTieringConfigurationsRequestCall.Returns.ListBucketIntelligentTieringConfigurationsOutput
}
func (f *S3API) ListBucketIntelligentTieringConfigurationsWithContext(param1 context.Context, param2 *s3.ListBucketIntelligentTieringConfigurationsInput, param3 ...request.Option) (*s3.ListBucketIntelligentTieringConfigurationsOutput, error) {
	f.ListBucketIntelligentTieringConfigurationsWithContextCall.mutex.Lock()
	defer f.ListBucketIntelligentTieringConfigurationsWithContextCall.mutex.Unlock()
	f.ListBucketIntelligentTieringConfigurationsWithContextCall.CallCount++
	f.ListBucketIntelligentTieringConfigurationsWithContextCall.Receives.Context = param1
	f.ListBucketIntelligentTieringConfigurationsWithContextCall.Receives.ListBucketIntelligentTieringConfigurationsInput = param2
	f.ListBucketIntelligentTieringConfigurationsWithContextCall.Receives.OptionSlice = param3
	if f.ListBucketIntelligentTieringConfigurationsWithContextCall.Stub != nil {
		return f.ListBucketIntelligentTieringConfigurationsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.ListBucketIntelligentTieringConfigurationsWithContextCall.Returns.ListBucketIntelligentTieringConfigurationsOutput, f.ListBucketIntelligentTieringConfigurationsWithContextCall.Returns.Error
}
func (f *S3API) ListBucketInventoryConfigurations(param1 *s3.ListBucketInventoryConfigurationsInput) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	f.ListBucketInventoryConfigurationsCall.mutex.Lock()
	defer f.ListBucketInventoryConfigurationsCall.mutex.Unlock()
	f.ListBucketInventoryConfigurationsCall.CallCount++
	f.ListBucketInventoryConfigurationsCall.Receives.ListBucketInventoryConfigurationsInput = param1
	if f.ListBucketInventoryConfigurationsCall.Stub != nil {
		return f.ListBucketInventoryConfigurationsCall.Stub(param1)
	}
	return f.ListBucketInventoryConfigurationsCall.Returns.ListBucketInventoryConfigurationsOutput, f.ListBucketInventoryConfigurationsCall.Returns.Error
}
func (f *S3API) ListBucketInventoryConfigurationsRequest(param1 *s3.ListBucketInventoryConfigurationsInput) (*request.Request, *s3.ListBucketInventoryConfigurationsOutput) {
	f.ListBucketInventoryConfigurationsRequestCall.mutex.Lock()
	defer f.ListBucketInventoryConfigurationsRequestCall.mutex.Unlock()
	f.ListBucketInventoryConfigurationsRequestCall.CallCount++
	f.ListBucketInventoryConfigurationsRequestCall.Receives.ListBucketInventoryConfigurationsInput = param1
	if f.ListBucketInventoryConfigurationsRequestCall.Stub != nil {
		return f.ListBucketInventoryConfigurationsRequestCall.Stub(param1)
	}
	return f.ListBucketInventoryConfigurationsRequestCall.Returns.Request, f.ListBucketInventoryConfigurationsRequestCall.Returns.ListBucketInventoryConfigurationsOutput
}
func (f *S3API) ListBucketInventoryConfigurationsWithContext(param1 context.Context, param2 *s3.ListBucketInventoryConfigurationsInput, param3 ...request.Option) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	f.ListBucketInventoryConfigurationsWithContextCall.mutex.Lock()
	defer f.ListBucketInventoryConfigurationsWithContextCall.mutex.Unlock()
	f.ListBucketInventoryConfigurationsWithContextCall.CallCount++
	f.ListBucketInventoryConfigurationsWithContextCall.Receives.Context = param1
	f.ListBucketInventoryConfigurationsWithContextCall.Receives.ListBucketInventoryConfigurationsInput = param2
	f.ListBucketInventoryConfigurationsWithContextCall.Receives.OptionSlice = param3
	if f.ListBucketInventoryConfigurationsWithContextCall.Stub != nil {
		return f.ListBucketInventoryConfigurationsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.ListBucketInventoryConfigurationsWithContextCall.Returns.ListBucketInventoryConfigurationsOutput, f.ListBucketInventoryConfigurationsWithContextCall.Returns.Error
}
func (f *S3API) ListBucketMetricsConfigurations(param1 *s3.ListBucketMetricsConfigurationsInput) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	f.ListBucketMetricsConfigurationsCall.mutex.Lock()
	defer f.ListBucketMetricsConfigurationsCall.mutex.Unlock()
	f.ListBucketMetricsConfigurationsCall.CallCount++
	f.ListBucketMetricsConfigurationsCall.Receives.ListBucketMetricsConfigurationsInput = param1
	if f.ListBucketMetricsConfigurationsCall.Stub != nil {
		return f.ListBucketMetricsConfigurationsCall.Stub(param1)
	}
	return f.ListBucketMetricsConfigurationsCall.Returns.ListBucketMetricsConfigurationsOutput, f.ListBucketMetricsConfigurationsCall.Returns.Error
}
func (f *S3API) ListBucketMetricsConfigurationsRequest(param1 *s3.ListBucketMetricsConfigurationsInput) (*request.Request, *s3.ListBucketMetricsConfigurationsOutput) {
	f.ListBucketMetricsConfigurationsRequestCall.mutex.Lock()
	defer f.ListBucketMetricsConfigurationsRequestCall.mutex.Unlock()
	f.ListBucketMetricsConfigurationsRequestCall.CallCount++
	f.ListBucketMetricsConfigurationsRequestCall.Receives.ListBucketMetricsConfigurationsInput = param1
	if f.ListBucketMetricsConfigurationsRequestCall.Stub != nil {
		return f.ListBucketMetricsConfigurationsRequestCall.Stub(param1)
	}
	return f.ListBucketMetricsConfigurationsRequestCall.Returns.Request, f.ListBucketMetricsConfigurationsRequestCall.Returns.ListBucketMetricsConfigurationsOutput
}
func (f *S3API) ListBucketMetricsConfigurationsWithContext(param1 context.Context, param2 *s3.ListBucketMetricsConfigurationsInput, param3 ...request.Option) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	f.ListBucketMetricsConfigurationsWithContextCall.mutex.Lock()
	defer f.ListBucketMetricsConfigurationsWithContextCall.mutex.Unlock()
	f.ListBucketMetricsConfigurationsWithContextCall.CallCount++
	f.ListBucketMetricsConfigurationsWithContextCall.Receives.Context = param1
	f.ListBucketMetricsConfigurationsWithContextCall.Receives.ListBucketMetricsConfigurationsInput = param2
	f.ListBucketMetricsConfigurationsWithContextCall.Receives.OptionSlice = param3
	if f.ListBucketMetricsConfigurationsWithContextCall.Stub != nil {
		return f.ListBucketMetricsConfigurationsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.ListBucketMetricsConfigurationsWithContextCall.Returns.ListBucketMetricsConfigurationsOutput, f.ListBucketMetricsConfigurationsWithContextCall.Returns.Error
}
func (f *S3API) ListBuckets(param1 *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	f.ListBucketsCall.mutex.Lock()
	defer f.ListBucketsCall.mutex.Unlock()
	f.ListBucketsCall.CallCount++
	f.ListBucketsCall.Receives.ListBucketsInput = param1
	if f.ListBucketsCall.Stub != nil {
		return f.ListBucketsCall.Stub(param1)
	}
	return f.ListBucketsCall.Returns.ListBucketsOutput, f.ListBucketsCall.Returns.Error
}
func (f *S3API) ListBucketsRequest(param1 *s3.ListBucketsInput) (*request.Request, *s3.ListBucketsOutput) {
	f.ListBucketsRequestCall.mutex.Lock()
	defer f.ListBucketsRequestCall.mutex.Unlock()
	f.ListBucketsRequestCall.CallCount++
	f.ListBucketsRequestCall.Receives.ListBucketsInput = param1
	if f.ListBucketsRequestCall.Stub != nil {
		return f.ListBucketsRequestCall.Stub(param1)
	}
	return f.ListBucketsRequestCall.Returns.Request, f.ListBucketsRequestCall.Returns.ListBucketsOutput
}
func (f *S3API) ListBucketsWithContext(param1 context.Context, param2 *s3.ListBucketsInput, param3 ...request.Option) (*s3.ListBucketsOutput, error) {
	f.ListBucketsWithContextCall.mutex.Lock()
	defer f.ListBucketsWithContextCall.mutex.Unlock()
	f.ListBucketsWithContextCall.CallCount++
	f.ListBucketsWithContextCall.Receives.Context = param1
	f.ListBucketsWithContextCall.Receives.ListBucketsInput = param2
	f.ListBucketsWithContextCall.Receives.OptionSlice = param3
	if f.ListBucketsWithContextCall.Stub != nil {
		return f.ListBucketsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.ListBucketsWithContextCall.Returns.ListBucketsOutput, f.ListBucketsWithContextCall.Returns.Error
}
func (f *S3API) ListMultipartUploads(param1 *s3.ListMultipartUploadsInput) (*s3.ListMultipartUploadsOutput, error) {
	f.ListMultipartUploadsCall.mutex.Lock()
	defer f.ListMultipartUploadsCall.mutex.Unlock()
	f.ListMultipartUploadsCall.CallCount++
	f.ListMultipartUploadsCall.Receives.ListMultipartUploadsInput = param1
	if f.ListMultipartUploadsCall.Stub != nil {
		return f.ListMultipartUploadsCall.Stub(param1)
	}
	return f.ListMultipartUploadsCall.Returns.ListMultipartUploadsOutput, f.ListMultipartUploadsCall.Returns.Error
}
func (f *S3API) ListMultipartUploadsPages(param1 *s3.ListMultipartUploadsInput, param2 func(*s3.ListMultipartUploadsOutput, bool) bool) error {
	f.ListMultipartUploadsPagesCall.mutex.Lock()
	defer f.ListMultipartUploadsPagesCall.mutex.Unlock()
	f.ListMultipartUploadsPagesCall.CallCount++
	f.ListMultipartUploadsPagesCall.Receives.ListMultipartUploadsInput = param1
	f.ListMultipartUploadsPagesCall.Receives.FuncS3ListMultipartUploadsOutputBoolBool = param2
	if f.ListMultipartUploadsPagesCall.Stub != nil {
		return f.ListMultipartUploadsPagesCall.Stub(param1, param2)
	}
	return f.ListMultipartUploadsPagesCall.Returns.Error
}
func (f *S3API) ListMultipartUploadsPagesWithContext(param1 context.Context, param2 *s3.ListMultipartUploadsInput, param3 func(*s3.ListMultipartUploadsOutput, bool) bool, param4 ...request.Option) error {
	f.ListMultipartUploadsPagesWithContextCall.mutex.Lock()
	defer f.ListMultipartUploadsPagesWithContextCall.mutex.Unlock()
	f.ListMultipartUploadsPagesWithContextCall.CallCount++
	f.ListMultipartUploadsPagesWithContextCall.Receives.Context = param1
	f.ListMultipartUploadsPagesWithContextCall.Receives.ListMultipartUploadsInput = param2
	f.ListMultipartUploadsPagesWithContextCall.Receives.FuncS3ListMultipartUploadsOutputBoolBool = param3
	f.ListMultipartUploadsPagesWithContextCall.Receives.OptionSlice = param4
	if f.ListMultipartUploadsPagesWithContextCall.Stub != nil {
		return f.ListMultipartUploadsPagesWithContextCall.Stub(param1, param2, param3, param4...)
	}
	return f.ListMultipartUploadsPagesWithContextCall.Returns.Error
}
func (f *S3API) ListMultipartUploadsRequest(param1 *s3.ListMultipartUploadsInput) (*request.Request, *s3.ListMultipartUploadsOutput) {
	f.ListMultipartUploadsRequestCall.mutex.Lock()
	defer f.ListMultipartUploadsRequestCall.mutex.Unlock()
	f.ListMultipartUploadsRequestCall.CallCount++
	f.ListMultipartUploadsRequestCall.Receives.ListMultipartUploadsInput = param1
	if f.ListMultipartUploadsRequestCall.Stub != nil {
		return f.ListMultipartUploadsRequestCall.Stub(param1)
	}
	return f.ListMultipartUploadsRequestCall.Returns.Request, f.ListMultipartUploadsRequestCall.Returns.ListMultipartUploadsOutput
}
func (f *S3API) ListMultipartUploadsWithContext(param1 context.Context, param2 *s3.ListMultipartUploadsInput, param3 ...request.Option) (*s3.ListMultipartUploadsOutput, error) {
	f.ListMultipartUploadsWithContextCall.mutex.Lock()
	defer f.ListMultipartUploadsWithContextCall.mutex.Unlock()
	f.ListMultipartUploadsWithContextCall.CallCount++
	f.ListMultipartUploadsWithContextCall.Receives.Context = param1
	f.ListMultipartUploadsWithContextCall.Receives.ListMultipartUploadsInput = param2
	f.ListMultipartUploadsWithContextCall.Receives.OptionSlice = param3
	if f.ListMultipartUploadsWithContextCall.Stub != nil {
		return f.ListMultipartUploadsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.ListMultipartUploadsWithContextCall.Returns.ListMultipartUploadsOutput, f.ListMultipartUploadsWithContextCall.Returns.Error
}
func (f *S3API) ListObjectVersions(param1 *s3.ListObjectVersionsInput) (*s3.ListObjectVersionsOutput, error) {
	f.ListObjectVersionsCall.mutex.Lock()
	defer f.ListObjectVersionsCall.mutex.Unlock()
	f.ListObjectVersionsCall.CallCount++
	f.ListObjectVersionsCall.Receives.ListObjectVersionsInput = param1
	if f.ListObjectVersionsCall.Stub != nil {
		return f.ListObjectVersionsCall.Stub(param1)
	}
	return f.ListObjectVersionsCall.Returns.ListObjectVersionsOutput, f.ListObjectVersionsCall.Returns.Error
}
func (f *S3API) ListObjectVersionsPages(param1 *s3.ListObjectVersionsInput, param2 func(*s3.ListObjectVersionsOutput, bool) bool) error {
	f.ListObjectVersionsPagesCall.mutex.Lock()
	defer f.ListObjectVersionsPagesCall.mutex.Unlock()
	f.ListObjectVersionsPagesCall.CallCount++
	f.ListObjectVersionsPagesCall.Receives.ListObjectVersionsInput = param1
	f.ListObjectVersionsPagesCall.Receives.FuncS3ListObjectVersionsOutputBoolBool = param2
	if f.ListObjectVersionsPagesCall.Stub != nil {
		return f.ListObjectVersionsPagesCall.Stub(param1, param2)
	}
	return f.ListObjectVersionsPagesCall.Returns.Error
}
func (f *S3API) ListObjectVersionsPagesWithContext(param1 context.Context, param2 *s3.ListObjectVersionsInput, param3 func(*s3.ListObjectVersionsOutput, bool) bool, param4 ...request.Option) error {
	f.ListObjectVersionsPagesWithContextCall.mutex.Lock()
	defer f.ListObjectVersionsPagesWithContextCall.mutex.Unlock()
	f.ListObjectVersionsPagesWithContextCall.CallCount++
	f.ListObjectVersionsPagesWithContextCall.Receives.Context = param1
	f.ListObjectVersionsPagesWithContextCall.Receives.ListObjectVersionsInput = param2
	f.ListObjectVersionsPagesWithContextCall.Receives.FuncS3ListObjectVersionsOutputBoolBool = param3
	f.ListObjectVersionsPagesWithContextCall.Receives.OptionSlice = param4
	if f.ListObjectVersionsPagesWithContextCall.Stub != nil {
		return f.ListObjectVersionsPagesWithContextCall.Stub(param1, param2, param3, param4...)
	}
	return f.ListObjectVersionsPagesWithContextCall.Returns.Error
}
func (f *S3API) ListObjectVersionsRequest(param1 *s3.ListObjectVersionsInput) (*request.Request, *s3.ListObjectVersionsOutput) {
	f.ListObjectVersionsRequestCall.mutex.Lock()
	defer f.ListObjectVersionsRequestCall.mutex.Unlock()
	f.ListObjectVersionsRequestCall.CallCount++
	f.ListObjectVersionsRequestCall.Receives.ListObjectVersionsInput = param1
	if f.ListObjectVersionsRequestCall.Stub != nil {
		return f.ListObjectVersionsRequestCall.Stub(param1)
	}
	return f.ListObjectVersionsRequestCall.Returns.Request, f.ListObjectVersionsRequestCall.Returns.ListObjectVersionsOutput
}
func (f *S3API) ListObjectVersionsWithContext(param1 context.Context, param2 *s3.ListObjectVersionsInput, param3 ...request.Option) (*s3.ListObjectVersionsOutput, error) {
	f.ListObjectVersionsWithContextCall.mutex.Lock()
	defer f.ListObjectVersionsWithContextCall.mutex.Unlock()
	f.ListObjectVersionsWithContextCall.CallCount++
	f.ListObjectVersionsWithContextCall.Receives.Context = param1
	f.ListObjectVersionsWithContextCall.Receives.ListObjectVersionsInput = param2
	f.ListObjectVersionsWithContextCall.Receives.OptionSlice = param3
	if f.ListObjectVersionsWithContextCall.Stub != nil {
		return f.ListObjectVersionsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.ListObjectVersionsWithContextCall.Returns.ListObjectVersionsOutput, f.ListObjectVersionsWithContextCall.Returns.Error
}
func (f *S3API) ListObjects(param1 *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	f.ListObjectsCall.mutex.Lock()
	defer f.ListObjectsCall.mutex.Unlock()
	f.ListObjectsCall.CallCount++
	f.ListObjectsCall.Receives.ListObjectsInput = param1
	if f.ListObjectsCall.Stub != nil {
		return f.ListObjectsCall.Stub(param1)
	}
	return f.ListObjectsCall.Returns.ListObjectsOutput, f.ListObjectsCall.Returns.Error
}
func (f *S3API) ListObjectsPages(param1 *s3.ListObjectsInput, param2 func(*s3.ListObjectsOutput, bool) bool) error {
	f.ListObjectsPagesCall.mutex.Lock()
	defer f.ListObjectsPagesCall.mutex.Unlock()
	f.ListObjectsPagesCall.CallCount++
	f.ListObjectsPagesCall.Receives.ListObjectsInput = param1
	f.ListObjectsPagesCall.Receives.FuncS3ListObjectsOutputBoolBool = param2
	if f.ListObjectsPagesCall.Stub != nil {
		return f.ListObjectsPagesCall.Stub(param1, param2)
	}
	return f.ListObjectsPagesCall.Returns.Error
}
func (f *S3API) ListObjectsPagesWithContext(param1 context.Context, param2 *s3.ListObjectsInput, param3 func(*s3.ListObjectsOutput, bool) bool, param4 ...request.Option) error {
	f.ListObjectsPagesWithContextCall.mutex.Lock()
	defer f.ListObjectsPagesWithContextCall.mutex.Unlock()
	f.ListObjectsPagesWithContextCall.CallCount++
	f.ListObjectsPagesWithContextCall.Receives.Context = param1
	f.ListObjectsPagesWithContextCall.Receives.ListObjectsInput = param2
	f.ListObjectsPagesWithContextCall.Receives.FuncS3ListObjectsOutputBoolBool = param3
	f.ListObjectsPagesWithContextCall.Receives.OptionSlice = param4
	if f.ListObjectsPagesWithContextCall.Stub != nil {
		return f.ListObjectsPagesWithContextCall.Stub(param1, param2, param3, param4...)
	}
	return f.ListObjectsPagesWithContextCall.Returns.Error
}
func (f *S3API) ListObjectsRequest(param1 *s3.ListObjectsInput) (*request.Request, *s3.ListObjectsOutput) {
	f.ListObjectsRequestCall.mutex.Lock()
	defer f.ListObjectsRequestCall.mutex.Unlock()
	f.ListObjectsRequestCall.CallCount++
	f.ListObjectsRequestCall.Receives.ListObjectsInput = param1
	if f.ListObjectsRequestCall.Stub != nil {
		return f.ListObjectsRequestCall.Stub(param1)
	}
	return f.ListObjectsRequestCall.Returns.Request, f.ListObjectsRequestCall.Returns.ListObjectsOutput
}
func (f *S3API) ListObjectsV2(param1 *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	f.ListObjectsV2Call.mutex.Lock()
	defer f.ListObjectsV2Call.mutex.Unlock()
	f.ListObjectsV2Call.CallCount++
	f.ListObjectsV2Call.Receives.ListObjectsV2Input = param1
	if f.ListObjectsV2Call.Stub != nil {
		return f.ListObjectsV2Call.Stub(param1)
	}
	return f.ListObjectsV2Call.Returns.ListObjectsV2Output, f.ListObjectsV2Call.Returns.Error
}
func (f *S3API) ListObjectsV2Pages(param1 *s3.ListObjectsV2Input, param2 func(*s3.ListObjectsV2Output, bool) bool) error {
	f.ListObjectsV2PagesCall.mutex.Lock()
	defer f.ListObjectsV2PagesCall.mutex.Unlock()
	f.ListObjectsV2PagesCall.CallCount++
	f.ListObjectsV2PagesCall.Receives.ListObjectsV2Input = param1
	f.ListObjectsV2PagesCall.Receives.FuncS3ListObjectsV2OutputBoolBool = param2
	if f.ListObjectsV2PagesCall.Stub != nil {
		return f.ListObjectsV2PagesCall.Stub(param1, param2)
	}
	return f.ListObjectsV2PagesCall.Returns.Error
}
func (f *S3API) ListObjectsV2PagesWithContext(param1 context.Context, param2 *s3.ListObjectsV2Input, param3 func(*s3.ListObjectsV2Output, bool) bool, param4 ...request.Option) error {
	f.ListObjectsV2PagesWithContextCall.mutex.Lock()
	defer f.ListObjectsV2PagesWithContextCall.mutex.Unlock()
	f.ListObjectsV2PagesWithContextCall.CallCount++
	f.ListObjectsV2PagesWithContextCall.Receives.Context = param1
	f.ListObjectsV2PagesWithContextCall.Receives.ListObjectsV2Input = param2
	f.ListObjectsV2PagesWithContextCall.Receives.FuncS3ListObjectsV2OutputBoolBool = param3
	f.ListObjectsV2PagesWithContextCall.Receives.OptionSlice = param4
	if f.ListObjectsV2PagesWithContextCall.Stub != nil {
		return f.ListObjectsV2PagesWithContextCall.Stub(param1, param2, param3, param4...)
	}
	return f.ListObjectsV2PagesWithContextCall.Returns.Error
}
func (f *S3API) ListObjectsV2Request(param1 *s3.ListObjectsV2Input) (*request.Request, *s3.ListObjectsV2Output) {
	f.ListObjectsV2RequestCall.mutex.Lock()
	defer f.ListObjectsV2RequestCall.mutex.Unlock()
	f.ListObjectsV2RequestCall.CallCount++
	f.ListObjectsV2RequestCall.Receives.ListObjectsV2Input = param1
	if f.ListObjectsV2RequestCall.Stub != nil {
		return f.ListObjectsV2RequestCall.Stub(param1)
	}
	return f.ListObjectsV2RequestCall.Returns.Request, f.ListObjectsV2RequestCall.Returns.ListObjectsV2Output
}
func (f *S3API) ListObjectsV2WithContext(param1 context.Context, param2 *s3.ListObjectsV2Input, param3 ...request.Option) (*s3.ListObjectsV2Output, error) {
	f.ListObjectsV2WithContextCall.mutex.Lock()
	defer f.ListObjectsV2WithContextCall.mutex.Unlock()
	f.ListObjectsV2WithContextCall.CallCount++
	f.ListObjectsV2WithContextCall.Receives.Context = param1
	f.ListObjectsV2WithContextCall.Receives.ListObjectsV2Input = param2
	f.ListObjectsV2WithContextCall.Receives.OptionSlice = param3
	if f.ListObjectsV2WithContextCall.Stub != nil {
		return f.ListObjectsV2WithContextCall.Stub(param1, param2, param3...)
	}
	return f.ListObjectsV2WithContextCall.Returns.ListObjectsV2Output, f.ListObjectsV2WithContextCall.Returns.Error
}
func (f *S3API) ListObjectsWithContext(param1 context.Context, param2 *s3.ListObjectsInput, param3 ...request.Option) (*s3.ListObjectsOutput, error) {
	f.ListObjectsWithContextCall.mutex.Lock()
	defer f.ListObjectsWithContextCall.mutex.Unlock()
	f.ListObjectsWithContextCall.CallCount++
	f.ListObjectsWithContextCall.Receives.Context = param1
	f.ListObjectsWithContextCall.Receives.ListObjectsInput = param2
	f.ListObjectsWithContextCall.Receives.OptionSlice = param3
	if f.ListObjectsWithContextCall.Stub != nil {
		return f.ListObjectsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.ListObjectsWithContextCall.Returns.ListObjectsOutput, f.ListObjectsWithContextCall.Returns.Error
}
func (f *S3API) ListParts(param1 *s3.ListPartsInput) (*s3.ListPartsOutput, error) {
	f.ListPartsCall.mutex.Lock()
	defer f.ListPartsCall.mutex.Unlock()
	f.ListPartsCall.CallCount++
	f.ListPartsCall.Receives.ListPartsInput = param1
	if f.ListPartsCall.Stub != nil {
		return f.ListPartsCall.Stub(param1)
	}
	return f.ListPartsCall.Returns.ListPartsOutput, f.ListPartsCall.Returns.Error
}
func (f *S3API) ListPartsPages(param1 *s3.ListPartsInput, param2 func(*s3.ListPartsOutput, bool) bool) error {
	f.ListPartsPagesCall.mutex.Lock()
	defer f.ListPartsPagesCall.mutex.Unlock()
	f.ListPartsPagesCall.CallCount++
	f.ListPartsPagesCall.Receives.ListPartsInput = param1
	f.ListPartsPagesCall.Receives.FuncS3ListPartsOutputBoolBool = param2
	if f.ListPartsPagesCall.Stub != nil {
		return f.ListPartsPagesCall.Stub(param1, param2)
	}
	return f.ListPartsPagesCall.Returns.Error
}
func (f *S3API) ListPartsPagesWithContext(param1 context.Context, param2 *s3.ListPartsInput, param3 func(*s3.ListPartsOutput, bool) bool, param4 ...request.Option) error {
	f.ListPartsPagesWithContextCall.mutex.Lock()
	defer f.ListPartsPagesWithContextCall.mutex.Unlock()
	f.ListPartsPagesWithContextCall.CallCount++
	f.ListPartsPagesWithContextCall.Receives.Context = param1
	f.ListPartsPagesWithContextCall.Receives.ListPartsInput = param2
	f.ListPartsPagesWithContextCall.Receives.FuncS3ListPartsOutputBoolBool = param3
	f.ListPartsPagesWithContextCall.Receives.OptionSlice = param4
	if f.ListPartsPagesWithContextCall.Stub != nil {
		return f.ListPartsPagesWithContextCall.Stub(param1, param2, param3, param4...)
	}
	return f.ListPartsPagesWithContextCall.Returns.Error
}
func (f *S3API) ListPartsRequest(param1 *s3.ListPartsInput) (*request.Request, *s3.ListPartsOutput) {
	f.ListPartsRequestCall.mutex.Lock()
	defer f.ListPartsRequestCall.mutex.Unlock()
	f.ListPartsRequestCall.CallCount++
	f.ListPartsRequestCall.Receives.ListPartsInput = param1
	if f.ListPartsRequestCall.Stub != nil {
		return f.ListPartsRequestCall.Stub(param1)
	}
	return f.ListPartsRequestCall.Returns.Request, f.ListPartsRequestCall.Returns.ListPartsOutput
}
func (f *S3API) ListPartsWithContext(param1 context.Context, param2 *s3.ListPartsInput, param3 ...request.Option) (*s3.ListPartsOutput, error) {
	f.ListPartsWithContextCall.mutex.Lock()
	defer f.ListPartsWithContextCall.mutex.Unlock()
	f.ListPartsWithContextCall.CallCount++
	f.ListPartsWithContextCall.Receives.Context = param1
	f.ListPartsWithContextCall.Receives.ListPartsInput = param2
	f.ListPartsWithContextCall.Receives.OptionSlice = param3
	if f.ListPartsWithContextCall.Stub != nil {
		return f.ListPartsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.ListPartsWithContextCall.Returns.ListPartsOutput, f.ListPartsWithContextCall.Returns.Error
}
func (f *S3API) PutBucketAccelerateConfiguration(param1 *s3.PutBucketAccelerateConfigurationInput) (*s3.PutBucketAccelerateConfigurationOutput, error) {
	f.PutBucketAccelerateConfigurationCall.mutex.Lock()
	defer f.PutBucketAccelerateConfigurationCall.mutex.Unlock()
	f.PutBucketAccelerateConfigurationCall.CallCount++
	f.PutBucketAccelerateConfigurationCall.Receives.PutBucketAccelerateConfigurationInput = param1
	if f.PutBucketAccelerateConfigurationCall.Stub != nil {
		return f.PutBucketAccelerateConfigurationCall.Stub(param1)
	}
	return f.PutBucketAccelerateConfigurationCall.Returns.PutBucketAccelerateConfigurationOutput, f.PutBucketAccelerateConfigurationCall.Returns.Error
}
func (f *S3API) PutBucketAccelerateConfigurationRequest(param1 *s3.PutBucketAccelerateConfigurationInput) (*request.Request, *s3.PutBucketAccelerateConfigurationOutput) {
	f.PutBucketAccelerateConfigurationRequestCall.mutex.Lock()
	defer f.PutBucketAccelerateConfigurationRequestCall.mutex.Unlock()
	f.PutBucketAccelerateConfigurationRequestCall.CallCount++
	f.PutBucketAccelerateConfigurationRequestCall.Receives.PutBucketAccelerateConfigurationInput = param1
	if f.PutBucketAccelerateConfigurationRequestCall.Stub != nil {
		return f.PutBucketAccelerateConfigurationRequestCall.Stub(param1)
	}
	return f.PutBucketAccelerateConfigurationRequestCall.Returns.Request, f.PutBucketAccelerateConfigurationRequestCall.Returns.PutBucketAccelerateConfigurationOutput
}
func (f *S3API) PutBucketAccelerateConfigurationWithContext(param1 context.Context, param2 *s3.PutBucketAccelerateConfigurationInput, param3 ...request.Option) (*s3.PutBucketAccelerateConfigurationOutput, error) {
	f.PutBucketAccelerateConfigurationWithContextCall.mutex.Lock()
	defer f.PutBucketAccelerateConfigurationWithContextCall.mutex.Unlock()
	f.PutBucketAccelerateConfigurationWithContextCall.CallCount++
	f.PutBucketAccelerateConfigurationWithContextCall.Receives.Context = param1
	f.PutBucketAccelerateConfigurationWithContextCall.Receives.PutBucketAccelerateConfigurationInput = param2
	f.PutBucketAccelerateConfigurationWithContextCall.Receives.OptionSlice = param3
	if f.PutBucketAccelerateConfigurationWithContextCall.Stub != nil {
		return f.PutBucketAccelerateConfigurationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutBucketAccelerateConfigurationWithContextCall.Returns.PutBucketAccelerateConfigurationOutput, f.PutBucketAccelerateConfigurationWithContextCall.Returns.Error
}
func (f *S3API) PutBucketAcl(param1 *s3.PutBucketAclInput) (*s3.PutBucketAclOutput, error) {
	f.PutBucketAclCall.mutex.Lock()
	defer f.PutBucketAclCall.mutex.Unlock()
	f.PutBucketAclCall.CallCount++
	f.PutBucketAclCall.Receives.PutBucketAclInput = param1
	if f.PutBucketAclCall.Stub != nil {
		return f.PutBucketAclCall.Stub(param1)
	}
	return f.PutBucketAclCall.Returns.PutBucketAclOutput, f.PutBucketAclCall.Returns.Error
}
func (f *S3API) PutBucketAclRequest(param1 *s3.PutBucketAclInput) (*request.Request, *s3.PutBucketAclOutput) {
	f.PutBucketAclRequestCall.mutex.Lock()
	defer f.PutBucketAclRequestCall.mutex.Unlock()
	f.PutBucketAclRequestCall.CallCount++
	f.PutBucketAclRequestCall.Receives.PutBucketAclInput = param1
	if f.PutBucketAclRequestCall.Stub != nil {
		return f.PutBucketAclRequestCall.Stub(param1)
	}
	return f.PutBucketAclRequestCall.Returns.Request, f.PutBucketAclRequestCall.Returns.PutBucketAclOutput
}
func (f *S3API) PutBucketAclWithContext(param1 context.Context, param2 *s3.PutBucketAclInput, param3 ...request.Option) (*s3.PutBucketAclOutput, error) {
	f.PutBucketAclWithContextCall.mutex.Lock()
	defer f.PutBucketAclWithContextCall.mutex.Unlock()
	f.PutBucketAclWithContextCall.CallCount++
	f.PutBucketAclWithContextCall.Receives.Context = param1
	f.PutBucketAclWithContextCall.Receives.PutBucketAclInput = param2
	f.PutBucketAclWithContextCall.Receives.OptionSlice = param3
	if f.PutBucketAclWithContextCall.Stub != nil {
		return f.PutBucketAclWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutBucketAclWithContextCall.Returns.PutBucketAclOutput, f.PutBucketAclWithContextCall.Returns.Error
}
func (f *S3API) PutBucketAnalyticsConfiguration(param1 *s3.PutBucketAnalyticsConfigurationInput) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
	f.PutBucketAnalyticsConfigurationCall.mutex.Lock()
	defer f.PutBucketAnalyticsConfigurationCall.mutex.Unlock()
	f.PutBucketAnalyticsConfigurationCall.CallCount++
	f.PutBucketAnalyticsConfigurationCall.Receives.PutBucketAnalyticsConfigurationInput = param1
	if f.PutBucketAnalyticsConfigurationCall.Stub != nil {
		return f.PutBucketAnalyticsConfigurationCall.Stub(param1)
	}
	return f.PutBucketAnalyticsConfigurationCall.Returns.PutBucketAnalyticsConfigurationOutput, f.PutBucketAnalyticsConfigurationCall.Returns.Error
}
func (f *S3API) PutBucketAnalyticsConfigurationRequest(param1 *s3.PutBucketAnalyticsConfigurationInput) (*request.Request, *s3.PutBucketAnalyticsConfigurationOutput) {
	f.PutBucketAnalyticsConfigurationRequestCall.mutex.Lock()
	defer f.PutBucketAnalyticsConfigurationRequestCall.mutex.Unlock()
	f.PutBucketAnalyticsConfigurationRequestCall.CallCount++
	f.PutBucketAnalyticsConfigurationRequestCall.Receives.PutBucketAnalyticsConfigurationInput = param1
	if f.PutBucketAnalyticsConfigurationRequestCall.Stub != nil {
		return f.PutBucketAnalyticsConfigurationRequestCall.Stub(param1)
	}
	return f.PutBucketAnalyticsConfigurationRequestCall.Returns.Request, f.PutBucketAnalyticsConfigurationRequestCall.Returns.PutBucketAnalyticsConfigurationOutput
}
func (f *S3API) PutBucketAnalyticsConfigurationWithContext(param1 context.Context, param2 *s3.PutBucketAnalyticsConfigurationInput, param3 ...request.Option) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
	f.PutBucketAnalyticsConfigurationWithContextCall.mutex.Lock()
	defer f.PutBucketAnalyticsConfigurationWithContextCall.mutex.Unlock()
	f.PutBucketAnalyticsConfigurationWithContextCall.CallCount++
	f.PutBucketAnalyticsConfigurationWithContextCall.Receives.Context = param1
	f.PutBucketAnalyticsConfigurationWithContextCall.Receives.PutBucketAnalyticsConfigurationInput = param2
	f.PutBucketAnalyticsConfigurationWithContextCall.Receives.OptionSlice = param3
	if f.PutBucketAnalyticsConfigurationWithContextCall.Stub != nil {
		return f.PutBucketAnalyticsConfigurationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutBucketAnalyticsConfigurationWithContextCall.Returns.PutBucketAnalyticsConfigurationOutput, f.PutBucketAnalyticsConfigurationWithContextCall.Returns.Error
}
func (f *S3API) PutBucketCors(param1 *s3.PutBucketCorsInput) (*s3.PutBucketCorsOutput, error) {
	f.PutBucketCorsCall.mutex.Lock()
	defer f.PutBucketCorsCall.mutex.Unlock()
	f.PutBucketCorsCall.CallCount++
	f.PutBucketCorsCall.Receives.PutBucketCorsInput = param1
	if f.PutBucketCorsCall.Stub != nil {
		return f.PutBucketCorsCall.Stub(param1)
	}
	return f.PutBucketCorsCall.Returns.PutBucketCorsOutput, f.PutBucketCorsCall.Returns.Error
}
func (f *S3API) PutBucketCorsRequest(param1 *s3.PutBucketCorsInput) (*request.Request, *s3.PutBucketCorsOutput) {
	f.PutBucketCorsRequestCall.mutex.Lock()
	defer f.PutBucketCorsRequestCall.mutex.Unlock()
	f.PutBucketCorsRequestCall.CallCount++
	f.PutBucketCorsRequestCall.Receives.PutBucketCorsInput = param1
	if f.PutBucketCorsRequestCall.Stub != nil {
		return f.PutBucketCorsRequestCall.Stub(param1)
	}
	return f.PutBucketCorsRequestCall.Returns.Request, f.PutBucketCorsRequestCall.Returns.PutBucketCorsOutput
}
func (f *S3API) PutBucketCorsWithContext(param1 context.Context, param2 *s3.PutBucketCorsInput, param3 ...request.Option) (*s3.PutBucketCorsOutput, error) {
	f.PutBucketCorsWithContextCall.mutex.Lock()
	defer f.PutBucketCorsWithContextCall.mutex.Unlock()
	f.PutBucketCorsWithContextCall.CallCount++
	f.PutBucketCorsWithContextCall.Receives.Context = param1
	f.PutBucketCorsWithContextCall.Receives.PutBucketCorsInput = param2
	f.PutBucketCorsWithContextCall.Receives.OptionSlice = param3
	if f.PutBucketCorsWithContextCall.Stub != nil {
		return f.PutBucketCorsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutBucketCorsWithContextCall.Returns.PutBucketCorsOutput, f.PutBucketCorsWithContextCall.Returns.Error
}
func (f *S3API) PutBucketEncryption(param1 *s3.PutBucketEncryptionInput) (*s3.PutBucketEncryptionOutput, error) {
	f.PutBucketEncryptionCall.mutex.Lock()
	defer f.PutBucketEncryptionCall.mutex.Unlock()
	f.PutBucketEncryptionCall.CallCount++
	f.PutBucketEncryptionCall.Receives.PutBucketEncryptionInput = param1
	if f.PutBucketEncryptionCall.Stub != nil {
		return f.PutBucketEncryptionCall.Stub(param1)
	}
	return f.PutBucketEncryptionCall.Returns.PutBucketEncryptionOutput, f.PutBucketEncryptionCall.Returns.Error
}
func (f *S3API) PutBucketEncryptionRequest(param1 *s3.PutBucketEncryptionInput) (*request.Request, *s3.PutBucketEncryptionOutput) {
	f.PutBucketEncryptionRequestCall.mutex.Lock()
	defer f.PutBucketEncryptionRequestCall.mutex.Unlock()
	f.PutBucketEncryptionRequestCall.CallCount++
	f.PutBucketEncryptionRequestCall.Receives.PutBucketEncryptionInput = param1
	if f.PutBucketEncryptionRequestCall.Stub != nil {
		return f.PutBucketEncryptionRequestCall.Stub(param1)
	}
	return f.PutBucketEncryptionRequestCall.Returns.Request, f.PutBucketEncryptionRequestCall.Returns.PutBucketEncryptionOutput
}
func (f *S3API) PutBucketEncryptionWithContext(param1 context.Context, param2 *s3.PutBucketEncryptionInput, param3 ...request.Option) (*s3.PutBucketEncryptionOutput, error) {
	f.PutBucketEncryptionWithContextCall.mutex.Lock()
	defer f.PutBucketEncryptionWithContextCall.mutex.Unlock()
	f.PutBucketEncryptionWithContextCall.CallCount++
	f.PutBucketEncryptionWithContextCall.Receives.Context = param1
	f.PutBucketEncryptionWithContextCall.Receives.PutBucketEncryptionInput = param2
	f.PutBucketEncryptionWithContextCall.Receives.OptionSlice = param3
	if f.PutBucketEncryptionWithContextCall.Stub != nil {
		return f.PutBucketEncryptionWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutBucketEncryptionWithContextCall.Returns.PutBucketEncryptionOutput, f.PutBucketEncryptionWithContextCall.Returns.Error
}
func (f *S3API) PutBucketIntelligentTieringConfiguration(param1 *s3.PutBucketIntelligentTieringConfigurationInput) (*s3.PutBucketIntelligentTieringConfigurationOutput, error) {
	f.PutBucketIntelligentTieringConfigurationCall.mutex.Lock()
	defer f.PutBucketIntelligentTieringConfigurationCall.mutex.Unlock()
	f.PutBucketIntelligentTieringConfigurationCall.CallCount++
	f.PutBucketIntelligentTieringConfigurationCall.Receives.PutBucketIntelligentTieringConfigurationInput = param1
	if f.PutBucketIntelligentTieringConfigurationCall.Stub != nil {
		return f.PutBucketIntelligentTieringConfigurationCall.Stub(param1)
	}
	return f.PutBucketIntelligentTieringConfigurationCall.Returns.PutBucketIntelligentTieringConfigurationOutput, f.PutBucketIntelligentTieringConfigurationCall.Returns.Error
}
func (f *S3API) PutBucketIntelligentTieringConfigurationRequest(param1 *s3.PutBucketIntelligentTieringConfigurationInput) (*request.Request, *s3.PutBucketIntelligentTieringConfigurationOutput) {
	f.PutBucketIntelligentTieringConfigurationRequestCall.mutex.Lock()
	defer f.PutBucketIntelligentTieringConfigurationRequestCall.mutex.Unlock()
	f.PutBucketIntelligentTieringConfigurationRequestCall.CallCount++
	f.PutBucketIntelligentTieringConfigurationRequestCall.Receives.PutBucketIntelligentTieringConfigurationInput = param1
	if f.PutBucketIntelligentTieringConfigurationRequestCall.Stub != nil {
		return f.PutBucketIntelligentTieringConfigurationRequestCall.Stub(param1)
	}
	return f.PutBucketIntelligentTieringConfigurationRequestCall.Returns.Request, f.PutBucketIntelligentTieringConfigurationRequestCall.Returns.PutBucketIntelligentTieringConfigurationOutput
}
func (f *S3API) PutBucketIntelligentTieringConfigurationWithContext(param1 context.Context, param2 *s3.PutBucketIntelligentTieringConfigurationInput, param3 ...request.Option) (*s3.PutBucketIntelligentTieringConfigurationOutput, error) {
	f.PutBucketIntelligentTieringConfigurationWithContextCall.mutex.Lock()
	defer f.PutBucketIntelligentTieringConfigurationWithContextCall.mutex.Unlock()
	f.PutBucketIntelligentTieringConfigurationWithContextCall.CallCount++
	f.PutBucketIntelligentTieringConfigurationWithContextCall.Receives.Context = param1
	f.PutBucketIntelligentTieringConfigurationWithContextCall.Receives.PutBucketIntelligentTieringConfigurationInput = param2
	f.PutBucketIntelligentTieringConfigurationWithContextCall.Receives.OptionSlice = param3
	if f.PutBucketIntelligentTieringConfigurationWithContextCall.Stub != nil {
		return f.PutBucketIntelligentTieringConfigurationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutBucketIntelligentTieringConfigurationWithContextCall.Returns.PutBucketIntelligentTieringConfigurationOutput, f.PutBucketIntelligentTieringConfigurationWithContextCall.Returns.Error
}
func (f *S3API) PutBucketInventoryConfiguration(param1 *s3.PutBucketInventoryConfigurationInput) (*s3.PutBucketInventoryConfigurationOutput, error) {
	f.PutBucketInventoryConfigurationCall.mutex.Lock()
	defer f.PutBucketInventoryConfigurationCall.mutex.Unlock()
	f.PutBucketInventoryConfigurationCall.CallCount++
	f.PutBucketInventoryConfigurationCall.Receives.PutBucketInventoryConfigurationInput = param1
	if f.PutBucketInventoryConfigurationCall.Stub != nil {
		return f.PutBucketInventoryConfigurationCall.Stub(param1)
	}
	return f.PutBucketInventoryConfigurationCall.Returns.PutBucketInventoryConfigurationOutput, f.PutBucketInventoryConfigurationCall.Returns.Error
}
func (f *S3API) PutBucketInventoryConfigurationRequest(param1 *s3.PutBucketInventoryConfigurationInput) (*request.Request, *s3.PutBucketInventoryConfigurationOutput) {
	f.PutBucketInventoryConfigurationRequestCall.mutex.Lock()
	defer f.PutBucketInventoryConfigurationRequestCall.mutex.Unlock()
	f.PutBucketInventoryConfigurationRequestCall.CallCount++
	f.PutBucketInventoryConfigurationRequestCall.Receives.PutBucketInventoryConfigurationInput = param1
	if f.PutBucketInventoryConfigurationRequestCall.Stub != nil {
		return f.PutBucketInventoryConfigurationRequestCall.Stub(param1)
	}
	return f.PutBucketInventoryConfigurationRequestCall.Returns.Request, f.PutBucketInventoryConfigurationRequestCall.Returns.PutBucketInventoryConfigurationOutput
}
func (f *S3API) PutBucketInventoryConfigurationWithContext(param1 context.Context, param2 *s3.PutBucketInventoryConfigurationInput, param3 ...request.Option) (*s3.PutBucketInventoryConfigurationOutput, error) {
	f.PutBucketInventoryConfigurationWithContextCall.mutex.Lock()
	defer f.PutBucketInventoryConfigurationWithContextCall.mutex.Unlock()
	f.PutBucketInventoryConfigurationWithContextCall.CallCount++
	f.PutBucketInventoryConfigurationWithContextCall.Receives.Context = param1
	f.PutBucketInventoryConfigurationWithContextCall.Receives.PutBucketInventoryConfigurationInput = param2
	f.PutBucketInventoryConfigurationWithContextCall.Receives.OptionSlice = param3
	if f.PutBucketInventoryConfigurationWithContextCall.Stub != nil {
		return f.PutBucketInventoryConfigurationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutBucketInventoryConfigurationWithContextCall.Returns.PutBucketInventoryConfigurationOutput, f.PutBucketInventoryConfigurationWithContextCall.Returns.Error
}
func (f *S3API) PutBucketLifecycle(param1 *s3.PutBucketLifecycleInput) (*s3.PutBucketLifecycleOutput, error) {
	f.PutBucketLifecycleCall.mutex.Lock()
	defer f.PutBucketLifecycleCall.mutex.Unlock()
	f.PutBucketLifecycleCall.CallCount++
	f.PutBucketLifecycleCall.Receives.PutBucketLifecycleInput = param1
	if f.PutBucketLifecycleCall.Stub != nil {
		return f.PutBucketLifecycleCall.Stub(param1)
	}
	return f.PutBucketLifecycleCall.Returns.PutBucketLifecycleOutput, f.PutBucketLifecycleCall.Returns.Error
}
func (f *S3API) PutBucketLifecycleConfiguration(param1 *s3.PutBucketLifecycleConfigurationInput) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	f.PutBucketLifecycleConfigurationCall.mutex.Lock()
	defer f.PutBucketLifecycleConfigurationCall.mutex.Unlock()
	f.PutBucketLifecycleConfigurationCall.CallCount++
	f.PutBucketLifecycleConfigurationCall.Receives.PutBucketLifecycleConfigurationInput = param1
	if f.PutBucketLifecycleConfigurationCall.Stub != nil {
		return f.PutBucketLifecycleConfigurationCall.Stub(param1)
	}
	return f.PutBucketLifecycleConfigurationCall.Returns.PutBucketLifecycleConfigurationOutput, f.PutBucketLifecycleConfigurationCall.Returns.Error
}
func (f *S3API) PutBucketLifecycleConfigurationRequest(param1 *s3.PutBucketLifecycleConfigurationInput) (*request.Request, *s3.PutBucketLifecycleConfigurationOutput) {
	f.PutBucketLifecycleConfigurationRequestCall.mutex.Lock()
	defer f.PutBucketLifecycleConfigurationRequestCall.mutex.Unlock()
	f.PutBucketLifecycleConfigurationRequestCall.CallCount++
	f.PutBucketLifecycleConfigurationRequestCall.Receives.PutBucketLifecycleConfigurationInput = param1
	if f.PutBucketLifecycleConfigurationRequestCall.Stub != nil {
		return f.PutBucketLifecycleConfigurationRequestCall.Stub(param1)
	}
	return f.PutBucketLifecycleConfigurationRequestCall.Returns.Request, f.PutBucketLifecycleConfigurationRequestCall.Returns.PutBucketLifecycleConfigurationOutput
}
func (f *S3API) PutBucketLifecycleConfigurationWithContext(param1 context.Context, param2 *s3.PutBucketLifecycleConfigurationInput, param3 ...request.Option) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	f.PutBucketLifecycleConfigurationWithContextCall.mutex.Lock()
	defer f.PutBucketLifecycleConfigurationWithContextCall.mutex.Unlock()
	f.PutBucketLifecycleConfigurationWithContextCall.CallCount++
	f.PutBucketLifecycleConfigurationWithContextCall.Receives.Context = param1
	f.PutBucketLifecycleConfigurationWithContextCall.Receives.PutBucketLifecycleConfigurationInput = param2
	f.PutBucketLifecycleConfigurationWithContextCall.Receives.OptionSlice = param3
	if f.PutBucketLifecycleConfigurationWithContextCall.Stub != nil {
		return f.PutBucketLifecycleConfigurationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutBucketLifecycleConfigurationWithContextCall.Returns.PutBucketLifecycleConfigurationOutput, f.PutBucketLifecycleConfigurationWithContextCall.Returns.Error
}
func (f *S3API) PutBucketLifecycleRequest(param1 *s3.PutBucketLifecycleInput) (*request.Request, *s3.PutBucketLifecycleOutput) {
	f.PutBucketLifecycleRequestCall.mutex.Lock()
	defer f.PutBucketLifecycleRequestCall.mutex.Unlock()
	f.PutBucketLifecycleRequestCall.CallCount++
	f.PutBucketLifecycleRequestCall.Receives.PutBucketLifecycleInput = param1
	if f.PutBucketLifecycleRequestCall.Stub != nil {
		return f.PutBucketLifecycleRequestCall.Stub(param1)
	}
	return f.PutBucketLifecycleRequestCall.Returns.Request, f.PutBucketLifecycleRequestCall.Returns.PutBucketLifecycleOutput
}
func (f *S3API) PutBucketLifecycleWithContext(param1 context.Context, param2 *s3.PutBucketLifecycleInput, param3 ...request.Option) (*s3.PutBucketLifecycleOutput, error) {
	f.PutBucketLifecycleWithContextCall.mutex.Lock()
	defer f.PutBucketLifecycleWithContextCall.mutex.Unlock()
	f.PutBucketLifecycleWithContextCall.CallCount++
	f.PutBucketLifecycleWithContextCall.Receives.Context = param1
	f.PutBucketLifecycleWithContextCall.Receives.PutBucketLifecycleInput = param2
	f.PutBucketLifecycleWithContextCall.Receives.OptionSlice = param3
	if f.PutBucketLifecycleWithContextCall.Stub != nil {
		return f.PutBucketLifecycleWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutBucketLifecycleWithContextCall.Returns.PutBucketLifecycleOutput, f.PutBucketLifecycleWithContextCall.Returns.Error
}
func (f *S3API) PutBucketLogging(param1 *s3.PutBucketLoggingInput) (*s3.PutBucketLoggingOutput, error) {
	f.PutBucketLoggingCall.mutex.Lock()
	defer f.PutBucketLoggingCall.mutex.Unlock()
	f.PutBucketLoggingCall.CallCount++
	f.PutBucketLoggingCall.Receives.PutBucketLoggingInput = param1
	if f.PutBucketLoggingCall.Stub != nil {
		return f.PutBucketLoggingCall.Stub(param1)
	}
	return f.PutBucketLoggingCall.Returns.PutBucketLoggingOutput, f.PutBucketLoggingCall.Returns.Error
}
func (f *S3API) PutBucketLoggingRequest(param1 *s3.PutBucketLoggingInput) (*request.Request, *s3.PutBucketLoggingOutput) {
	f.PutBucketLoggingRequestCall.mutex.Lock()
	defer f.PutBucketLoggingRequestCall.mutex.Unlock()
	f.PutBucketLoggingRequestCall.CallCount++
	f.PutBucketLoggingRequestCall.Receives.PutBucketLoggingInput = param1
	if f.PutBucketLoggingRequestCall.Stub != nil {
		return f.PutBucketLoggingRequestCall.Stub(param1)
	}
	return f.PutBucketLoggingRequestCall.Returns.Request, f.PutBucketLoggingRequestCall.Returns.PutBucketLoggingOutput
}
func (f *S3API) PutBucketLoggingWithContext(param1 context.Context, param2 *s3.PutBucketLoggingInput, param3 ...request.Option) (*s3.PutBucketLoggingOutput, error) {
	f.PutBucketLoggingWithContextCall.mutex.Lock()
	defer f.PutBucketLoggingWithContextCall.mutex.Unlock()
	f.PutBucketLoggingWithContextCall.CallCount++
	f.PutBucketLoggingWithContextCall.Receives.Context = param1
	f.PutBucketLoggingWithContextCall.Receives.PutBucketLoggingInput = param2
	f.PutBucketLoggingWithContextCall.Receives.OptionSlice = param3
	if f.PutBucketLoggingWithContextCall.Stub != nil {
		return f.PutBucketLoggingWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutBucketLoggingWithContextCall.Returns.PutBucketLoggingOutput, f.PutBucketLoggingWithContextCall.Returns.Error
}
func (f *S3API) PutBucketMetricsConfiguration(param1 *s3.PutBucketMetricsConfigurationInput) (*s3.PutBucketMetricsConfigurationOutput, error) {
	f.PutBucketMetricsConfigurationCall.mutex.Lock()
	defer f.PutBucketMetricsConfigurationCall.mutex.Unlock()
	f.PutBucketMetricsConfigurationCall.CallCount++
	f.PutBucketMetricsConfigurationCall.Receives.PutBucketMetricsConfigurationInput = param1
	if f.PutBucketMetricsConfigurationCall.Stub != nil {
		return f.PutBucketMetricsConfigurationCall.Stub(param1)
	}
	return f.PutBucketMetricsConfigurationCall.Returns.PutBucketMetricsConfigurationOutput, f.PutBucketMetricsConfigurationCall.Returns.Error
}
func (f *S3API) PutBucketMetricsConfigurationRequest(param1 *s3.PutBucketMetricsConfigurationInput) (*request.Request, *s3.PutBucketMetricsConfigurationOutput) {
	f.PutBucketMetricsConfigurationRequestCall.mutex.Lock()
	defer f.PutBucketMetricsConfigurationRequestCall.mutex.Unlock()
	f.PutBucketMetricsConfigurationRequestCall.CallCount++
	f.PutBucketMetricsConfigurationRequestCall.Receives.PutBucketMetricsConfigurationInput = param1
	if f.PutBucketMetricsConfigurationRequestCall.Stub != nil {
		return f.PutBucketMetricsConfigurationRequestCall.Stub(param1)
	}
	return f.PutBucketMetricsConfigurationRequestCall.Returns.Request, f.PutBucketMetricsConfigurationRequestCall.Returns.PutBucketMetricsConfigurationOutput
}
func (f *S3API) PutBucketMetricsConfigurationWithContext(param1 context.Context, param2 *s3.PutBucketMetricsConfigurationInput, param3 ...request.Option) (*s3.PutBucketMetricsConfigurationOutput, error) {
	f.PutBucketMetricsConfigurationWithContextCall.mutex.Lock()
	defer f.PutBucketMetricsConfigurationWithContextCall.mutex.Unlock()
	f.PutBucketMetricsConfigurationWithContextCall.CallCount++
	f.PutBucketMetricsConfigurationWithContextCall.Receives.Context = param1
	f.PutBucketMetricsConfigurationWithContextCall.Receives.PutBucketMetricsConfigurationInput = param2
	f.PutBucketMetricsConfigurationWithContextCall.Receives.OptionSlice = param3
	if f.PutBucketMetricsConfigurationWithContextCall.Stub != nil {
		return f.PutBucketMetricsConfigurationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutBucketMetricsConfigurationWithContextCall.Returns.PutBucketMetricsConfigurationOutput, f.PutBucketMetricsConfigurationWithContextCall.Returns.Error
}
func (f *S3API) PutBucketNotification(param1 *s3.PutBucketNotificationInput) (*s3.PutBucketNotificationOutput, error) {
	f.PutBucketNotificationCall.mutex.Lock()
	defer f.PutBucketNotificationCall.mutex.Unlock()
	f.PutBucketNotificationCall.CallCount++
	f.PutBucketNotificationCall.Receives.PutBucketNotificationInput = param1
	if f.PutBucketNotificationCall.Stub != nil {
		return f.PutBucketNotificationCall.Stub(param1)
	}
	return f.PutBucketNotificationCall.Returns.PutBucketNotificationOutput, f.PutBucketNotificationCall.Returns.Error
}
func (f *S3API) PutBucketNotificationConfiguration(param1 *s3.PutBucketNotificationConfigurationInput) (*s3.PutBucketNotificationConfigurationOutput, error) {
	f.PutBucketNotificationConfigurationCall.mutex.Lock()
	defer f.PutBucketNotificationConfigurationCall.mutex.Unlock()
	f.PutBucketNotificationConfigurationCall.CallCount++
	f.PutBucketNotificationConfigurationCall.Receives.PutBucketNotificationConfigurationInput = param1
	if f.PutBucketNotificationConfigurationCall.Stub != nil {
		return f.PutBucketNotificationConfigurationCall.Stub(param1)
	}
	return f.PutBucketNotificationConfigurationCall.Returns.PutBucketNotificationConfigurationOutput, f.PutBucketNotificationConfigurationCall.Returns.Error
}
func (f *S3API) PutBucketNotificationConfigurationRequest(param1 *s3.PutBucketNotificationConfigurationInput) (*request.Request, *s3.PutBucketNotificationConfigurationOutput) {
	f.PutBucketNotificationConfigurationRequestCall.mutex.Lock()
	defer f.PutBucketNotificationConfigurationRequestCall.mutex.Unlock()
	f.PutBucketNotificationConfigurationRequestCall.CallCount++
	f.PutBucketNotificationConfigurationRequestCall.Receives.PutBucketNotificationConfigurationInput = param1
	if f.PutBucketNotificationConfigurationRequestCall.Stub != nil {
		return f.PutBucketNotificationConfigurationRequestCall.Stub(param1)
	}
	return f.PutBucketNotificationConfigurationRequestCall.Returns.Request, f.PutBucketNotificationConfigurationRequestCall.Returns.PutBucketNotificationConfigurationOutput
}
func (f *S3API) PutBucketNotificationConfigurationWithContext(param1 context.Context, param2 *s3.PutBucketNotificationConfigurationInput, param3 ...request.Option) (*s3.PutBucketNotificationConfigurationOutput, error) {
	f.PutBucketNotificationConfigurationWithContextCall.mutex.Lock()
	defer f.PutBucketNotificationConfigurationWithContextCall.mutex.Unlock()
	f.PutBucketNotificationConfigurationWithContextCall.CallCount++
	f.PutBucketNotificationConfigurationWithContextCall.Receives.Context = param1
	f.PutBucketNotificationConfigurationWithContextCall.Receives.PutBucketNotificationConfigurationInput = param2
	f.PutBucketNotificationConfigurationWithContextCall.Receives.OptionSlice = param3
	if f.PutBucketNotificationConfigurationWithContextCall.Stub != nil {
		return f.PutBucketNotificationConfigurationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutBucketNotificationConfigurationWithContextCall.Returns.PutBucketNotificationConfigurationOutput, f.PutBucketNotificationConfigurationWithContextCall.Returns.Error
}
func (f *S3API) PutBucketNotificationRequest(param1 *s3.PutBucketNotificationInput) (*request.Request, *s3.PutBucketNotificationOutput) {
	f.PutBucketNotificationRequestCall.mutex.Lock()
	defer f.PutBucketNotificationRequestCall.mutex.Unlock()
	f.PutBucketNotificationRequestCall.CallCount++
	f.PutBucketNotificationRequestCall.Receives.PutBucketNotificationInput = param1
	if f.PutBucketNotificationRequestCall.Stub != nil {
		return f.PutBucketNotificationRequestCall.Stub(param1)
	}
	return f.PutBucketNotificationRequestCall.Returns.Request, f.PutBucketNotificationRequestCall.Returns.PutBucketNotificationOutput
}
func (f *S3API) PutBucketNotificationWithContext(param1 context.Context, param2 *s3.PutBucketNotificationInput, param3 ...request.Option) (*s3.PutBucketNotificationOutput, error) {
	f.PutBucketNotificationWithContextCall.mutex.Lock()
	defer f.PutBucketNotificationWithContextCall.mutex.Unlock()
	f.PutBucketNotificationWithContextCall.CallCount++
	f.PutBucketNotificationWithContextCall.Receives.Context = param1
	f.PutBucketNotificationWithContextCall.Receives.PutBucketNotificationInput = param2
	f.PutBucketNotificationWithContextCall.Receives.OptionSlice = param3
	if f.PutBucketNotificationWithContextCall.Stub != nil {
		return f.PutBucketNotificationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutBucketNotificationWithContextCall.Returns.PutBucketNotificationOutput, f.PutBucketNotificationWithContextCall.Returns.Error
}
func (f *S3API) PutBucketOwnershipControls(param1 *s3.PutBucketOwnershipControlsInput) (*s3.PutBucketOwnershipControlsOutput, error) {
	f.PutBucketOwnershipControlsCall.mutex.Lock()
	defer f.PutBucketOwnershipControlsCall.mutex.Unlock()
	f.PutBucketOwnershipControlsCall.CallCount++
	f.PutBucketOwnershipControlsCall.Receives.PutBucketOwnershipControlsInput = param1
	if f.PutBucketOwnershipControlsCall.Stub != nil {
		return f.PutBucketOwnershipControlsCall.Stub(param1)
	}
	return f.PutBucketOwnershipControlsCall.Returns.PutBucketOwnershipControlsOutput, f.PutBucketOwnershipControlsCall.Returns.Error
}
func (f *S3API) PutBucketOwnershipControlsRequest(param1 *s3.PutBucketOwnershipControlsInput) (*request.Request, *s3.PutBucketOwnershipControlsOutput) {
	f.PutBucketOwnershipControlsRequestCall.mutex.Lock()
	defer f.PutBucketOwnershipControlsRequestCall.mutex.Unlock()
	f.PutBucketOwnershipControlsRequestCall.CallCount++
	f.PutBucketOwnershipControlsRequestCall.Receives.PutBucketOwnershipControlsInput = param1
	if f.PutBucketOwnershipControlsRequestCall.Stub != nil {
		return f.PutBucketOwnershipControlsRequestCall.Stub(param1)
	}
	return f.PutBucketOwnershipControlsRequestCall.Returns.Request, f.PutBucketOwnershipControlsRequestCall.Returns.PutBucketOwnershipControlsOutput
}
func (f *S3API) PutBucketOwnershipControlsWithContext(param1 context.Context, param2 *s3.PutBucketOwnershipControlsInput, param3 ...request.Option) (*s3.PutBucketOwnershipControlsOutput, error) {
	f.PutBucketOwnershipControlsWithContextCall.mutex.Lock()
	defer f.PutBucketOwnershipControlsWithContextCall.mutex.Unlock()
	f.PutBucketOwnershipControlsWithContextCall.CallCount++
	f.PutBucketOwnershipControlsWithContextCall.Receives.Context = param1
	f.PutBucketOwnershipControlsWithContextCall.Receives.PutBucketOwnershipControlsInput = param2
	f.PutBucketOwnershipControlsWithContextCall.Receives.OptionSlice = param3
	if f.PutBucketOwnershipControlsWithContextCall.Stub != nil {
		return f.PutBucketOwnershipControlsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutBucketOwnershipControlsWithContextCall.Returns.PutBucketOwnershipControlsOutput, f.PutBucketOwnershipControlsWithContextCall.Returns.Error
}
func (f *S3API) PutBucketPolicy(param1 *s3.PutBucketPolicyInput) (*s3.PutBucketPolicyOutput, error) {
	f.PutBucketPolicyCall.mutex.Lock()
	defer f.PutBucketPolicyCall.mutex.Unlock()
	f.PutBucketPolicyCall.CallCount++
	f.PutBucketPolicyCall.Receives.PutBucketPolicyInput = param1
	if f.PutBucketPolicyCall.Stub != nil {
		return f.PutBucketPolicyCall.Stub(param1)
	}
	return f.PutBucketPolicyCall.Returns.PutBucketPolicyOutput, f.PutBucketPolicyCall.Returns.Error
}
func (f *S3API) PutBucketPolicyRequest(param1 *s3.PutBucketPolicyInput) (*request.Request, *s3.PutBucketPolicyOutput) {
	f.PutBucketPolicyRequestCall.mutex.Lock()
	defer f.PutBucketPolicyRequestCall.mutex.Unlock()
	f.PutBucketPolicyRequestCall.CallCount++
	f.PutBucketPolicyRequestCall.Receives.PutBucketPolicyInput = param1
	if f.PutBucketPolicyRequestCall.Stub != nil {
		return f.PutBucketPolicyRequestCall.Stub(param1)
	}
	return f.PutBucketPolicyRequestCall.Returns.Request, f.PutBucketPolicyRequestCall.Returns.PutBucketPolicyOutput
}
func (f *S3API) PutBucketPolicyWithContext(param1 context.Context, param2 *s3.PutBucketPolicyInput, param3 ...request.Option) (*s3.PutBucketPolicyOutput, error) {
	f.PutBucketPolicyWithContextCall.mutex.Lock()
	defer f.PutBucketPolicyWithContextCall.mutex.Unlock()
	f.PutBucketPolicyWithContextCall.CallCount++
	f.PutBucketPolicyWithContextCall.Receives.Context = param1
	f.PutBucketPolicyWithContextCall.Receives.PutBucketPolicyInput = param2
	f.PutBucketPolicyWithContextCall.Receives.OptionSlice = param3
	if f.PutBucketPolicyWithContextCall.Stub != nil {
		return f.PutBucketPolicyWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutBucketPolicyWithContextCall.Returns.PutBucketPolicyOutput, f.PutBucketPolicyWithContextCall.Returns.Error
}
func (f *S3API) PutBucketReplication(param1 *s3.PutBucketReplicationInput) (*s3.PutBucketReplicationOutput, error) {
	f.PutBucketReplicationCall.mutex.Lock()
	defer f.PutBucketReplicationCall.mutex.Unlock()
	f.PutBucketReplicationCall.CallCount++
	f.PutBucketReplicationCall.Receives.PutBucketReplicationInput = param1
	if f.PutBucketReplicationCall.Stub != nil {
		return f.PutBucketReplicationCall.Stub(param1)
	}
	return f.PutBucketReplicationCall.Returns.PutBucketReplicationOutput, f.PutBucketReplicationCall.Returns.Error
}
func (f *S3API) PutBucketReplicationRequest(param1 *s3.PutBucketReplicationInput) (*request.Request, *s3.PutBucketReplicationOutput) {
	f.PutBucketReplicationRequestCall.mutex.Lock()
	defer f.PutBucketReplicationRequestCall.mutex.Unlock()
	f.PutBucketReplicationRequestCall.CallCount++
	f.PutBucketReplicationRequestCall.Receives.PutBucketReplicationInput = param1
	if f.PutBucketReplicationRequestCall.Stub != nil {
		return f.PutBucketReplicationRequestCall.Stub(param1)
	}
	return f.PutBucketReplicationRequestCall.Returns.Request, f.PutBucketReplicationRequestCall.Returns.PutBucketReplicationOutput
}
func (f *S3API) PutBucketReplicationWithContext(param1 context.Context, param2 *s3.PutBucketReplicationInput, param3 ...request.Option) (*s3.PutBucketReplicationOutput, error) {
	f.PutBucketReplicationWithContextCall.mutex.Lock()
	defer f.PutBucketReplicationWithContextCall.mutex.Unlock()
	f.PutBucketReplicationWithContextCall.CallCount++
	f.PutBucketReplicationWithContextCall.Receives.Context = param1
	f.PutBucketReplicationWithContextCall.Receives.PutBucketReplicationInput = param2
	f.PutBucketReplicationWithContextCall.Receives.OptionSlice = param3
	if f.PutBucketReplicationWithContextCall.Stub != nil {
		return f.PutBucketReplicationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutBucketReplicationWithContextCall.Returns.PutBucketReplicationOutput, f.PutBucketReplicationWithContextCall.Returns.Error
}
func (f *S3API) PutBucketRequestPayment(param1 *s3.PutBucketRequestPaymentInput) (*s3.PutBucketRequestPaymentOutput, error) {
	f.PutBucketRequestPaymentCall.mutex.Lock()
	defer f.PutBucketRequestPaymentCall.mutex.Unlock()
	f.PutBucketRequestPaymentCall.CallCount++
	f.PutBucketRequestPaymentCall.Receives.PutBucketRequestPaymentInput = param1
	if f.PutBucketRequestPaymentCall.Stub != nil {
		return f.PutBucketRequestPaymentCall.Stub(param1)
	}
	return f.PutBucketRequestPaymentCall.Returns.PutBucketRequestPaymentOutput, f.PutBucketRequestPaymentCall.Returns.Error
}
func (f *S3API) PutBucketRequestPaymentRequest(param1 *s3.PutBucketRequestPaymentInput) (*request.Request, *s3.PutBucketRequestPaymentOutput) {
	f.PutBucketRequestPaymentRequestCall.mutex.Lock()
	defer f.PutBucketRequestPaymentRequestCall.mutex.Unlock()
	f.PutBucketRequestPaymentRequestCall.CallCount++
	f.PutBucketRequestPaymentRequestCall.Receives.PutBucketRequestPaymentInput = param1
	if f.PutBucketRequestPaymentRequestCall.Stub != nil {
		return f.PutBucketRequestPaymentRequestCall.Stub(param1)
	}
	return f.PutBucketRequestPaymentRequestCall.Returns.Request, f.PutBucketRequestPaymentRequestCall.Returns.PutBucketRequestPaymentOutput
}
func (f *S3API) PutBucketRequestPaymentWithContext(param1 context.Context, param2 *s3.PutBucketRequestPaymentInput, param3 ...request.Option) (*s3.PutBucketRequestPaymentOutput, error) {
	f.PutBucketRequestPaymentWithContextCall.mutex.Lock()
	defer f.PutBucketRequestPaymentWithContextCall.mutex.Unlock()
	f.PutBucketRequestPaymentWithContextCall.CallCount++
	f.PutBucketRequestPaymentWithContextCall.Receives.Context = param1
	f.PutBucketRequestPaymentWithContextCall.Receives.PutBucketRequestPaymentInput = param2
	f.PutBucketRequestPaymentWithContextCall.Receives.OptionSlice = param3
	if f.PutBucketRequestPaymentWithContextCall.Stub != nil {
		return f.PutBucketRequestPaymentWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutBucketRequestPaymentWithContextCall.Returns.PutBucketRequestPaymentOutput, f.PutBucketRequestPaymentWithContextCall.Returns.Error
}
func (f *S3API) PutBucketTagging(param1 *s3.PutBucketTaggingInput) (*s3.PutBucketTaggingOutput, error) {
	f.PutBucketTaggingCall.mutex.Lock()
	defer f.PutBucketTaggingCall.mutex.Unlock()
	f.PutBucketTaggingCall.CallCount++
	f.PutBucketTaggingCall.Receives.PutBucketTaggingInput = param1
	if f.PutBucketTaggingCall.Stub != nil {
		return f.PutBucketTaggingCall.Stub(param1)
	}
	return f.PutBucketTaggingCall.Returns.PutBucketTaggingOutput, f.PutBucketTaggingCall.Returns.Error
}
func (f *S3API) PutBucketTaggingRequest(param1 *s3.PutBucketTaggingInput) (*request.Request, *s3.PutBucketTaggingOutput) {
	f.PutBucketTaggingRequestCall.mutex.Lock()
	defer f.PutBucketTaggingRequestCall.mutex.Unlock()
	f.PutBucketTaggingRequestCall.CallCount++
	f.PutBucketTaggingRequestCall.Receives.PutBucketTaggingInput = param1
	if f.PutBucketTaggingRequestCall.Stub != nil {
		return f.PutBucketTaggingRequestCall.Stub(param1)
	}
	return f.PutBucketTaggingRequestCall.Returns.Request, f.PutBucketTaggingRequestCall.Returns.PutBucketTaggingOutput
}
func (f *S3API) PutBucketTaggingWithContext(param1 context.Context, param2 *s3.PutBucketTaggingInput, param3 ...request.Option) (*s3.PutBucketTaggingOutput, error) {
	f.PutBucketTaggingWithContextCall.mutex.Lock()
	defer f.PutBucketTaggingWithContextCall.mutex.Unlock()
	f.PutBucketTaggingWithContextCall.CallCount++
	f.PutBucketTaggingWithContextCall.Receives.Context = param1
	f.PutBucketTaggingWithContextCall.Receives.PutBucketTaggingInput = param2
	f.PutBucketTaggingWithContextCall.Receives.OptionSlice = param3
	if f.PutBucketTaggingWithContextCall.Stub != nil {
		return f.PutBucketTaggingWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutBucketTaggingWithContextCall.Returns.PutBucketTaggingOutput, f.PutBucketTaggingWithContextCall.Returns.Error
}
func (f *S3API) PutBucketVersioning(param1 *s3.PutBucketVersioningInput) (*s3.PutBucketVersioningOutput, error) {
	f.PutBucketVersioningCall.mutex.Lock()
	defer f.PutBucketVersioningCall.mutex.Unlock()
	f.PutBucketVersioningCall.CallCount++
	f.PutBucketVersioningCall.Receives.PutBucketVersioningInput = param1
	if f.PutBucketVersioningCall.Stub != nil {
		return f.PutBucketVersioningCall.Stub(param1)
	}
	return f.PutBucketVersioningCall.Returns.PutBucketVersioningOutput, f.PutBucketVersioningCall.Returns.Error
}
func (f *S3API) PutBucketVersioningRequest(param1 *s3.PutBucketVersioningInput) (*request.Request, *s3.PutBucketVersioningOutput) {
	f.PutBucketVersioningRequestCall.mutex.Lock()
	defer f.PutBucketVersioningRequestCall.mutex.Unlock()
	f.PutBucketVersioningRequestCall.CallCount++
	f.PutBucketVersioningRequestCall.Receives.PutBucketVersioningInput = param1
	if f.PutBucketVersioningRequestCall.Stub != nil {
		return f.PutBucketVersioningRequestCall.Stub(param1)
	}
	return f.PutBucketVersioningRequestCall.Returns.Request, f.PutBucketVersioningRequestCall.Returns.PutBucketVersioningOutput
}
func (f *S3API) PutBucketVersioningWithContext(param1 context.Context, param2 *s3.PutBucketVersioningInput, param3 ...request.Option) (*s3.PutBucketVersioningOutput, error) {
	f.PutBucketVersioningWithContextCall.mutex.Lock()
	defer f.PutBucketVersioningWithContextCall.mutex.Unlock()
	f.PutBucketVersioningWithContextCall.CallCount++
	f.PutBucketVersioningWithContextCall.Receives.Context = param1
	f.PutBucketVersioningWithContextCall.Receives.PutBucketVersioningInput = param2
	f.PutBucketVersioningWithContextCall.Receives.OptionSlice = param3
	if f.PutBucketVersioningWithContextCall.Stub != nil {
		return f.PutBucketVersioningWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutBucketVersioningWithContextCall.Returns.PutBucketVersioningOutput, f.PutBucketVersioningWithContextCall.Returns.Error
}
func (f *S3API) PutBucketWebsite(param1 *s3.PutBucketWebsiteInput) (*s3.PutBucketWebsiteOutput, error) {
	f.PutBucketWebsiteCall.mutex.Lock()
	defer f.PutBucketWebsiteCall.mutex.Unlock()
	f.PutBucketWebsiteCall.CallCount++
	f.PutBucketWebsiteCall.Receives.PutBucketWebsiteInput = param1
	if f.PutBucketWebsiteCall.Stub != nil {
		return f.PutBucketWebsiteCall.Stub(param1)
	}
	return f.PutBucketWebsiteCall.Returns.PutBucketWebsiteOutput, f.PutBucketWebsiteCall.Returns.Error
}
func (f *S3API) PutBucketWebsiteRequest(param1 *s3.PutBucketWebsiteInput) (*request.Request, *s3.PutBucketWebsiteOutput) {
	f.PutBucketWebsiteRequestCall.mutex.Lock()
	defer f.PutBucketWebsiteRequestCall.mutex.Unlock()
	f.PutBucketWebsiteRequestCall.CallCount++
	f.PutBucketWebsiteRequestCall.Receives.PutBucketWebsiteInput = param1
	if f.PutBucketWebsiteRequestCall.Stub != nil {
		return f.PutBucketWebsiteRequestCall.Stub(param1)
	}
	return f.PutBucketWebsiteRequestCall.Returns.Request, f.PutBucketWebsiteRequestCall.Returns.PutBucketWebsiteOutput
}
func (f *S3API) PutBucketWebsiteWithContext(param1 context.Context, param2 *s3.PutBucketWebsiteInput, param3 ...request.Option) (*s3.PutBucketWebsiteOutput, error) {
	f.PutBucketWebsiteWithContextCall.mutex.Lock()
	defer f.PutBucketWebsiteWithContextCall.mutex.Unlock()
	f.PutBucketWebsiteWithContextCall.CallCount++
	f.PutBucketWebsiteWithContextCall.Receives.Context = param1
	f.PutBucketWebsiteWithContextCall.Receives.PutBucketWebsiteInput = param2
	f.PutBucketWebsiteWithContextCall.Receives.OptionSlice = param3
	if f.PutBucketWebsiteWithContextCall.Stub != nil {
		return f.PutBucketWebsiteWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutBucketWebsiteWithContextCall.Returns.PutBucketWebsiteOutput, f.PutBucketWebsiteWithContextCall.Returns.Error
}
func (f *S3API) PutObject(param1 *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	f.PutObjectCall.mutex.Lock()
	defer f.PutObjectCall.mutex.Unlock()
	f.PutObjectCall.CallCount++
	f.PutObjectCall.Receives.PutObjectInput = param1
	if f.PutObjectCall.Stub != nil {
		return f.PutObjectCall.Stub(param1)
	}
	return f.PutObjectCall.Returns.PutObjectOutput, f.PutObjectCall.Returns.Error
}
func (f *S3API) PutObjectAcl(param1 *s3.PutObjectAclInput) (*s3.PutObjectAclOutput, error) {
	f.PutObjectAclCall.mutex.Lock()
	defer f.PutObjectAclCall.mutex.Unlock()
	f.PutObjectAclCall.CallCount++
	f.PutObjectAclCall.Receives.PutObjectAclInput = param1
	if f.PutObjectAclCall.Stub != nil {
		return f.PutObjectAclCall.Stub(param1)
	}
	return f.PutObjectAclCall.Returns.PutObjectAclOutput, f.PutObjectAclCall.Returns.Error
}
func (f *S3API) PutObjectAclRequest(param1 *s3.PutObjectAclInput) (*request.Request, *s3.PutObjectAclOutput) {
	f.PutObjectAclRequestCall.mutex.Lock()
	defer f.PutObjectAclRequestCall.mutex.Unlock()
	f.PutObjectAclRequestCall.CallCount++
	f.PutObjectAclRequestCall.Receives.PutObjectAclInput = param1
	if f.PutObjectAclRequestCall.Stub != nil {
		return f.PutObjectAclRequestCall.Stub(param1)
	}
	return f.PutObjectAclRequestCall.Returns.Request, f.PutObjectAclRequestCall.Returns.PutObjectAclOutput
}
func (f *S3API) PutObjectAclWithContext(param1 context.Context, param2 *s3.PutObjectAclInput, param3 ...request.Option) (*s3.PutObjectAclOutput, error) {
	f.PutObjectAclWithContextCall.mutex.Lock()
	defer f.PutObjectAclWithContextCall.mutex.Unlock()
	f.PutObjectAclWithContextCall.CallCount++
	f.PutObjectAclWithContextCall.Receives.Context = param1
	f.PutObjectAclWithContextCall.Receives.PutObjectAclInput = param2
	f.PutObjectAclWithContextCall.Receives.OptionSlice = param3
	if f.PutObjectAclWithContextCall.Stub != nil {
		return f.PutObjectAclWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutObjectAclWithContextCall.Returns.PutObjectAclOutput, f.PutObjectAclWithContextCall.Returns.Error
}
func (f *S3API) PutObjectLegalHold(param1 *s3.PutObjectLegalHoldInput) (*s3.PutObjectLegalHoldOutput, error) {
	f.PutObjectLegalHoldCall.mutex.Lock()
	defer f.PutObjectLegalHoldCall.mutex.Unlock()
	f.PutObjectLegalHoldCall.CallCount++
	f.PutObjectLegalHoldCall.Receives.PutObjectLegalHoldInput = param1
	if f.PutObjectLegalHoldCall.Stub != nil {
		return f.PutObjectLegalHoldCall.Stub(param1)
	}
	return f.PutObjectLegalHoldCall.Returns.PutObjectLegalHoldOutput, f.PutObjectLegalHoldCall.Returns.Error
}
func (f *S3API) PutObjectLegalHoldRequest(param1 *s3.PutObjectLegalHoldInput) (*request.Request, *s3.PutObjectLegalHoldOutput) {
	f.PutObjectLegalHoldRequestCall.mutex.Lock()
	defer f.PutObjectLegalHoldRequestCall.mutex.Unlock()
	f.PutObjectLegalHoldRequestCall.CallCount++
	f.PutObjectLegalHoldRequestCall.Receives.PutObjectLegalHoldInput = param1
	if f.PutObjectLegalHoldRequestCall.Stub != nil {
		return f.PutObjectLegalHoldRequestCall.Stub(param1)
	}
	return f.PutObjectLegalHoldRequestCall.Returns.Request, f.PutObjectLegalHoldRequestCall.Returns.PutObjectLegalHoldOutput
}
func (f *S3API) PutObjectLegalHoldWithContext(param1 context.Context, param2 *s3.PutObjectLegalHoldInput, param3 ...request.Option) (*s3.PutObjectLegalHoldOutput, error) {
	f.PutObjectLegalHoldWithContextCall.mutex.Lock()
	defer f.PutObjectLegalHoldWithContextCall.mutex.Unlock()
	f.PutObjectLegalHoldWithContextCall.CallCount++
	f.PutObjectLegalHoldWithContextCall.Receives.Context = param1
	f.PutObjectLegalHoldWithContextCall.Receives.PutObjectLegalHoldInput = param2
	f.PutObjectLegalHoldWithContextCall.Receives.OptionSlice = param3
	if f.PutObjectLegalHoldWithContextCall.Stub != nil {
		return f.PutObjectLegalHoldWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutObjectLegalHoldWithContextCall.Returns.PutObjectLegalHoldOutput, f.PutObjectLegalHoldWithContextCall.Returns.Error
}
func (f *S3API) PutObjectLockConfiguration(param1 *s3.PutObjectLockConfigurationInput) (*s3.PutObjectLockConfigurationOutput, error) {
	f.PutObjectLockConfigurationCall.mutex.Lock()
	defer f.PutObjectLockConfigurationCall.mutex.Unlock()
	f.PutObjectLockConfigurationCall.CallCount++
	f.PutObjectLockConfigurationCall.Receives.PutObjectLockConfigurationInput = param1
	if f.PutObjectLockConfigurationCall.Stub != nil {
		return f.PutObjectLockConfigurationCall.Stub(param1)
	}
	return f.PutObjectLockConfigurationCall.Returns.PutObjectLockConfigurationOutput, f.PutObjectLockConfigurationCall.Returns.Error
}
func (f *S3API) PutObjectLockConfigurationRequest(param1 *s3.PutObjectLockConfigurationInput) (*request.Request, *s3.PutObjectLockConfigurationOutput) {
	f.PutObjectLockConfigurationRequestCall.mutex.Lock()
	defer f.PutObjectLockConfigurationRequestCall.mutex.Unlock()
	f.PutObjectLockConfigurationRequestCall.CallCount++
	f.PutObjectLockConfigurationRequestCall.Receives.PutObjectLockConfigurationInput = param1
	if f.PutObjectLockConfigurationRequestCall.Stub != nil {
		return f.PutObjectLockConfigurationRequestCall.Stub(param1)
	}
	return f.PutObjectLockConfigurationRequestCall.Returns.Request, f.PutObjectLockConfigurationRequestCall.Returns.PutObjectLockConfigurationOutput
}
func (f *S3API) PutObjectLockConfigurationWithContext(param1 context.Context, param2 *s3.PutObjectLockConfigurationInput, param3 ...request.Option) (*s3.PutObjectLockConfigurationOutput, error) {
	f.PutObjectLockConfigurationWithContextCall.mutex.Lock()
	defer f.PutObjectLockConfigurationWithContextCall.mutex.Unlock()
	f.PutObjectLockConfigurationWithContextCall.CallCount++
	f.PutObjectLockConfigurationWithContextCall.Receives.Context = param1
	f.PutObjectLockConfigurationWithContextCall.Receives.PutObjectLockConfigurationInput = param2
	f.PutObjectLockConfigurationWithContextCall.Receives.OptionSlice = param3
	if f.PutObjectLockConfigurationWithContextCall.Stub != nil {
		return f.PutObjectLockConfigurationWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutObjectLockConfigurationWithContextCall.Returns.PutObjectLockConfigurationOutput, f.PutObjectLockConfigurationWithContextCall.Returns.Error
}
func (f *S3API) PutObjectRequest(param1 *s3.PutObjectInput) (*request.Request, *s3.PutObjectOutput) {
	f.PutObjectRequestCall.mutex.Lock()
	defer f.PutObjectRequestCall.mutex.Unlock()
	f.PutObjectRequestCall.CallCount++
	f.PutObjectRequestCall.Receives.PutObjectInput = param1
	if f.PutObjectRequestCall.Stub != nil {
		return f.PutObjectRequestCall.Stub(param1)
	}
	return f.PutObjectRequestCall.Returns.Request, f.PutObjectRequestCall.Returns.PutObjectOutput
}
func (f *S3API) PutObjectRetention(param1 *s3.PutObjectRetentionInput) (*s3.PutObjectRetentionOutput, error) {
	f.PutObjectRetentionCall.mutex.Lock()
	defer f.PutObjectRetentionCall.mutex.Unlock()
	f.PutObjectRetentionCall.CallCount++
	f.PutObjectRetentionCall.Receives.PutObjectRetentionInput = param1
	if f.PutObjectRetentionCall.Stub != nil {
		return f.PutObjectRetentionCall.Stub(param1)
	}
	return f.PutObjectRetentionCall.Returns.PutObjectRetentionOutput, f.PutObjectRetentionCall.Returns.Error
}
func (f *S3API) PutObjectRetentionRequest(param1 *s3.PutObjectRetentionInput) (*request.Request, *s3.PutObjectRetentionOutput) {
	f.PutObjectRetentionRequestCall.mutex.Lock()
	defer f.PutObjectRetentionRequestCall.mutex.Unlock()
	f.PutObjectRetentionRequestCall.CallCount++
	f.PutObjectRetentionRequestCall.Receives.PutObjectRetentionInput = param1
	if f.PutObjectRetentionRequestCall.Stub != nil {
		return f.PutObjectRetentionRequestCall.Stub(param1)
	}
	return f.PutObjectRetentionRequestCall.Returns.Request, f.PutObjectRetentionRequestCall.Returns.PutObjectRetentionOutput
}
func (f *S3API) PutObjectRetentionWithContext(param1 context.Context, param2 *s3.PutObjectRetentionInput, param3 ...request.Option) (*s3.PutObjectRetentionOutput, error) {
	f.PutObjectRetentionWithContextCall.mutex.Lock()
	defer f.PutObjectRetentionWithContextCall.mutex.Unlock()
	f.PutObjectRetentionWithContextCall.CallCount++
	f.PutObjectRetentionWithContextCall.Receives.Context = param1
	f.PutObjectRetentionWithContextCall.Receives.PutObjectRetentionInput = param2
	f.PutObjectRetentionWithContextCall.Receives.OptionSlice = param3
	if f.PutObjectRetentionWithContextCall.Stub != nil {
		return f.PutObjectRetentionWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutObjectRetentionWithContextCall.Returns.PutObjectRetentionOutput, f.PutObjectRetentionWithContextCall.Returns.Error
}
func (f *S3API) PutObjectTagging(param1 *s3.PutObjectTaggingInput) (*s3.PutObjectTaggingOutput, error) {
	f.PutObjectTaggingCall.mutex.Lock()
	defer f.PutObjectTaggingCall.mutex.Unlock()
	f.PutObjectTaggingCall.CallCount++
	f.PutObjectTaggingCall.Receives.PutObjectTaggingInput = param1
	if f.PutObjectTaggingCall.Stub != nil {
		return f.PutObjectTaggingCall.Stub(param1)
	}
	return f.PutObjectTaggingCall.Returns.PutObjectTaggingOutput, f.PutObjectTaggingCall.Returns.Error
}
func (f *S3API) PutObjectTaggingRequest(param1 *s3.PutObjectTaggingInput) (*request.Request, *s3.PutObjectTaggingOutput) {
	f.PutObjectTaggingRequestCall.mutex.Lock()
	defer f.PutObjectTaggingRequestCall.mutex.Unlock()
	f.PutObjectTaggingRequestCall.CallCount++
	f.PutObjectTaggingRequestCall.Receives.PutObjectTaggingInput = param1
	if f.PutObjectTaggingRequestCall.Stub != nil {
		return f.PutObjectTaggingRequestCall.Stub(param1)
	}
	return f.PutObjectTaggingRequestCall.Returns.Request, f.PutObjectTaggingRequestCall.Returns.PutObjectTaggingOutput
}
func (f *S3API) PutObjectTaggingWithContext(param1 context.Context, param2 *s3.PutObjectTaggingInput, param3 ...request.Option) (*s3.PutObjectTaggingOutput, error) {
	f.PutObjectTaggingWithContextCall.mutex.Lock()
	defer f.PutObjectTaggingWithContextCall.mutex.Unlock()
	f.PutObjectTaggingWithContextCall.CallCount++
	f.PutObjectTaggingWithContextCall.Receives.Context = param1
	f.PutObjectTaggingWithContextCall.Receives.PutObjectTaggingInput = param2
	f.PutObjectTaggingWithContextCall.Receives.OptionSlice = param3
	if f.PutObjectTaggingWithContextCall.Stub != nil {
		return f.PutObjectTaggingWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutObjectTaggingWithContextCall.Returns.PutObjectTaggingOutput, f.PutObjectTaggingWithContextCall.Returns.Error
}
func (f *S3API) PutObjectWithContext(param1 context.Context, param2 *s3.PutObjectInput, param3 ...request.Option) (*s3.PutObjectOutput, error) {
	f.PutObjectWithContextCall.mutex.Lock()
	defer f.PutObjectWithContextCall.mutex.Unlock()
	f.PutObjectWithContextCall.CallCount++
	f.PutObjectWithContextCall.Receives.Context = param1
	f.PutObjectWithContextCall.Receives.PutObjectInput = param2
	f.PutObjectWithContextCall.Receives.OptionSlice = param3
	if f.PutObjectWithContextCall.Stub != nil {
		return f.PutObjectWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutObjectWithContextCall.Returns.PutObjectOutput, f.PutObjectWithContextCall.Returns.Error
}
func (f *S3API) PutPublicAccessBlock(param1 *s3.PutPublicAccessBlockInput) (*s3.PutPublicAccessBlockOutput, error) {
	f.PutPublicAccessBlockCall.mutex.Lock()
	defer f.PutPublicAccessBlockCall.mutex.Unlock()
	f.PutPublicAccessBlockCall.CallCount++
	f.PutPublicAccessBlockCall.Receives.PutPublicAccessBlockInput = param1
	if f.PutPublicAccessBlockCall.Stub != nil {
		return f.PutPublicAccessBlockCall.Stub(param1)
	}
	return f.PutPublicAccessBlockCall.Returns.PutPublicAccessBlockOutput, f.PutPublicAccessBlockCall.Returns.Error
}
func (f *S3API) PutPublicAccessBlockRequest(param1 *s3.PutPublicAccessBlockInput) (*request.Request, *s3.PutPublicAccessBlockOutput) {
	f.PutPublicAccessBlockRequestCall.mutex.Lock()
	defer f.PutPublicAccessBlockRequestCall.mutex.Unlock()
	f.PutPublicAccessBlockRequestCall.CallCount++
	f.PutPublicAccessBlockRequestCall.Receives.PutPublicAccessBlockInput = param1
	if f.PutPublicAccessBlockRequestCall.Stub != nil {
		return f.PutPublicAccessBlockRequestCall.Stub(param1)
	}
	return f.PutPublicAccessBlockRequestCall.Returns.Request, f.PutPublicAccessBlockRequestCall.Returns.PutPublicAccessBlockOutput
}
func (f *S3API) PutPublicAccessBlockWithContext(param1 context.Context, param2 *s3.PutPublicAccessBlockInput, param3 ...request.Option) (*s3.PutPublicAccessBlockOutput, error) {
	f.PutPublicAccessBlockWithContextCall.mutex.Lock()
	defer f.PutPublicAccessBlockWithContextCall.mutex.Unlock()
	f.PutPublicAccessBlockWithContextCall.CallCount++
	f.PutPublicAccessBlockWithContextCall.Receives.Context = param1
	f.PutPublicAccessBlockWithContextCall.Receives.PutPublicAccessBlockInput = param2
	f.PutPublicAccessBlockWithContextCall.Receives.OptionSlice = param3
	if f.PutPublicAccessBlockWithContextCall.Stub != nil {
		return f.PutPublicAccessBlockWithContextCall.Stub(param1, param2, param3...)
	}
	return f.PutPublicAccessBlockWithContextCall.Returns.PutPublicAccessBlockOutput, f.PutPublicAccessBlockWithContextCall.Returns.Error
}
func (f *S3API) RestoreObject(param1 *s3.RestoreObjectInput) (*s3.RestoreObjectOutput, error) {
	f.RestoreObjectCall.mutex.Lock()
	defer f.RestoreObjectCall.mutex.Unlock()
	f.RestoreObjectCall.CallCount++
	f.RestoreObjectCall.Receives.RestoreObjectInput = param1
	if f.RestoreObjectCall.Stub != nil {
		return f.RestoreObjectCall.Stub(param1)
	}
	return f.RestoreObjectCall.Returns.RestoreObjectOutput, f.RestoreObjectCall.Returns.Error
}
func (f *S3API) RestoreObjectRequest(param1 *s3.RestoreObjectInput) (*request.Request, *s3.RestoreObjectOutput) {
	f.RestoreObjectRequestCall.mutex.Lock()
	defer f.RestoreObjectRequestCall.mutex.Unlock()
	f.RestoreObjectRequestCall.CallCount++
	f.RestoreObjectRequestCall.Receives.RestoreObjectInput = param1
	if f.RestoreObjectRequestCall.Stub != nil {
		return f.RestoreObjectRequestCall.Stub(param1)
	}
	return f.RestoreObjectRequestCall.Returns.Request, f.RestoreObjectRequestCall.Returns.RestoreObjectOutput
}
func (f *S3API) RestoreObjectWithContext(param1 context.Context, param2 *s3.RestoreObjectInput, param3 ...request.Option) (*s3.RestoreObjectOutput, error) {
	f.RestoreObjectWithContextCall.mutex.Lock()
	defer f.RestoreObjectWithContextCall.mutex.Unlock()
	f.RestoreObjectWithContextCall.CallCount++
	f.RestoreObjectWithContextCall.Receives.Context = param1
	f.RestoreObjectWithContextCall.Receives.RestoreObjectInput = param2
	f.RestoreObjectWithContextCall.Receives.OptionSlice = param3
	if f.RestoreObjectWithContextCall.Stub != nil {
		return f.RestoreObjectWithContextCall.Stub(param1, param2, param3...)
	}
	return f.RestoreObjectWithContextCall.Returns.RestoreObjectOutput, f.RestoreObjectWithContextCall.Returns.Error
}
func (f *S3API) SelectObjectContent(param1 *s3.SelectObjectContentInput) (*s3.SelectObjectContentOutput, error) {
	f.SelectObjectContentCall.mutex.Lock()
	defer f.SelectObjectContentCall.mutex.Unlock()
	f.SelectObjectContentCall.CallCount++
	f.SelectObjectContentCall.Receives.SelectObjectContentInput = param1
	if f.SelectObjectContentCall.Stub != nil {
		return f.SelectObjectContentCall.Stub(param1)
	}
	return f.SelectObjectContentCall.Returns.SelectObjectContentOutput, f.SelectObjectContentCall.Returns.Error
}
func (f *S3API) SelectObjectContentRequest(param1 *s3.SelectObjectContentInput) (*request.Request, *s3.SelectObjectContentOutput) {
	f.SelectObjectContentRequestCall.mutex.Lock()
	defer f.SelectObjectContentRequestCall.mutex.Unlock()
	f.SelectObjectContentRequestCall.CallCount++
	f.SelectObjectContentRequestCall.Receives.SelectObjectContentInput = param1
	if f.SelectObjectContentRequestCall.Stub != nil {
		return f.SelectObjectContentRequestCall.Stub(param1)
	}
	return f.SelectObjectContentRequestCall.Returns.Request, f.SelectObjectContentRequestCall.Returns.SelectObjectContentOutput
}
func (f *S3API) SelectObjectContentWithContext(param1 context.Context, param2 *s3.SelectObjectContentInput, param3 ...request.Option) (*s3.SelectObjectContentOutput, error) {
	f.SelectObjectContentWithContextCall.mutex.Lock()
	defer f.SelectObjectContentWithContextCall.mutex.Unlock()
	f.SelectObjectContentWithContextCall.CallCount++
	f.SelectObjectContentWithContextCall.Receives.Context = param1
	f.SelectObjectContentWithContextCall.Receives.SelectObjectContentInput = param2
	f.SelectObjectContentWithContextCall.Receives.OptionSlice = param3
	if f.SelectObjectContentWithContextCall.Stub != nil {
		return f.SelectObjectContentWithContextCall.Stub(param1, param2, param3...)
	}
	return f.SelectObjectContentWithContextCall.Returns.SelectObjectContentOutput, f.SelectObjectContentWithContextCall.Returns.Error
}
func (f *S3API) UploadPart(param1 *s3.UploadPartInput) (*s3.UploadPartOutput, error) {
	f.UploadPartCall.mutex.Lock()
	defer f.UploadPartCall.mutex.Unlock()
	f.UploadPartCall.CallCount++
	f.UploadPartCall.Receives.UploadPartInput = param1
	if f.UploadPartCall.Stub != nil {
		return f.UploadPartCall.Stub(param1)
	}
	return f.UploadPartCall.Returns.UploadPartOutput, f.UploadPartCall.Returns.Error
}
func (f *S3API) UploadPartCopy(param1 *s3.UploadPartCopyInput) (*s3.UploadPartCopyOutput, error) {
	f.UploadPartCopyCall.mutex.Lock()
	defer f.UploadPartCopyCall.mutex.Unlock()
	f.UploadPartCopyCall.CallCount++
	f.UploadPartCopyCall.Receives.UploadPartCopyInput = param1
	if f.UploadPartCopyCall.Stub != nil {
		return f.UploadPartCopyCall.Stub(param1)
	}
	return f.UploadPartCopyCall.Returns.UploadPartCopyOutput, f.UploadPartCopyCall.Returns.Error
}
func (f *S3API) UploadPartCopyRequest(param1 *s3.UploadPartCopyInput) (*request.Request, *s3.UploadPartCopyOutput) {
	f.UploadPartCopyRequestCall.mutex.Lock()
	defer f.UploadPartCopyRequestCall.mutex.Unlock()
	f.UploadPartCopyRequestCall.CallCount++
	f.UploadPartCopyRequestCall.Receives.UploadPartCopyInput = param1
	if f.UploadPartCopyRequestCall.Stub != nil {
		return f.UploadPartCopyRequestCall.Stub(param1)
	}
	return f.UploadPartCopyRequestCall.Returns.Request, f.UploadPartCopyRequestCall.Returns.UploadPartCopyOutput
}
func (f *S3API) UploadPartCopyWithContext(param1 context.Context, param2 *s3.UploadPartCopyInput, param3 ...request.Option) (*s3.UploadPartCopyOutput, error) {
	f.UploadPartCopyWithContextCall.mutex.Lock()
	defer f.UploadPartCopyWithContextCall.mutex.Unlock()
	f.UploadPartCopyWithContextCall.CallCount++
	f.UploadPartCopyWithContextCall.Receives.Context = param1
	f.UploadPartCopyWithContextCall.Receives.UploadPartCopyInput = param2
	f.UploadPartCopyWithContextCall.Receives.OptionSlice = param3
	if f.UploadPartCopyWithContextCall.Stub != nil {
		return f.UploadPartCopyWithContextCall.Stub(param1, param2, param3...)
	}
	return f.UploadPartCopyWithContextCall.Returns.UploadPartCopyOutput, f.UploadPartCopyWithContextCall.Returns.Error
}
func (f *S3API) UploadPartRequest(param1 *s3.UploadPartInput) (*request.Request, *s3.UploadPartOutput) {
	f.UploadPartRequestCall.mutex.Lock()
	defer f.UploadPartRequestCall.mutex.Unlock()
	f.UploadPartRequestCall.CallCount++
	f.UploadPartRequestCall.Receives.UploadPartInput = param1
	if f.UploadPartRequestCall.Stub != nil {
		return f.UploadPartRequestCall.Stub(param1)
	}
	return f.UploadPartRequestCall.Returns.Request, f.UploadPartRequestCall.Returns.UploadPartOutput
}
func (f *S3API) UploadPartWithContext(param1 context.Context, param2 *s3.UploadPartInput, param3 ...request.Option) (*s3.UploadPartOutput, error) {
	f.UploadPartWithContextCall.mutex.Lock()
	defer f.UploadPartWithContextCall.mutex.Unlock()
	f.UploadPartWithContextCall.CallCount++
	f.UploadPartWithContextCall.Receives.Context = param1
	f.UploadPartWithContextCall.Receives.UploadPartInput = param2
	f.UploadPartWithContextCall.Receives.OptionSlice = param3
	if f.UploadPartWithContextCall.Stub != nil {
		return f.UploadPartWithContextCall.Stub(param1, param2, param3...)
	}
	return f.UploadPartWithContextCall.Returns.UploadPartOutput, f.UploadPartWithContextCall.Returns.Error
}
func (f *S3API) WaitUntilBucketExists(param1 *s3.HeadBucketInput) error {
	f.WaitUntilBucketExistsCall.mutex.Lock()
	defer f.WaitUntilBucketExistsCall.mutex.Unlock()
	f.WaitUntilBucketExistsCall.CallCount++
	f.WaitUntilBucketExistsCall.Receives.HeadBucketInput = param1
	if f.WaitUntilBucketExistsCall.Stub != nil {
		return f.WaitUntilBucketExistsCall.Stub(param1)
	}
	return f.WaitUntilBucketExistsCall.Returns.Error
}
func (f *S3API) WaitUntilBucketExistsWithContext(param1 context.Context, param2 *s3.HeadBucketInput, param3 ...request.WaiterOption) error {
	f.WaitUntilBucketExistsWithContextCall.mutex.Lock()
	defer f.WaitUntilBucketExistsWithContextCall.mutex.Unlock()
	f.WaitUntilBucketExistsWithContextCall.CallCount++
	f.WaitUntilBucketExistsWithContextCall.Receives.Context = param1
	f.WaitUntilBucketExistsWithContextCall.Receives.HeadBucketInput = param2
	f.WaitUntilBucketExistsWithContextCall.Receives.WaiterOptionSlice = param3
	if f.WaitUntilBucketExistsWithContextCall.Stub != nil {
		return f.WaitUntilBucketExistsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.WaitUntilBucketExistsWithContextCall.Returns.Error
}
func (f *S3API) WaitUntilBucketNotExists(param1 *s3.HeadBucketInput) error {
	f.WaitUntilBucketNotExistsCall.mutex.Lock()
	defer f.WaitUntilBucketNotExistsCall.mutex.Unlock()
	f.WaitUntilBucketNotExistsCall.CallCount++
	f.WaitUntilBucketNotExistsCall.Receives.HeadBucketInput = param1
	if f.WaitUntilBucketNotExistsCall.Stub != nil {
		return f.WaitUntilBucketNotExistsCall.Stub(param1)
	}
	return f.WaitUntilBucketNotExistsCall.Returns.Error
}
func (f *S3API) WaitUntilBucketNotExistsWithContext(param1 context.Context, param2 *s3.HeadBucketInput, param3 ...request.WaiterOption) error {
	f.WaitUntilBucketNotExistsWithContextCall.mutex.Lock()
	defer f.WaitUntilBucketNotExistsWithContextCall.mutex.Unlock()
	f.WaitUntilBucketNotExistsWithContextCall.CallCount++
	f.WaitUntilBucketNotExistsWithContextCall.Receives.Context = param1
	f.WaitUntilBucketNotExistsWithContextCall.Receives.HeadBucketInput = param2
	f.WaitUntilBucketNotExistsWithContextCall.Receives.WaiterOptionSlice = param3
	if f.WaitUntilBucketNotExistsWithContextCall.Stub != nil {
		return f.WaitUntilBucketNotExistsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.WaitUntilBucketNotExistsWithContextCall.Returns.Error
}
func (f *S3API) WaitUntilObjectExists(param1 *s3.HeadObjectInput) error {
	f.WaitUntilObjectExistsCall.mutex.Lock()
	defer f.WaitUntilObjectExistsCall.mutex.Unlock()
	f.WaitUntilObjectExistsCall.CallCount++
	f.WaitUntilObjectExistsCall.Receives.HeadObjectInput = param1
	if f.WaitUntilObjectExistsCall.Stub != nil {
		return f.WaitUntilObjectExistsCall.Stub(param1)
	}
	return f.WaitUntilObjectExistsCall.Returns.Error
}
func (f *S3API) WaitUntilObjectExistsWithContext(param1 context.Context, param2 *s3.HeadObjectInput, param3 ...request.WaiterOption) error {
	f.WaitUntilObjectExistsWithContextCall.mutex.Lock()
	defer f.WaitUntilObjectExistsWithContextCall.mutex.Unlock()
	f.WaitUntilObjectExistsWithContextCall.CallCount++
	f.WaitUntilObjectExistsWithContextCall.Receives.Context = param1
	f.WaitUntilObjectExistsWithContextCall.Receives.HeadObjectInput = param2
	f.WaitUntilObjectExistsWithContextCall.Receives.WaiterOptionSlice = param3
	if f.WaitUntilObjectExistsWithContextCall.Stub != nil {
		return f.WaitUntilObjectExistsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.WaitUntilObjectExistsWithContextCall.Returns.Error
}
func (f *S3API) WaitUntilObjectNotExists(param1 *s3.HeadObjectInput) error {
	f.WaitUntilObjectNotExistsCall.mutex.Lock()
	defer f.WaitUntilObjectNotExistsCall.mutex.Unlock()
	f.WaitUntilObjectNotExistsCall.CallCount++
	f.WaitUntilObjectNotExistsCall.Receives.HeadObjectInput = param1
	if f.WaitUntilObjectNotExistsCall.Stub != nil {
		return f.WaitUntilObjectNotExistsCall.Stub(param1)
	}
	return f.WaitUntilObjectNotExistsCall.Returns.Error
}
func (f *S3API) WaitUntilObjectNotExistsWithContext(param1 context.Context, param2 *s3.HeadObjectInput, param3 ...request.WaiterOption) error {
	f.WaitUntilObjectNotExistsWithContextCall.mutex.Lock()
	defer f.WaitUntilObjectNotExistsWithContextCall.mutex.Unlock()
	f.WaitUntilObjectNotExistsWithContextCall.CallCount++
	f.WaitUntilObjectNotExistsWithContextCall.Receives.Context = param1
	f.WaitUntilObjectNotExistsWithContextCall.Receives.HeadObjectInput = param2
	f.WaitUntilObjectNotExistsWithContextCall.Receives.WaiterOptionSlice = param3
	if f.WaitUntilObjectNotExistsWithContextCall.Stub != nil {
		return f.WaitUntilObjectNotExistsWithContextCall.Stub(param1, param2, param3...)
	}
	return f.WaitUntilObjectNotExistsWithContextCall.Returns.Error
}
func (f *S3API) WriteGetObjectResponse(param1 *s3.WriteGetObjectResponseInput) (*s3.WriteGetObjectResponseOutput, error) {
	f.WriteGetObjectResponseCall.mutex.Lock()
	defer f.WriteGetObjectResponseCall.mutex.Unlock()
	f.WriteGetObjectResponseCall.CallCount++
	f.WriteGetObjectResponseCall.Receives.WriteGetObjectResponseInput = param1
	if f.WriteGetObjectResponseCall.Stub != nil {
		return f.WriteGetObjectResponseCall.Stub(param1)
	}
	return f.WriteGetObjectResponseCall.Returns.WriteGetObjectResponseOutput, f.WriteGetObjectResponseCall.Returns.Error
}
func (f *S3API) WriteGetObjectResponseRequest(param1 *s3.WriteGetObjectResponseInput) (*request.Request, *s3.WriteGetObjectResponseOutput) {
	f.WriteGetObjectResponseRequestCall.mutex.Lock()
	defer f.WriteGetObjectResponseRequestCall.mutex.Unlock()
	f.WriteGetObjectResponseRequestCall.CallCount++
	f.WriteGetObjectResponseRequestCall.Receives.WriteGetObjectResponseInput = param1
	if f.WriteGetObjectResponseRequestCall.Stub != nil {
		return f.WriteGetObjectResponseRequestCall.Stub(param1)
	}
	return f.WriteGetObjectResponseRequestCall.Returns.Request, f.WriteGetObjectResponseRequestCall.Returns.WriteGetObjectResponseOutput
}
func (f *S3API) WriteGetObjectResponseWithContext(param1 context.Context, param2 *s3.WriteGetObjectResponseInput, param3 ...request.Option) (*s3.WriteGetObjectResponseOutput, error) {
	f.WriteGetObjectResponseWithContextCall.mutex.Lock()
	defer f.WriteGetObjectResponseWithContextCall.mutex.Unlock()
	f.WriteGetObjectResponseWithContextCall.CallCount++
	f.WriteGetObjectResponseWithContextCall.Receives.Context = param1
	f.WriteGetObjectResponseWithContextCall.Receives.WriteGetObjectResponseInput = param2
	f.WriteGetObjectResponseWithContextCall.Receives.OptionSlice = param3
	if f.WriteGetObjectResponseWithContextCall.Stub != nil {
		return f.WriteGetObjectResponseWithContextCall.Stub(param1, param2, param3...)
	}
	return f.WriteGetObjectResponseWithContextCall.Returns.WriteGetObjectResponseOutput, f.WriteGetObjectResponseWithContextCall.Returns.Error
}
