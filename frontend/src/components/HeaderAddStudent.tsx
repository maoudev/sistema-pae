
import {Link} from "react-router-dom";
import HomeRoundedIcon from '@mui/icons-material/HomeRounded';
import React from "react";
import {Button} from "@mui/material";

const HeaderAddStudent = () => {

    return (
        <header className="flex w-full items-center justify-center py-5 outline-1 outline outline-cyan-950">
            <nav className="w-full">
                <ul className="flex flex-row items-center justify-evenly">
                    <li>
                        <Button
                            variant={"outlined"}
                            startIcon={<HomeRoundedIcon />}
                            component={Link}
                            to={"/"}
                        >Inicio</Button>
                    </li>
                </ul>
            </nav>
        </header>
    )
};

export default HeaderAddStudent;