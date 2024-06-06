import { Button } from "@mui/material";
import DownloadIcon from "@mui/icons-material/Download";

function DownloadXlsxButton() {
  const handleClick = async () => {
    const response = await fetch(`http://localhost:3000/students`, {
      method: "GET",
    });

    let link = document.createElement("a");
    link.href = window.URL.createObjectURL(await response.blob());
    link.download = new Date().toLocaleString() + "_PAE.xlsx";
    link.click();

    window.URL.revokeObjectURL(link.href);
  };
  return (
    <Button
      variant="outlined"
      startIcon={<DownloadIcon />}
      onClick={async () => await handleClick()}
    >
      Descargar Ficha
    </Button>
  );
}

export default DownloadXlsxButton;
