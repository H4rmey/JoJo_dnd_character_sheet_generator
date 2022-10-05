#!/bin/sh

sed '/^$/d' feats_raw | awk '
{
    line=$0
    if (line ~ "//") {
        getline
        print "- name:", "\"",$0,"\""
        getline
        if ( $0 ~ "Prerequisite") {
            print " ", $0
            getline
        }
        if ( $0 ~ "Description" ) {
            getline
        }
        print "  description:", "\"",$0,"\""
        getline
        print "  effects:"
        while (1)
        {
            getline
            print "    - \"", $0, "\""
            if ($0 ~ "//") {
                break
            }
        }
    }
}' | \
    sed 's/Prerequisite: /prereq: "/g' | awk '
    {
        if ( $0 ~ "prereq")
        {
            print $0, "\""
        }
        else
        {
            print $0
        }
    }' | \
    sed 's/\" /"/g' | \
    sed 's/"    /"/g' | \
    sed '/\/\//d' | \
    sed 's/ "/"/g' | \
    sed 's/-"/- "/g' | \
    sed 's/:"/: "/g' > "yaml/feats.yaml"

