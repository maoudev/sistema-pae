import {useState} from "react";
import { addStudent } from "../scripts/addStudent";
import alumnoAddRequest from "../types/alumnoAddRequest";
import {Alert, Button} from "@mui/material";

const Form = () => {

    const [rut, setRut] = useState("");
    const [nombre, setNombre] = useState("");
    const [nivel, setNivel] = useState("");
    const [letra, setLetra] = useState("");
    const [error, setError] = useState<boolean | null>(null);
    const [campos, setCampos] = useState<boolean | null>(null);

    const sendData = () => {
        setError(null);
        setCampos(null);
        const student: alumnoAddRequest = {
            rut: rut,
            nombre: nombre,
            nivel: Number(nivel),
            letra: letra
        };

        addStudent(student)
            .then((response) => {
                if (response.status === 200) {
                    setError(false);
                    clearFields();
                } else {
                    setError(true);
                    clearFields();
                }
            });
    };

    const clearFields = () => {
        setRut("");
        setNombre("");
        setNivel("");
        setLetra("");
    };

    const handleSubmit = () => {
        setError(null);
        setCampos(null);

        if (rut === "" || nombre === "" || nivel === "" || letra === "") {
            setCampos(true);
            return;
        }

        sendData();
    };

    return (
        <div className="w-2/3 h-fit mt-10 p-14 border absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 rounded-xl shadow-[0_20px_50px_rgba(8,_112,_184,_0.2)] border-[#1975d156] m-auto transition-shadow duration-300 ease-in-out">


            <h1 className={"text-3xl text-center font-bold mb-5"}>Registrar Estudiante</h1>

            <form className={"flex flex-col px-32 mb-4"}>
                <div className="mb-4">
                    <label className="block text-gray-400 text-sm font-bold mb-2" htmlFor="rut">
                        Rut
                    </label>
                    <input
                        value={rut}
                        type="text"
                        id="rut"
                        className="w-full px-3 py-2 border border-gray-600 rounded-md bg-gray-700 text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500"
                        placeholder="21784092-4 o 21784092"
                        onChange={(e) => setRut(e.target.value)}

                    />
                </div>
                <div className="mb-4">
                    <label className="block text-gray-400 text-sm font-bold mb-2" htmlFor="nombre">
                        Nombre
                    </label>
                    <input
                        value={nombre}
                        type="text"
                        id="nombre"
                        className="w-full px-3 py-2 border border-gray-600 rounded-md bg-gray-700 text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500"
                        placeholder="Paterno Materno Nombres"
                        onChange={(e) => setNombre(e.target.value)}
                    />
                </div>

                <div className="mb-4">
                    <label className="block text-gray-400 text-sm font-bold mb-2" htmlFor="nivel">
                        Nivel
                    </label>
                    <select
                        value={nivel}
                        id="nivel"
                        className="w-full px-3 py-2 border border-gray-600 rounded-md bg-gray-700 text-white focus:outline-none focus:ring-2 focus:ring-blue-500"
                        onChange={(e) => setNivel(e.target.value)}
                    >
                        <option value="" className="text-white">Selecciona una opción</option>
                        <option value="1" className="text-white">1°</option>
                        <option value="2" className="text-white">2°</option>
                        <option value="3" className="text-white">3°</option>
                        <option value="4" className="text-white">4°</option>
                    </select>
                </div>

                <div className="mb-4">
                    <label className="block text-gray-400 text-sm font-bold mb-2" htmlFor="letra">
                        Letra
                    </label>
                    <select
                        value={letra}
                        id="letra"
                        className="w-full px-3 py-2 border border-gray-600 rounded-md bg-gray-700 text-white focus:outline-none focus:ring-2 focus:ring-blue-500"
                        onChange={(e) => setLetra(e.target.value)}
                    >
                        <option value="" className="text-white">Selecciona una opción</option>
                        <option value="A" className="text-white">A</option>
                        <option value="B" className="text-white">B</option>
                        <option value="C" className="text-white">C</option>
                        <option value="D" className="text-white">D</option>
                        <option value="E" className="text-white">E</option>
                        <option value="F" className="text-white">F</option>
                        <option value="G" className="text-white">G</option>
                        <option value="H" className="text-white">H</option>
                        <option value="I" className="text-white">I</option>
                        <option value="J" className="text-white">J</option>
                    </select>
                </div>


                <Button
                    variant={"outlined"}
                    onClick={() => handleSubmit()}
                >
                    Enviar
                </Button>
            </form>

            {error != null ? (
                error ? <Alert  severity="error">Error al registrar alumno</Alert> : <Alert  severity="success">Alumno registrado correctamente</Alert>
            ) : null}

            {campos != null ? (
                campos ? <Alert  severity="warning">¡Por favor, complete todos los campos!</Alert>: null
            ) : null}
        </div>
    )
};

export default Form;