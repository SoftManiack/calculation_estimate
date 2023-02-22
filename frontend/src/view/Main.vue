<template>
    <div class="l-form">
        <div class="c-header">
            <h2>Опросник</h2>
        </div>
        <div class="c-form" v-for="(item, key) in getGroupsInquirerGetter" :key="key">
            <div class="c-form_title">
                <h3>{{item.title}}</h3>
            </div>
            <div class="c-form_body">
                <div v-for="(item, key) in item.questions" :key="key">
                    <h5>{{key + 1}}. {{item.text}}</h5>
                    <template v-if="item.options.length > 0" >
                        <b-form-radio-group
                            :options="item.options"
                            class="mb-3"
                            v-model="item.answer"
                            value-field="item"
                            text-field="name"
                            disabled-field="notEnabled"
                            plain
                            stacked
                            :state="!!item.answer"
                        ></b-form-radio-group>
                        <div class="px-2" v-if="item.subquery.uid != '' & item.subquery.conditions == item.answer">
                            <h5>{{item.text}}</h5>
                            <template v-if="item.subquery.options.length > 0" >
                                <b-form-radio-group
                                    :options="item.subquery.options"
                                    class="mb-3"
                                    v-model="item.subquery.answer"
                                    value-field="item"
                                    text-field="name"
                                    disabled-field="notEnabled"
                                    plain
                                    stacked
                                    :state="!!item.subquery.answer"
                                ></b-form-radio-group>
                            </template>
                            <template v-if="item.subquery.options.length == 0">
                                <b-form-input
                                    class="mb-3 c-form_number-input"
                                    type="number"
                                    v-model="item.subquery.answer"
                                    placeholder="0">
                                </b-form-input>
                            </template>
                        </div>
                    </template>
                    <template v-if="item.options.length == 0">
                        <b-form-input
                            class="mb-3 c-form_number-input"
                            type="number"
                            v-model="item.answer"
                            placeholder="0">
                        </b-form-input>
                    </template>
                </div>
            </div>
        </div>
        <div class="c-footer">  
            <b-button @click="getEstimates" variant="primary" >Получить смету</b-button>
            <hr>
            <div class="c-footer_files">
                <template v-if="getEstimateGetter[0].uid != ''">
                    <div v-for="(item, key) in getEstimateGetter" :key="key">
                        <a :href="`http://localhost/${item.uid}.pdf`" class="c-footer_file">
                            <img width="57" height="42" src="@/assets/icons/pdf.svg" />
                            <h5>{{item.name}}</h5>
                        </a>
                    </div>
                </template>
            </div>
        </div>
    </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
export default {
    name: "Main",
    data() {
      return {
        count: 0,
      }
    },
    computed: {
        ...mapGetters({
            getInquirerGetter: "getInquirerGetter",
            getGroupsInquirerGetter: "getGroupsInquirerGetter",
            getEstimateGetter: "getEstimateGetter"
        }),
    },
    methods: {
        ...mapActions({  
            getInquirerAction: "getInquirerAction",
            getEstimate: "getEstimateAction",
        }),
        async getEstimates(){
            await this.getEstimate(this.getInquirerGetter)
        }
    },
    async mounted() { 
        await this.getInquirerAction();
    },
}
</script>

<style>
</style>
