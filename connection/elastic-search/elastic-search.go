package elasticsearch

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
)

type ElasticSearch struct {
	//TODO: config
}

func (es *ElasticSearch) QueryRule(domainName string) (*elasticsearchservice.DescribeElasticsearchDomainOutput, error) {
	svc := elasticsearchservice.New(session.Must(session.NewSession()))

	// Specify the search query parameters
	searchInput := &elasticsearchservice.DescribeElasticsearchDomainInput{
		DomainName: aws.String(domainName),
		// ... specify other parameters as needed
	}

	// Execute the search query
	searchOutput, err := svc.DescribeElasticsearchDomain(searchInput)
	if err != nil {
		return nil, err
	}

	return searchOutput, nil
}
