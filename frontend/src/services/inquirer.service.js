import { instance } from '@/services/instances.service';

export const getInquirer = async () => {
    try {
        const response = await instance.get(`/inquirer`);
        return response.data
    } catch (error) {
        return error
    }
}

export const getEstimate = async ( inquirer ) => {
    try {
        console.log(inquirer)
        const response = await instance.post(`/inquirer`, inquirer);
        return response.data
    } catch (error) {
        return error
    }
}