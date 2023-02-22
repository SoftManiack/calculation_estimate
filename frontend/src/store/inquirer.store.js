import { getInquirer, getEstimate } from '@/services/inquirer.service';

const mutations = {
    SET_INQUIRER(state, data) {
        state.inquirer = data
    },
    SET_ESTIMATES(state, data) {
        state.estimates = data
    },
    SET_ERROR(state, error){
        state.error = error;
    },
}

const actions = {
    async getInquirerAction( { commit } ) {
        const data = await getInquirer();
        if(data){
            commit("SET_INQUIRER", data );
        }else{
            commit("INQUIRER_ERROR", "Нет данных");
        }
    },
    async getEstimateAction( { commit }, inquirer){
        const data = await getEstimate(inquirer);
        console.log("*")
        console.log(data)
        if(data){
            commit("SET_ESTIMATES", data );
        }else{
            commit("SET_ERROR", "Нет данных");
        }
    }
}

const getters = {
    getInquirerGetter: (state) => state.inquirer,
    getGroupsInquirerGetter: (state) => state.inquirer.groups,
    getEstimateGetter: (state) => state.estimates,
    errorInquirer: (state) => state.error,
}

const state = () => ({
    error: "",
    inquirer: {
        uid: "",
        groups: [
            {
                uid: "",
                title: "",
                questions: [
                    {
                        uid: "",
                        text: "",
                        options: [
                            "",
                        ],
                        answer: "",
                        subquery: {
                            uid: "",
                            text: "",
                            options: [
                                "",
                            ],
                            answer: "",
                            conditions: "да",
                            subquery: null
                        }
                    },
                ]
            }
        ]
    },
    estimates: [
        { 
            uid: "",
            name: "" 
        }
    ]
})

export default {
    mutations,
    getters,
    actions,
    state,
}
