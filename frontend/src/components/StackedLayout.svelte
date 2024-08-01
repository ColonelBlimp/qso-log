<script>
    import {mainMenuStore} from "../lib/stores.js";
    import MainMenuBurger from "./MainMenuBurger.svelte";

    let isMainMenuOpen;
    mainMenuStore.subscribe((value) => {
        isMainMenuOpen = value;
    });
</script>
<div class="min-h-full">
    <nav class="bg-indigo-600">
        <div class="mx-auto max-w-7xl px-8">
            <div class="flex h-16 items-center justify-between">
                <div class="flex items-center">
                    <div class="block">
                        <div class="ml-10 flex items-baseline space-x-4">
                            <a href="#" class="rounded-md bg-indigo-700 px-3 py-2 text-sm font-medium text-white" aria-current="page">Logbook</a>
                            <a href="#" class="rounded-md px-3 py-2 text-sm font-medium text-white hover:bg-indigo-500 hover:bg-opacity-75">Rig Control</a>
                            <a href="#" class="rounded-md px-3 py-2 text-sm font-medium text-white hover:bg-indigo-500 hover:bg-opacity-75">Rotator</a>
                        </div>
                    </div>
                </div>
                <div class="-m-2 flex">
<MainMenuBurger/>
                </div>
            </div>
        </div>
        {#if isMainMenuOpen}
        <div class="" id="mainMenu">
        </div>
        {/if}
    </nav>
</div>
