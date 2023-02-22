import Vue from 'vue'
import Vuex from 'vuex'
import inquirer from "@/store/inquirer.store"
import { loadingStatuses } from "@/store/statusesLoadingConst"

Vue.use(Vuex);

export default new Vuex.Store({
    getters: {
        getLoadingStatus: (state) => state.loadingStatus,
    },
    state: () => ({
        loadingStatus: loadingStatuses.Empty,
    }),
    mutations: {
        UPDATE_STATUS(state, data){
            state.loadingStatus = data;
        },
    },
    modules: {
        inquirer,
    }
})
