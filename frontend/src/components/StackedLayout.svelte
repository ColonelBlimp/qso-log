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

                <div class="block">
                    <div class="ml-4 flex items-center md:ml-6">
                        <!-- Profile dropdown -->
                        <div class="relative ml-3">
                            <div>
                                <MainMenuBurger/>
                            </div>

                            <!--
                              Dropdown menu, show/hide based on menu state.

                              Entering: "transition ease-out duration-100"
                                From: "transform opacity-0 scale-95"
                                To: "transform opacity-100 scale-100"
                              Leaving: "transition ease-in duration-75"
                                From: "transform opacity-100 scale-100"
                                To: "transform opacity-0 scale-95"
                            -->
                            {#if isMainMenuOpen}
                            <div class="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none" role="menu" aria-orientation="vertical" aria-labelledby="user-menu-button" tabindex="-1">
                                <!-- Active: "bg-gray-100", Not Active: "" -->
                                <a href="#" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-200 hover:bg-opacity-75" role="menuitem" tabindex="-1" id="user-menu-item-0">Your Profile</a>
                                <a href="#" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-200 hover:bg-opacity-75" role="menuitem" tabindex="-1" id="user-menu-item-1">Settings</a>
                                <a href="#" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-200 hover:bg-opacity-75" role="menuitem" tabindex="-1" id="user-menu-item-2">Sign out</a>
                            </div>
                            {/if}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </nav>
</div>
