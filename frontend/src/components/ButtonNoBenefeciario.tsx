import { Button } from "@mui/material";
import NoAccountsIcon from "@mui/icons-material/NoAccounts";
import React from "react";

interface ButtonNoBenefeciarioProps {
  handleNoBeneficiario: Function;
}

const ButtonNoBenefeciario: React.FC<ButtonNoBenefeciarioProps> = ({
  handleNoBeneficiario,
}) => {
  return (
    <Button
      startIcon={<NoAccountsIcon />}
      variant="outlined"
      onClick={() => handleNoBeneficiario()}
    >
      No beneficiario
    </Button>
  );
};

export default ButtonNoBenefeciario;
