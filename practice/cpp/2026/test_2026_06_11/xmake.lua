add_rules("mode.debug", "mode.release")

target("test_2026_06_11")
    set_kind("binary")
    add_files("src/*.cpp")
    set_languages("c++20")
    set_warnings("all", "error")
