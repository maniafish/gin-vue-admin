package request

import "gin-vue-admin/model"

type ProjectASearch struct{
    model.ProjectA
    PageInfo
}