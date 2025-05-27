package rpc_company

import (
	"context"

	pb "colinahealth_emr-api/internal/adapters/framework/left/company/grpc/pb"
	ct "colinahealth_emr-api/internal/application/core/domain/company"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddNewCompany gets the result of new company entry
func (grpca Adapter) AddNewCompany(ctx context.Context, req *pb.RequestVariables) (*pb.ResponseAnswers, error) {
	var err error
	ans := &pb.ResponseAnswers{}

	if req.GetName() == "" || req.GetWebsite() == "" || req.GetContact() == "" || req.GetEmail() == "" || req.GetCountry() == "" || req.GetState() == "" || req.GetZip() == "" {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetNewCompany(ct.Name(req.Name), ct.Contact(req.Contact), ct.Website(req.Website), ct.Email(req.Email), ct.Country(req.Country), ct.State(req.State), ct.Zip(req.Zip))
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.ResponseAnswers{
		Id: answer.CompanyID, Company: &pb.NewCompany{Uuid: answer.NewCompany.Uuid.String(), Name: string(answer.NewCompany.Name), Contact: string(answer.NewCompany.Contact), Website: string(answer.NewCompany.Website), Email: string(answer.NewCompany.Email), Country: string(answer.NewCompany.Country), State: string(answer.NewCompany.State), Zip: string(answer.NewCompany.Zip)},
	}

	return ans, nil
}