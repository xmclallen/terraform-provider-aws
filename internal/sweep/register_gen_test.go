// Code generated by internal/generate/sweeperregistration/main.go; DO NOT EDIT.

package sweep_test

import (
	"github.com/hashicorp/terraform-provider-aws/internal/service/accessanalyzer"
	"github.com/hashicorp/terraform-provider-aws/internal/service/acm"
	"github.com/hashicorp/terraform-provider-aws/internal/service/acmpca"
	"github.com/hashicorp/terraform-provider-aws/internal/service/amplify"
	"github.com/hashicorp/terraform-provider-aws/internal/service/apigateway"
	"github.com/hashicorp/terraform-provider-aws/internal/service/apigatewayv2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/appconfig"
	"github.com/hashicorp/terraform-provider-aws/internal/service/appflow"
	"github.com/hashicorp/terraform-provider-aws/internal/service/applicationinsights"
	"github.com/hashicorp/terraform-provider-aws/internal/service/appmesh"
	"github.com/hashicorp/terraform-provider-aws/internal/service/apprunner"
	"github.com/hashicorp/terraform-provider-aws/internal/service/appstream"
	"github.com/hashicorp/terraform-provider-aws/internal/service/appsync"
	"github.com/hashicorp/terraform-provider-aws/internal/service/athena"
	"github.com/hashicorp/terraform-provider-aws/internal/service/auditmanager"
	"github.com/hashicorp/terraform-provider-aws/internal/service/autoscaling"
	"github.com/hashicorp/terraform-provider-aws/internal/service/autoscalingplans"
	"github.com/hashicorp/terraform-provider-aws/internal/service/backup"
	"github.com/hashicorp/terraform-provider-aws/internal/service/batch"
	"github.com/hashicorp/terraform-provider-aws/internal/service/bcmdataexports"
	"github.com/hashicorp/terraform-provider-aws/internal/service/budgets"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloud9"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloudformation"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloudfront"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloudhsmv2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloudsearch"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloudtrail"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloudwatch"
	"github.com/hashicorp/terraform-provider-aws/internal/service/codeartifact"
	"github.com/hashicorp/terraform-provider-aws/internal/service/codebuild"
	"github.com/hashicorp/terraform-provider-aws/internal/service/codegurureviewer"
	"github.com/hashicorp/terraform-provider-aws/internal/service/codepipeline"
	"github.com/hashicorp/terraform-provider-aws/internal/service/codestarconnections"
	"github.com/hashicorp/terraform-provider-aws/internal/service/codestarnotifications"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cognitoidentity"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cognitoidp"
	"github.com/hashicorp/terraform-provider-aws/internal/service/configservice"
	"github.com/hashicorp/terraform-provider-aws/internal/service/connect"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cur"
	"github.com/hashicorp/terraform-provider-aws/internal/service/dataexchange"
	"github.com/hashicorp/terraform-provider-aws/internal/service/datasync"
	"github.com/hashicorp/terraform-provider-aws/internal/service/dax"
	"github.com/hashicorp/terraform-provider-aws/internal/service/deploy"
	"github.com/hashicorp/terraform-provider-aws/internal/service/devicefarm"
	"github.com/hashicorp/terraform-provider-aws/internal/service/directconnect"
	"github.com/hashicorp/terraform-provider-aws/internal/service/dlm"
	"github.com/hashicorp/terraform-provider-aws/internal/service/dms"
	"github.com/hashicorp/terraform-provider-aws/internal/service/docdb"
	"github.com/hashicorp/terraform-provider-aws/internal/service/docdbelastic"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ds"
	"github.com/hashicorp/terraform-provider-aws/internal/service/dynamodb"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ecr"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ecrpublic"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ecs"
	"github.com/hashicorp/terraform-provider-aws/internal/service/efs"
	"github.com/hashicorp/terraform-provider-aws/internal/service/eks"
	"github.com/hashicorp/terraform-provider-aws/internal/service/elasticache"
	"github.com/hashicorp/terraform-provider-aws/internal/service/elasticbeanstalk"
	"github.com/hashicorp/terraform-provider-aws/internal/service/elasticsearch"
	"github.com/hashicorp/terraform-provider-aws/internal/service/elb"
	"github.com/hashicorp/terraform-provider-aws/internal/service/elbv2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/emr"
	"github.com/hashicorp/terraform-provider-aws/internal/service/emrcontainers"
	"github.com/hashicorp/terraform-provider-aws/internal/service/emrserverless"
	"github.com/hashicorp/terraform-provider-aws/internal/service/events"
	"github.com/hashicorp/terraform-provider-aws/internal/service/evidently"
	"github.com/hashicorp/terraform-provider-aws/internal/service/finspace"
	"github.com/hashicorp/terraform-provider-aws/internal/service/firehose"
	"github.com/hashicorp/terraform-provider-aws/internal/service/fis"
	"github.com/hashicorp/terraform-provider-aws/internal/service/fsx"
	"github.com/hashicorp/terraform-provider-aws/internal/service/gamelift"
	"github.com/hashicorp/terraform-provider-aws/internal/service/glacier"
	"github.com/hashicorp/terraform-provider-aws/internal/service/globalaccelerator"
	"github.com/hashicorp/terraform-provider-aws/internal/service/glue"
	"github.com/hashicorp/terraform-provider-aws/internal/service/grafana"
	"github.com/hashicorp/terraform-provider-aws/internal/service/guardduty"
	"github.com/hashicorp/terraform-provider-aws/internal/service/iam"
	"github.com/hashicorp/terraform-provider-aws/internal/service/imagebuilder"
	"github.com/hashicorp/terraform-provider-aws/internal/service/internetmonitor"
	"github.com/hashicorp/terraform-provider-aws/internal/service/iot"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kafka"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kafkaconnect"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kendra"
	"github.com/hashicorp/terraform-provider-aws/internal/service/keyspaces"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kinesis"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kinesisanalytics"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kinesisanalyticsv2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kms"
	"github.com/hashicorp/terraform-provider-aws/internal/service/lakeformation"
	"github.com/hashicorp/terraform-provider-aws/internal/service/lambda"
	"github.com/hashicorp/terraform-provider-aws/internal/service/lexmodels"
	"github.com/hashicorp/terraform-provider-aws/internal/service/lexv2models"
	"github.com/hashicorp/terraform-provider-aws/internal/service/licensemanager"
	"github.com/hashicorp/terraform-provider-aws/internal/service/lightsail"
	"github.com/hashicorp/terraform-provider-aws/internal/service/location"
	"github.com/hashicorp/terraform-provider-aws/internal/service/logs"
	"github.com/hashicorp/terraform-provider-aws/internal/service/medialive"
	"github.com/hashicorp/terraform-provider-aws/internal/service/mediapackage"
	"github.com/hashicorp/terraform-provider-aws/internal/service/memorydb"
	"github.com/hashicorp/terraform-provider-aws/internal/service/mq"
	"github.com/hashicorp/terraform-provider-aws/internal/service/mwaa"
	"github.com/hashicorp/terraform-provider-aws/internal/service/neptune"
	"github.com/hashicorp/terraform-provider-aws/internal/service/networkfirewall"
	"github.com/hashicorp/terraform-provider-aws/internal/service/networkmanager"
	"github.com/hashicorp/terraform-provider-aws/internal/service/opensearch"
	"github.com/hashicorp/terraform-provider-aws/internal/service/opensearchserverless"
	"github.com/hashicorp/terraform-provider-aws/internal/service/opsworks"
	"github.com/hashicorp/terraform-provider-aws/internal/service/pinpoint"
	"github.com/hashicorp/terraform-provider-aws/internal/service/pipes"
	"github.com/hashicorp/terraform-provider-aws/internal/service/qldb"
	"github.com/hashicorp/terraform-provider-aws/internal/service/quicksight"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ram"
	"github.com/hashicorp/terraform-provider-aws/internal/service/rds"
	"github.com/hashicorp/terraform-provider-aws/internal/service/redshift"
	"github.com/hashicorp/terraform-provider-aws/internal/service/redshiftserverless"
	"github.com/hashicorp/terraform-provider-aws/internal/service/resourceexplorer2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/resourcegroups"
	"github.com/hashicorp/terraform-provider-aws/internal/service/route53"
	"github.com/hashicorp/terraform-provider-aws/internal/service/route53recoverycontrolconfig"
	"github.com/hashicorp/terraform-provider-aws/internal/service/route53resolver"
	"github.com/hashicorp/terraform-provider-aws/internal/service/rum"
	"github.com/hashicorp/terraform-provider-aws/internal/service/s3"
	"github.com/hashicorp/terraform-provider-aws/internal/service/s3control"
	"github.com/hashicorp/terraform-provider-aws/internal/service/sagemaker"
	"github.com/hashicorp/terraform-provider-aws/internal/service/scheduler"
	"github.com/hashicorp/terraform-provider-aws/internal/service/schemas"
	"github.com/hashicorp/terraform-provider-aws/internal/service/secretsmanager"
	"github.com/hashicorp/terraform-provider-aws/internal/service/servicecatalog"
	"github.com/hashicorp/terraform-provider-aws/internal/service/servicediscovery"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ses"
	"github.com/hashicorp/terraform-provider-aws/internal/service/sesv2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/sfn"
	"github.com/hashicorp/terraform-provider-aws/internal/service/shield"
	"github.com/hashicorp/terraform-provider-aws/internal/service/signer"
	"github.com/hashicorp/terraform-provider-aws/internal/service/simpledb"
	"github.com/hashicorp/terraform-provider-aws/internal/service/sns"
	"github.com/hashicorp/terraform-provider-aws/internal/service/sqs"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ssm"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ssmcontacts"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ssmincidents"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ssoadmin"
	"github.com/hashicorp/terraform-provider-aws/internal/service/storagegateway"
	"github.com/hashicorp/terraform-provider-aws/internal/service/swf"
	"github.com/hashicorp/terraform-provider-aws/internal/service/synthetics"
	"github.com/hashicorp/terraform-provider-aws/internal/service/timestreamwrite"
	"github.com/hashicorp/terraform-provider-aws/internal/service/transcribe"
	"github.com/hashicorp/terraform-provider-aws/internal/service/transfer"
	"github.com/hashicorp/terraform-provider-aws/internal/service/verifiedpermissions"
	"github.com/hashicorp/terraform-provider-aws/internal/service/vpclattice"
	"github.com/hashicorp/terraform-provider-aws/internal/service/waf"
	"github.com/hashicorp/terraform-provider-aws/internal/service/wafregional"
	"github.com/hashicorp/terraform-provider-aws/internal/service/wafv2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/workspaces"
	"github.com/hashicorp/terraform-provider-aws/internal/service/xray"
)

