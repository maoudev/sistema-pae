import alumnoAddRequest from "../types/alumnoAddRequest";

export const addStudent = async (alumno: alumnoAddRequest) : Promise<Response> => {
    const response = await fetch(`http://localhost:3000/student`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(alumno),
    })

    return response;
};