import axios from "axios";

export const instance = axios.create({
    baseURL: "http://45.141.101.31:8001/v1",
});

//http://45.141.101.31:8001/v1
/*PORT=8001
PDF=C:/Програмирование/Проекты/calculation-estimate/pdf/
PATH_SOURCES=C:/Програмирование/Проекты/calculation-estimate/sources/
*/