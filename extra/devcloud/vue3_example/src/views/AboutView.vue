<template>
  <div class="about">
    <h2 id="name">{{ person }}</h2>
    <input v-model="person.name" type="text" />
    <input v-model="skill" @keyup.enter="addSkile(skill)" type="text" />
  </div>
</template>

<script setup>
import { shallowReactive, ref } from "vue";

let skill = ref("");

// 使用ref来为基础类型 构造响应式变量
let person = shallowReactive({
  name: "张三",
  profile: { city: "北京" },
  skills: ["Golang", "Vue"],
});

let addSkile = (s) => {
  person.skills.push(s);
  person.profile.skill_count = person.skills.length;
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
