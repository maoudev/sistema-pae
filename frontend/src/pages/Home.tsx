import { useEffect, useState } from "react";
import useScanDetection from "use-scan-detection";
import useSound from "use-sound";

import adelante from "./../assets/sounds/adelante.mp3";
import noBeneficiario from "./../assets/sounds/nobeneficiario.mp3";
import ingreso from "./../assets/sounds/ingreso.mp3";
import Header from "../components/Header";
import Alumno from "../types/alumno";

function Home() {
  const [datos, setDatos] = useState<string>("");
  const [error, setError] = useState<null | string>(null);
  const [requestError, setRequestError] = useState<null | string>(null);
  const [alumno, setAlumno] = useState<Alumno | null>(null);
  const [fecha, setFecha] = useState<string>("");

  const [play] = useSound(adelante);
  const [play2] = useSound(noBeneficiario);
  const [play3] = useSound(ingreso);

  const handleClick = (value: string) => {
    setDatos(value);
  };

  useEffect(() => {
    setAlumno(null);
    setDatos("");
    setRequestError(null);
    setError(null);
    if (datos !== "") {
      const getData = async () => {
        try {
          if (/^[0-9]+[-|‐]{1}[0-9kK]{1}$/.test(datos)) {
            const response = await fetch(
              `http://localhost:3000/student/search/${datos}`,
              {
                method: "GET",
              }
            );

            if (response.status === 200) {
              const alum = (await response.json()) as Alumno;

              setAlumno(alum);
              setDatos("");
              setRequestError(null);
              setError(null);
            }

            if (response.status === 404) {
              setRequestError("Error 404");
            }

            if (response.status === 500) {
              setRequestError("Error: 500");
            }
          }

          if (/^\d{8}$/.test(datos)) {
            const response = await fetch(
              `http://localhost:3000/student/scanner/${datos}`,
              {
                method: "GET",
              }
            );

            if (response.status === 200) {
              const alum = (await response.json()) as Alumno;

              setAlumno(alum);
              setDatos("");
              setRequestError(null);
              setError(null);
            }

            if (response.status === 404) {
              setRequestError("Error 404");
            }

            if (response.status === 500) {
              setRequestError("Error: 500");
            }
          }
        } catch (error) {
          setError("Error de Conexión");
        }
      };

      setFecha(new Date().toLocaleString());
      getData();
    }
  }, [datos]);

  useEffect(() => {
    if (requestError != null) {
      play2();
    }

    if (requestError === null && alumno !== null) {
      if (alumno.almorzo) {
        play3();
      } else {
        play();
      }
    }
  }, [alumno]);

  useScanDetection({
    onComplete: (data) => {
      setDatos(data.toString());
    },
    onError: (error) => {
      setError(error.toString());
    },
  });

  const handleNoBeneficiario = () => {
    const insertData = async () => {
      const response = await fetch(
        `http://localhost:3000/student/non-beneficiary`,
        {
          method: "POST",
        }
      );

      if (response.status === 200) {
        play();
      }
    };

    insertData();
  };

  return (
    <div>
      <Header
        setData={handleClick}
        handleNoBeneficiario={handleNoBeneficiario}
      />

      <div className="w-2/3 h-fit mt-10 px-20 py-20 border absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 rounded-xl shadow-[0_20px_50px_rgba(8,_112,_184,_0.2)] border-[#1975d156] m-auto transition-shadow duration-300 ease-in-out">
        <h2 className="text-3xl text-center font-bold">
          {alumno ? alumno.nombre.toUpperCase() : "SISTEMA PAE"}
        </h2>
        <p className="text-xl mt-16 text-center">
          Rut: {alumno ? alumno.rut : ""}
        </p>
        <p className="text-xl mt-4 text-center">
          Curso: {alumno ? alumno.nivel + "°" + alumno.letra : ""}
        </p>
        <p className="text-xl mt-4 mb-16 text-center">Fecha: {fecha}</p>

        {requestError != null || error != null ? (
          <p className="text-2xl text-red-500 font-bold mt-4 text-center">
            No Se Encontró Al Alumno
          </p>
        ) : alumno ? (
          <p
            className={
              alumno.almorzo
                ? "text-2xl text-red-500 font-bold mt-4 text-center"
                : "text-2xl text-green-500 font-bold mt-4 text-center"
            }
          >
            {alumno.almorzo
              ? "El Alumno Ya Hizo Ingreso Al Comedor"
              : "El Alumno Puede Ingresar Al Comedor"}
          </p>
        ) : (
          <p className="text-2xl font-bold mt-4 text-center">
            Esperando Lectura
          </p>
        )}
      </div>
    </div>
  );
}

export default Home;
