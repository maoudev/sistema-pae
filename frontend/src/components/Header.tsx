import React, { useState } from "react";
import {Button, IconButton} from "@mui/material";
import Paper from "@mui/material/Paper";
import InputBase from "@mui/material/InputBase";
import AddCircleOutlineIcon from '@mui/icons-material/AddCircleOutline';
import SearchOutlinedIcon from "@mui/icons-material/SearchOutlined";
import DownloadXlsxButton from "./DownloadXlsxButton";
import ButtonNoBenefeciario from "./ButtonNoBenefeciario";
import {Link} from "react-router-dom";

interface HeaderProps {
  setData: Function;
  handleNoBeneficiario: Function;
}

const Header: React.FC<HeaderProps> = ({ setData, handleNoBeneficiario }) => {
  const [rut, setRut] = useState("");

  const sendData = (value: string) => {

    const txtrut = document.getElementById("txtrut") as HTMLInputElement;
    txtrut.disabled = true;
    txtrut.disabled = false;


    setData(value);
    setRut("");
  };

  return (
    <header className="flex w-full items-center justify-center py-5 outline-1 outline outline-cyan-950">
      <nav className="w-full">
        <ul className="flex flex-row items-center justify-evenly">
          <li>
            {
              <Paper
                sx={{
                  p: "2px",
                  display: "flex",
                  alignItems: "center",
                  width: 200,
                  height: 40,
                  backgroundColor: "transparent",
                  border: "1px solid #166bcc",
                }}
              >
                <InputBase
                    id={"txtrut"}
                  sx={{ ml: 1, flex: 1, color: "white" }}
                  placeholder="Buscar por rut"
                  inputProps={{ "aria-label": "Buscador por Rut" }}
                  value={rut}
                  onChange={(e) => setRut(e.target.value)}
                  onKeyDown={(e) => {
                    if (e.key === "Enter") sendData(rut);
                  }} // Add this line to handle keydown event
                />
                <IconButton
                  type="button"
                  sx={{ p: "10px" }}
                  aria-label="search"
                  onClick={() => sendData(rut)}
                >
                  <SearchOutlinedIcon color="primary" />
                </IconButton>
              </Paper>
            }
          </li>
          <li>
            <DownloadXlsxButton />
          </li>
          <li>
            <ButtonNoBenefeciario handleNoBeneficiario={handleNoBeneficiario} />
          </li>
          <li>
            <Button
                variant={"outlined"}
                startIcon={<AddCircleOutlineIcon/>}
                component={Link}
                to={"/add"}
            >Registrar Estudiante</Button>
          </li>
        </ul>
      </nav>
    </header>
  );
};

export default Header;
