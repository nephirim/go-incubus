package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"strings"
)

const ProposalTypeUpdateFees = "UpdateFantokenFeesProposal"

func init() {
	govtypes.RegisterProposalType(ProposalTypeUpdateFees)
	govtypes.RegisterProposalTypeCodec(&UpdateFeesProposal{}, "go-incubus/fantoken/UpdateFeesProposal")
}

var _ govtypes.Content = &UpdateFeesProposal{}

func NewUpdateFeesProposal(title, description string, issueFee, mintFee, burnFee sdk.Coin) govtypes.Content {
	return &UpdateFeesProposal{
		Title:       title,
		Description: description,
		IssueFee:    issueFee,
		MintFee:     mintFee,
		BurnFee:     burnFee,
	}
}

func (p *UpdateFeesProposal) GetTitle() string { return p.Title }

func (p *UpdateFeesProposal) GetDescription() string { return p.Description }

func (p *UpdateFeesProposal) ProposalRoute() string { return RouterKey }

func (p *UpdateFeesProposal) ProposalType() string { return ProposalTypeUpdateFees }

func (p *UpdateFeesProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	if err := ValidateFees(p.IssueFee, p.MintFee, p.BurnFee); err != nil {
		return err
	}

	return nil
}

func (p UpdateFeesProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`Update Fantoken Fees Proposal:
  Title:       %s
  Description: %s
  Issue Fee:   %s
  Mint Fee:    %s
  Burn Fee:    %s
`, p.Title, p.Description, p.IssueFee, p.MintFee, p.BurnFee))
	return b.String()
}
