<template>
  <div class="about">
    <h2 id="name">{{ person }}</h2>
    <input v-model="name.value" type="text" />
    <input v-model="skill" @keyup.enter="addSkile(skill)" type="text" />
  </div>
</template>

<script setup>
import { ref } from "vue";

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
