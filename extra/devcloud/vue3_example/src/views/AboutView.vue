<template>
  <div class="about">
    <h2 id="name">{{ person }}</h2>
    <div>
      <input v-model="name" type="text" />
      <input v-model="skill" @keyup.enter="addSkile(skill)" type="text" />
    </div>
    <div>
      <ButtonCounter style="width: 220px" />
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import ButtonCounter from "@/components/ButtonCounter.vue";

import { inject } from "vue";

// 这里也可以获取默认值: inject(<变量名称>, <变量默认值>), 如果获取不到变量 就使用默认值
const count = inject("count");

let skill = ref("");

// 使用ref来构造一个对象
let person = {
  name: ref("张三"),
  profile: ref({ city: "北京" }),
  skills: ref(["Golang", "Vue"]),
};

// 解构赋值
let { name, profile, skills } = person;

let addSkile = (s) => {
  skills.value.push(s);
  profile.skill_count = skills.value.length;
  console.log(count.value);
};
</script>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