func registerSweepers() {
	accessanalyzer.RegisterSweepers()
	acm.RegisterSweepers()
	acmpca.RegisterSweepers()
	amplify.RegisterSweepers()
	apigateway.RegisterSweepers()
	apigatewayv2.RegisterSweepers()
	appconfig.RegisterSweepers()
	appflow.RegisterSweepers()
	applicationinsights.RegisterSweepers()
	appmesh.RegisterSweepers()
	apprunner.RegisterSweepers()
	appstream.RegisterSweepers()
	appsync.RegisterSweepers()
	athena.RegisterSweepers()
	auditmanager.RegisterSweepers()
	autoscaling.RegisterSweepers()
	autoscalingplans.RegisterSweepers()
	backup.RegisterSweepers()
	batch.RegisterSweepers()
	bcmdataexports.RegisterSweepers()
	budgets.RegisterSweepers()
	cloud9.RegisterSweepers()
	cloudformation.RegisterSweepers()
	cloudfront.RegisterSweepers()
	cloudhsmv2.RegisterSweepers()
	cloudsearch.RegisterSweepers()
	cloudtrail.RegisterSweepers()
	cloudwatch.RegisterSweepers()
	codeartifact.RegisterSweepers()
	codebuild.RegisterSweepers()
	codegurureviewer.RegisterSweepers()
	codepipeline.RegisterSweepers()
	codestarconnections.RegisterSweepers()
	codestarnotifications.RegisterSweepers()
	cognitoidentity.RegisterSweepers()
	cognitoidp.RegisterSweepers()
	configservice.RegisterSweepers()
	connect.RegisterSweepers()
	cur.RegisterSweepers()
	dataexchange.RegisterSweepers()
	datasync.RegisterSweepers()
	dax.RegisterSweepers()
	deploy.RegisterSweepers()
	devicefarm.RegisterSweepers()
	directconnect.RegisterSweepers()
	dlm.RegisterSweepers()
	dms.RegisterSweepers()
	docdb.RegisterSweepers()
	docdbelastic.RegisterSweepers()
	ds.RegisterSweepers()
	dynamodb.RegisterSweepers()
	ec2.RegisterSweepers()
	ecr.RegisterSweepers()
	ecrpublic.RegisterSweepers()
	ecs.RegisterSweepers()
	efs.RegisterSweepers()
	eks.RegisterSweepers()
	elasticache.RegisterSweepers()
	elasticbeanstalk.RegisterSweepers()
	elasticsearch.RegisterSweepers()
	elb.RegisterSweepers()
	elbv2.RegisterSweepers()
	emr.RegisterSweepers()
	emrcontainers.RegisterSweepers()
	emrserverless.RegisterSweepers()
	events.RegisterSweepers()
	evidently.RegisterSweepers()
	finspace.RegisterSweepers()
	firehose.RegisterSweepers()
	fis.RegisterSweepers()
	fsx.RegisterSweepers()
	gamelift.RegisterSweepers()
	glacier.RegisterSweepers()
	globalaccelerator.RegisterSweepers()
	glue.RegisterSweepers()
	grafana.RegisterSweepers()
	guardduty.RegisterSweepers()
	iam.RegisterSweepers()
	imagebuilder.RegisterSweepers()
	internetmonitor.RegisterSweepers()
	iot.RegisterSweepers()
	kafka.RegisterSweepers()
	kafkaconnect.RegisterSweepers()
	kendra.RegisterSweepers()
	keyspaces.RegisterSweepers()
	kinesis.RegisterSweepers()
	kinesisanalytics.RegisterSweepers()
	kinesisanalyticsv2.RegisterSweepers()
	kms.RegisterSweepers()
	lakeformation.RegisterSweepers()
	lambda.RegisterSweepers()
	lexmodels.RegisterSweepers()
	lexv2models.RegisterSweepers()
	licensemanager.RegisterSweepers()
	lightsail.RegisterSweepers()
	location.RegisterSweepers()
	logs.RegisterSweepers()
	medialive.RegisterSweepers()
	mediapackage.RegisterSweepers()
	memorydb.RegisterSweepers()
	mq.RegisterSweepers()
	mwaa.RegisterSweepers()
	neptune.RegisterSweepers()
	networkfirewall.RegisterSweepers()
	networkmanager.RegisterSweepers()
	opensearch.RegisterSweepers()
	opensearchserverless.RegisterSweepers()
	opsworks.RegisterSweepers()
	pinpoint.RegisterSweepers()
	pipes.RegisterSweepers()
	qldb.RegisterSweepers()
	quicksight.RegisterSweepers()
	ram.RegisterSweepers()
	rds.RegisterSweepers()
	redshift.RegisterSweepers()
	redshiftserverless.RegisterSweepers()
	resourceexplorer2.RegisterSweepers()
	resourcegroups.RegisterSweepers()
	route53.RegisterSweepers()
	route53recoverycontrolconfig.RegisterSweepers()
	route53resolver.RegisterSweepers()
	rum.RegisterSweepers()
	s3.RegisterSweepers()
	s3control.RegisterSweepers()
	sagemaker.RegisterSweepers()
	scheduler.RegisterSweepers()
	schemas.RegisterSweepers()
	secretsmanager.RegisterSweepers()
	servicecatalog.RegisterSweepers()
	servicediscovery.RegisterSweepers()
	ses.RegisterSweepers()
	sesv2.RegisterSweepers()
	sfn.RegisterSweepers()
	shield.RegisterSweepers()
	signer.RegisterSweepers()
	simpledb.RegisterSweepers()
	sns.RegisterSweepers()
	sqs.RegisterSweepers()
	ssm.RegisterSweepers()
	ssmcontacts.RegisterSweepers()
	ssmincidents.RegisterSweepers()
	ssoadmin.RegisterSweepers()
	storagegateway.RegisterSweepers()
	swf.RegisterSweepers()
	synthetics.RegisterSweepers()
	timestreamwrite.RegisterSweepers()
	transcribe.RegisterSweepers()
	transfer.RegisterSweepers()
	verifiedpermissions.RegisterSweepers()
	vpclattice.RegisterSweepers()
	waf.RegisterSweepers()
	wafregional.RegisterSweepers()
	wafv2.RegisterSweepers()
	workspaces.RegisterSweepers()
	xray.RegisterSweepers()
}
