<template>
    <div class="iadmin_aboutus">
        <el-card class="ia_card" :body-style="{ padding: '0px' }">
        <div style="padding: 14px;">
            <div class="ia_textwrap">
                <img src="../assets/img/favicon-96x96.png" alt="Image">
                <div class="ia_text">
                    <p>Nlan Admin V1.0</p>
                    <a href="https://github.com/mynet1314" target="_blank">
                        <el-button type="text" class="button">@mynet1314</el-button>
                    </a>
                </div>
            </div>
            <p class="time">基于docker的多用户SSR & V2ray管理平台.</p>
        </div>
        </el-card>
    </div>
</template>

<script>
</script>

<style lang="less">
    .ia_card {
        width: 400px;
        margin: 50px auto;
    }
    .ia_img {
        width: 100%;
    }
    .ia_textwrap {
        display: flex;
        justify-content: space-around;
    }
    .ia_text {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
    }
</style>
