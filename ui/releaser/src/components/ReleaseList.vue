<script setup lang="ts">
import { ref, reactive } from 'vue'
import type { FormInstance } from 'element-plus'

interface Project {
    name: string
    address: string
    branch: string
}

interface Artifact {
}

interface Release {
    name: string
    repositories: Project[]
}

const releaseList = ref([] as Release[])
function load() {
    fetch("/v1/releases").then(function(r: any) {
        return r.json();
    }).then(function(d: any) {
        releaseList.value = d;
    })
}
load()

// creation
const creationDialog = ref(false)
const formRelease = reactive({
    repositories: [{} as Project]
} as Release)
const formReleaseRef = ref<FormInstance>()
const save = async (formEl: FormInstance | undefined) => {
    console.log(formRelease)
    fetch('/v1/releases', {
        method: "POST",
        body: JSON.stringify(formRelease)
    })
}
</script>

<template>
    <div>
        <el-button type="primary" @click="creationDialog=true">Create</el-button>
        <el-button type="primary" @click="load()">Refresh</el-button>
    </div>
    <el-table :data="releaseList" style="width: 100%">
        <el-table-column prop="name" label="Name" width="100" />
        <el-table-column label="Repository" width="200">
            <template #default="scope">
                <div v-if="scope.row.repositories.length === 1">
                    <a :href="scope.row.repositories[0].address" target="_blank">
                        {{ scope.row.repositories[0].name }}
                        <i class="gg-git-branch"></i>
                        {{ scope.row.repositories[0].branch }}
                    </a>
                </div>
            </template>
        </el-table-column>
        <el-table-column label="Image" width="200">
            <template #default="scope">
                <div v-if="scope.row.status && scope.row.status.artifacts && scope.row.status.artifacts.length === 1">
                    {{ scope.row.status.artifacts[0].image.image }}:{{ scope.row.status.artifacts[0].image.tag }}
                </div>
            </template>
        </el-table-column>
        <el-table-column prop="status.completionTime" label="Completion" width="200" />
        <el-table-column prop="phase" label="Phase" />
    </el-table>

    <el-drawer v-model="creationDialog">
        <template #header>
            <h4>Create a new release</h4>
        </template>
        <template #default>
            <el-form
                ref="formReleaseRef"
                label-width="auto"
                :model="formRelease"
                style="max-width: 600px"
            >
                <el-form-item label="Name">
                    <el-input v-model="formRelease.name" />
                </el-form-item>
                <el-form-item label="Repository">
                    <el-input v-model="formRelease.repositories[0].address" />
                </el-form-item>
                <el-form-item label="Branch">
                    <el-input v-model="formRelease.repositories[0].branch" />
                </el-form-item>
            </el-form>
        </template>
        <template #footer>
            <el-button type="primary" @click="save(formReleaseRef)">Save</el-button>
        </template>
    </el-drawer>
</template>

<style scoped>
.gg-git-branch {
    box-sizing: border-box;
    position: relative;
    display: block;
    transform: scale(var(--ggs,1));
    width: 2px;
    height: 14px;
    background: currentColor
}

.gg-git-branch::after,
.gg-git-branch::before {
    content: "";
    display: block;
    box-sizing: border-box;
    position: absolute
}

.gg-git-branch::before {
    border-right: 2px solid;
    border-bottom: 2px solid;
    border-bottom-right-radius: 4px;
    bottom: 0;
    width: 8px;
    height: 6px;
    left: 0
}

.gg-git-branch::after {
    width: 4px;
    height: 4px;
    background: currentColor;
    box-shadow:
    0 12px 0 0,
    6px 6px 0 0;
    border-radius: 100%;
    left: -1px;
    top: -1px
} 
</style>
