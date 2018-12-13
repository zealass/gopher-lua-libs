// inspect.lua for gopher-lua
package inspect

const lua_inspect = "bG9jYWwgaW5zcGVjdCA9ewogIF9WRVJTSU9OID0gJ2luc3BlY3QubHVhIDMuMS4wJywKICBfVVJMICAgICA9ICdodHRwOi8vZ2l0aHViLmNvbS9raWtpdG8vaW5zcGVjdC5sdWEnLAogIF9ERVNDUklQVElPTiA9ICdodW1hbi1yZWFkYWJsZSByZXByZXNlbnRhdGlvbnMgb2YgdGFibGVzJywKICBfTElDRU5TRSA9IFtbCiAgICBNSVQgTElDRU5TRQoKICAgIENvcHlyaWdodCAoYykgMjAxMyBFbnJpcXVlIEdhcmPDrWEgQ290YQoKICAgIFBlcm1pc3Npb24gaXMgaGVyZWJ5IGdyYW50ZWQsIGZyZWUgb2YgY2hhcmdlLCB0byBhbnkgcGVyc29uIG9idGFpbmluZyBhCiAgICBjb3B5IG9mIHRoaXMgc29mdHdhcmUgYW5kIGFzc29jaWF0ZWQgZG9jdW1lbnRhdGlvbiBmaWxlcyAodGhlCiAgICAiU29mdHdhcmUiKSwgdG8gZGVhbCBpbiB0aGUgU29mdHdhcmUgd2l0aG91dCByZXN0cmljdGlvbiwgaW5jbHVkaW5nCiAgICB3aXRob3V0IGxpbWl0YXRpb24gdGhlIHJpZ2h0cyB0byB1c2UsIGNvcHksIG1vZGlmeSwgbWVyZ2UsIHB1Ymxpc2gsCiAgICBkaXN0cmlidXRlLCBzdWJsaWNlbnNlLCBhbmQvb3Igc2VsbCBjb3BpZXMgb2YgdGhlIFNvZnR3YXJlLCBhbmQgdG8KICAgIHBlcm1pdCBwZXJzb25zIHRvIHdob20gdGhlIFNvZnR3YXJlIGlzIGZ1cm5pc2hlZCB0byBkbyBzbywgc3ViamVjdCB0bwogICAgdGhlIGZvbGxvd2luZyBjb25kaXRpb25zOgoKICAgIFRoZSBhYm92ZSBjb3B5cmlnaHQgbm90aWNlIGFuZCB0aGlzIHBlcm1pc3Npb24gbm90aWNlIHNoYWxsIGJlIGluY2x1ZGVkCiAgICBpbiBhbGwgY29waWVzIG9yIHN1YnN0YW50aWFsIHBvcnRpb25zIG9mIHRoZSBTb2Z0d2FyZS4KCiAgICBUSEUgU09GVFdBUkUgSVMgUFJPVklERUQgIkFTIElTIiwgV0lUSE9VVCBXQVJSQU5UWSBPRiBBTlkgS0lORCwgRVhQUkVTUwogICAgT1IgSU1QTElFRCwgSU5DTFVESU5HIEJVVCBOT1QgTElNSVRFRCBUTyBUSEUgV0FSUkFOVElFUyBPRgogICAgTUVSQ0hBTlRBQklMSVRZLCBGSVRORVNTIEZPUiBBIFBBUlRJQ1VMQVIgUFVSUE9TRSBBTkQgTk9OSU5GUklOR0VNRU5ULgogICAgSU4gTk8gRVZFTlQgU0hBTEwgVEhFIEFVVEhPUlMgT1IgQ09QWVJJR0hUIEhPTERFUlMgQkUgTElBQkxFIEZPUiBBTlkKICAgIENMQUlNLCBEQU1BR0VTIE9SIE9USEVSIExJQUJJTElUWSwgV0hFVEhFUiBJTiBBTiBBQ1RJT04gT0YgQ09OVFJBQ1QsCiAgICBUT1JUIE9SIE9USEVSV0lTRSwgQVJJU0lORyBGUk9NLCBPVVQgT0YgT1IgSU4gQ09OTkVDVElPTiBXSVRIIFRIRQogICAgU09GVFdBUkUgT1IgVEhFIFVTRSBPUiBPVEhFUiBERUFMSU5HUyBJTiBUSEUgU09GVFdBUkUuCiAgXV0KfQoKbG9jYWwgdG9zdHJpbmcgPSB0b3N0cmluZwoKaW5zcGVjdC5LRVkgICAgICAgPSBzZXRtZXRhdGFibGUoe30sIHtfX3Rvc3RyaW5nID0gZnVuY3Rpb24oKSByZXR1cm4gJ2luc3BlY3QuS0VZJyBlbmR9KQppbnNwZWN0Lk1FVEFUQUJMRSA9IHNldG1ldGF0YWJsZSh7fSwge19fdG9zdHJpbmcgPSBmdW5jdGlvbigpIHJldHVybiAnaW5zcGVjdC5NRVRBVEFCTEUnIGVuZH0pCgpsb2NhbCBmdW5jdGlvbiByYXdwYWlycyh0KQogIHJldHVybiBuZXh0LCB0LCBuaWwKZW5kCgotLSBBcG9zdHJvcGhpemVzIHRoZSBzdHJpbmcgaWYgaXQgaGFzIHF1b3RlcywgYnV0IG5vdCBhcGhvc3Ryb3BoZXMKLS0gT3RoZXJ3aXNlLCBpdCByZXR1cm5zIGEgcmVndWxhciBxdW90ZWQgc3RyaW5nCmxvY2FsIGZ1bmN0aW9uIHNtYXJ0UXVvdGUoc3RyKQogIGlmIHN0cjptYXRjaCgnIicpIGFuZCBub3Qgc3RyOm1hdGNoKCInIikgdGhlbgogICAgcmV0dXJuICInIiAuLiBzdHIgLi4gIiciCiAgZW5kCiAgcmV0dXJuICciJyAuLiBzdHI6Z3N1YignIicsICdcXCInKSAuLiAnIicKZW5kCgotLSBcYSA9PiAnXFxhJywgXDAgPT4gJ1xcMCcsIDMxID0+ICdcMzEnCmxvY2FsIHNob3J0Q29udHJvbENoYXJFc2NhcGVzID0gewogIFsiXGEiXSA9ICJcXGEiLCAgWyJcYiJdID0gIlxcYiIsIFsiXGYiXSA9ICJcXGYiLCBbIlxuIl0gPSAiXFxuIiwKICBbIlxyIl0gPSAiXFxyIiwgIFsiXHQiXSA9ICJcXHQiLCBbIlx2Il0gPSAiXFx2Igp9CmxvY2FsIGxvbmdDb250cm9sQ2hhckVzY2FwZXMgPSB7fSAtLSBcYSA9PiBuaWwsIFwwID0+IFwwMDAsIDMxID0+IFwwMzEKZm9yIGk9MCwgMzEgZG8KICBsb2NhbCBjaCA9IHN0cmluZy5jaGFyKGkpCiAgaWYgbm90IHNob3J0Q29udHJvbENoYXJFc2NhcGVzW2NoXSB0aGVuCiAgICBzaG9ydENvbnRyb2xDaGFyRXNjYXBlc1tjaF0gPSAiXFwiLi5pCiAgICBsb25nQ29udHJvbENoYXJFc2NhcGVzW2NoXSAgPSBzdHJpbmcuZm9ybWF0KCJcXCUwM2QiLCBpKQogIGVuZAplbmQKCmxvY2FsIGZ1bmN0aW9uIGVzY2FwZShzdHIpCiAgcmV0dXJuIChzdHI6Z3N1YigiXFwiLCAiXFxcXCIpCiAgICAgICAgICAgICA6Z3N1YigiKCVjKSVmWzAtOV0iLCBsb25nQ29udHJvbENoYXJFc2NhcGVzKQogICAgICAgICAgICAgOmdzdWIoIiVjIiwgc2hvcnRDb250cm9sQ2hhckVzY2FwZXMpKQplbmQKCmxvY2FsIGZ1bmN0aW9uIGlzSWRlbnRpZmllcihzdHIpCiAgcmV0dXJuIHR5cGUoc3RyKSA9PSAnc3RyaW5nJyBhbmQgc3RyOm1hdGNoKCAiXltfJWFdW18lYSVkXSokIiApCmVuZAoKbG9jYWwgZnVuY3Rpb24gaXNTZXF1ZW5jZUtleShrLCBzZXF1ZW5jZUxlbmd0aCkKICByZXR1cm4gdHlwZShrKSA9PSAnbnVtYmVyJwogICAgIGFuZCAxIDw9IGsKICAgICBhbmQgayA8PSBzZXF1ZW5jZUxlbmd0aAogICAgIGFuZCBtYXRoLmZsb29yKGspID09IGsKZW5kCgpsb2NhbCBkZWZhdWx0VHlwZU9yZGVycyA9IHsKICBbJ251bWJlciddICAgPSAxLCBbJ2Jvb2xlYW4nXSAgPSAyLCBbJ3N0cmluZyddID0gMywgWyd0YWJsZSddID0gNCwKICBbJ2Z1bmN0aW9uJ10gPSA1LCBbJ3VzZXJkYXRhJ10gPSA2LCBbJ3RocmVhZCddID0gNwp9Cgpsb2NhbCBmdW5jdGlvbiBzb3J0S2V5cyhhLCBiKQogIGxvY2FsIHRhLCB0YiA9IHR5cGUoYSksIHR5cGUoYikKCiAgLS0gc3RyaW5ncyBhbmQgbnVtYmVycyBhcmUgc29ydGVkIG51bWVyaWNhbGx5L2FscGhhYmV0aWNhbGx5CiAgaWYgdGEgPT0gdGIgYW5kICh0YSA9PSAnc3RyaW5nJyBvciB0YSA9PSAnbnVtYmVyJykgdGhlbiByZXR1cm4gYSA8IGIgZW5kCgogIGxvY2FsIGR0YSwgZHRiID0gZGVmYXVsdFR5cGVPcmRlcnNbdGFdLCBkZWZhdWx0VHlwZU9yZGVyc1t0Yl0KICAtLSBUd28gZGVmYXVsdCB0eXBlcyBhcmUgY29tcGFyZWQgYWNjb3JkaW5nIHRvIHRoZSBkZWZhdWx0VHlwZU9yZGVycyB0YWJsZQogIGlmIGR0YSBhbmQgZHRiIHRoZW4gcmV0dXJuIGRlZmF1bHRUeXBlT3JkZXJzW3RhXSA8IGRlZmF1bHRUeXBlT3JkZXJzW3RiXQogIGVsc2VpZiBkdGEgICAgIHRoZW4gcmV0dXJuIHRydWUgIC0tIGRlZmF1bHQgdHlwZXMgYmVmb3JlIGN1c3RvbSBvbmVzCiAgZWxzZWlmIGR0YiAgICAgdGhlbiByZXR1cm4gZmFsc2UgLS0gY3VzdG9tIHR5cGVzIGFmdGVyIGRlZmF1bHQgb25lcwogIGVuZAoKICAtLSBjdXN0b20gdHlwZXMgYXJlIHNvcnRlZCBvdXQgYWxwaGFiZXRpY2FsbHkKICByZXR1cm4gdGEgPCB0YgplbmQKCi0tIEZvciBpbXBsZW1lbnRhdGlvbiByZWFzb25zLCB0aGUgYmVoYXZpb3Igb2YgcmF3bGVuICYgIyBpcyAidW5kZWZpbmVkIiB3aGVuCi0tIHRhYmxlcyBhcmVuJ3QgcHVyZSBzZXF1ZW5jZXMuIFNvIHdlIGltcGxlbWVudCBvdXIgb3duICMgb3BlcmF0b3IuCmxvY2FsIGZ1bmN0aW9uIGdldFNlcXVlbmNlTGVuZ3RoKHQpCiAgbG9jYWwgbGVuID0gMQogIGxvY2FsIHYgPSByYXdnZXQodCxsZW4pCiAgd2hpbGUgdiB+PSBuaWwgZG8KICAgIGxlbiA9IGxlbiArIDEKICAgIHYgPSByYXdnZXQodCxsZW4pCiAgZW5kCiAgcmV0dXJuIGxlbiAtIDEKZW5kCgpsb2NhbCBmdW5jdGlvbiBnZXROb25TZXF1ZW50aWFsS2V5cyh0KQogIGxvY2FsIGtleXMsIGtleXNMZW5ndGggPSB7fSwgMAogIGxvY2FsIHNlcXVlbmNlTGVuZ3RoID0gZ2V0U2VxdWVuY2VMZW5ndGgodCkKICBmb3IgayxfIGluIHJhd3BhaXJzKHQpIGRvCiAgICBpZiBub3QgaXNTZXF1ZW5jZUtleShrLCBzZXF1ZW5jZUxlbmd0aCkgdGhlbgogICAgICBrZXlzTGVuZ3RoID0ga2V5c0xlbmd0aCArIDEKICAgICAga2V5c1trZXlzTGVuZ3RoXSA9IGsKICAgIGVuZAogIGVuZAogIHRhYmxlLnNvcnQoa2V5cywgc29ydEtleXMpCiAgcmV0dXJuIGtleXMsIGtleXNMZW5ndGgsIHNlcXVlbmNlTGVuZ3RoCmVuZAoKbG9jYWwgZnVuY3Rpb24gY291bnRUYWJsZUFwcGVhcmFuY2VzKHQsIHRhYmxlQXBwZWFyYW5jZXMpCiAgdGFibGVBcHBlYXJhbmNlcyA9IHRhYmxlQXBwZWFyYW5jZXMgb3Ige30KCiAgaWYgdHlwZSh0KSA9PSAndGFibGUnIHRoZW4KICAgIGlmIG5vdCB0YWJsZUFwcGVhcmFuY2VzW3RdIHRoZW4KICAgICAgdGFibGVBcHBlYXJhbmNlc1t0XSA9IDEKICAgICAgZm9yIGssdiBpbiByYXdwYWlycyh0KSBkbwogICAgICAgIGNvdW50VGFibGVBcHBlYXJhbmNlcyhrLCB0YWJsZUFwcGVhcmFuY2VzKQogICAgICAgIGNvdW50VGFibGVBcHBlYXJhbmNlcyh2LCB0YWJsZUFwcGVhcmFuY2VzKQogICAgICBlbmQKICAgICAgY291bnRUYWJsZUFwcGVhcmFuY2VzKGdldG1ldGF0YWJsZSh0KSwgdGFibGVBcHBlYXJhbmNlcykKICAgIGVsc2UKICAgICAgdGFibGVBcHBlYXJhbmNlc1t0XSA9IHRhYmxlQXBwZWFyYW5jZXNbdF0gKyAxCiAgICBlbmQKICBlbmQKCiAgcmV0dXJuIHRhYmxlQXBwZWFyYW5jZXMKZW5kCgpsb2NhbCBjb3B5U2VxdWVuY2UgPSBmdW5jdGlvbihzKQogIGxvY2FsIGNvcHksIGxlbiA9IHt9LCAjcwogIGZvciBpPTEsIGxlbiBkbyBjb3B5W2ldID0gc1tpXSBlbmQKICByZXR1cm4gY29weSwgbGVuCmVuZAoKbG9jYWwgZnVuY3Rpb24gbWFrZVBhdGgocGF0aCwgLi4uKQogIGxvY2FsIGtleXMgPSB7Li4ufQogIGxvY2FsIG5ld1BhdGgsIGxlbiA9IGNvcHlTZXF1ZW5jZShwYXRoKQogIGZvciBpPTEsICNrZXlzIGRvCiAgICBuZXdQYXRoW2xlbiArIGldID0ga2V5c1tpXQogIGVuZAogIHJldHVybiBuZXdQYXRoCmVuZAoKbG9jYWwgZnVuY3Rpb24gcHJvY2Vzc1JlY3Vyc2l2ZShwcm9jZXNzLCBpdGVtLCBwYXRoLCB2aXNpdGVkKQogIGlmIGl0ZW0gPT0gbmlsIHRoZW4gcmV0dXJuIG5pbCBlbmQKICBpZiB2aXNpdGVkW2l0ZW1dIHRoZW4gcmV0dXJuIHZpc2l0ZWRbaXRlbV0gZW5kCgogIGxvY2FsIHByb2Nlc3NlZCA9IHByb2Nlc3MoaXRlbSwgcGF0aCkKICBpZiB0eXBlKHByb2Nlc3NlZCkgPT0gJ3RhYmxlJyB0aGVuCiAgICBsb2NhbCBwcm9jZXNzZWRDb3B5ID0ge30KICAgIHZpc2l0ZWRbaXRlbV0gPSBwcm9jZXNzZWRDb3B5CiAgICBsb2NhbCBwcm9jZXNzZWRLZXkKCiAgICBmb3Igayx2IGluIHJhd3BhaXJzKHByb2Nlc3NlZCkgZG8KICAgICAgcHJvY2Vzc2VkS2V5ID0gcHJvY2Vzc1JlY3Vyc2l2ZShwcm9jZXNzLCBrLCBtYWtlUGF0aChwYXRoLCBrLCBpbnNwZWN0LktFWSksIHZpc2l0ZWQpCiAgICAgIGlmIHByb2Nlc3NlZEtleSB+PSBuaWwgdGhlbgogICAgICAgIHByb2Nlc3NlZENvcHlbcHJvY2Vzc2VkS2V5XSA9IHByb2Nlc3NSZWN1cnNpdmUocHJvY2VzcywgdiwgbWFrZVBhdGgocGF0aCwgcHJvY2Vzc2VkS2V5KSwgdmlzaXRlZCkKICAgICAgZW5kCiAgICBlbmQKCiAgICBsb2NhbCBtdCAgPSBwcm9jZXNzUmVjdXJzaXZlKHByb2Nlc3MsIGdldG1ldGF0YWJsZShwcm9jZXNzZWQpLCBtYWtlUGF0aChwYXRoLCBpbnNwZWN0Lk1FVEFUQUJMRSksIHZpc2l0ZWQpCiAgICBpZiB0eXBlKG10KSB+PSAndGFibGUnIHRoZW4gbXQgPSBuaWwgZW5kIC0tIGlnbm9yZSBub3QgbmlsL3RhYmxlIF9fbWV0YXRhYmxlIGZpZWxkCiAgICBzZXRtZXRhdGFibGUocHJvY2Vzc2VkQ29weSwgbXQpCiAgICBwcm9jZXNzZWQgPSBwcm9jZXNzZWRDb3B5CiAgZW5kCiAgcmV0dXJuIHByb2Nlc3NlZAplbmQKCgoKLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLQoKbG9jYWwgSW5zcGVjdG9yID0ge30KbG9jYWwgSW5zcGVjdG9yX210ID0ge19faW5kZXggPSBJbnNwZWN0b3J9CgpmdW5jdGlvbiBJbnNwZWN0b3I6cHV0cyguLi4pCiAgbG9jYWwgYXJncyAgID0gey4uLn0KICBsb2NhbCBidWZmZXIgPSBzZWxmLmJ1ZmZlcgogIGxvY2FsIGxlbiAgICA9ICNidWZmZXIKICBmb3IgaT0xLCAjYXJncyBkbwogICAgbGVuID0gbGVuICsgMQogICAgYnVmZmVyW2xlbl0gPSBhcmdzW2ldCiAgZW5kCmVuZAoKZnVuY3Rpb24gSW5zcGVjdG9yOmRvd24oZikKICBzZWxmLmxldmVsID0gc2VsZi5sZXZlbCArIDEKICBmKCkKICBzZWxmLmxldmVsID0gc2VsZi5sZXZlbCAtIDEKZW5kCgpmdW5jdGlvbiBJbnNwZWN0b3I6dGFiaWZ5KCkKICBzZWxmOnB1dHMoc2VsZi5uZXdsaW5lLCBzdHJpbmcucmVwKHNlbGYuaW5kZW50LCBzZWxmLmxldmVsKSkKZW5kCgpmdW5jdGlvbiBJbnNwZWN0b3I6YWxyZWFkeVZpc2l0ZWQodikKICByZXR1cm4gc2VsZi5pZHNbdl0gfj0gbmlsCmVuZAoKZnVuY3Rpb24gSW5zcGVjdG9yOmdldElkKHYpCiAgbG9jYWwgaWQgPSBzZWxmLmlkc1t2XQogIGlmIG5vdCBpZCB0aGVuCiAgICBsb2NhbCB0diA9IHR5cGUodikKICAgIGlkICAgICAgICAgICAgICA9IChzZWxmLm1heElkc1t0dl0gb3IgMCkgKyAxCiAgICBzZWxmLm1heElkc1t0dl0gPSBpZAogICAgc2VsZi5pZHNbdl0gICAgID0gaWQKICBlbmQKICByZXR1cm4gdG9zdHJpbmcoaWQpCmVuZAoKZnVuY3Rpb24gSW5zcGVjdG9yOnB1dEtleShrKQogIGlmIGlzSWRlbnRpZmllcihrKSB0aGVuIHJldHVybiBzZWxmOnB1dHMoaykgZW5kCiAgc2VsZjpwdXRzKCJbIikKICBzZWxmOnB1dFZhbHVlKGspCiAgc2VsZjpwdXRzKCJdIikKZW5kCgpmdW5jdGlvbiBJbnNwZWN0b3I6cHV0VGFibGUodCkKICBpZiB0ID09IGluc3BlY3QuS0VZIG9yIHQgPT0gaW5zcGVjdC5NRVRBVEFCTEUgdGhlbgogICAgc2VsZjpwdXRzKHRvc3RyaW5nKHQpKQogIGVsc2VpZiBzZWxmOmFscmVhZHlWaXNpdGVkKHQpIHRoZW4KICAgIHNlbGY6cHV0cygnPHRhYmxlICcsIHNlbGY6Z2V0SWQodCksICc+JykKICBlbHNlaWYgc2VsZi5sZXZlbCA+PSBzZWxmLmRlcHRoIHRoZW4KICAgIHNlbGY6cHV0cygney4uLn0nKQogIGVsc2UKICAgIGlmIHNlbGYudGFibGVBcHBlYXJhbmNlc1t0XSA+IDEgdGhlbiBzZWxmOnB1dHMoJzwnLCBzZWxmOmdldElkKHQpLCAnPicpIGVuZAoKICAgIGxvY2FsIG5vblNlcXVlbnRpYWxLZXlzLCBub25TZXF1ZW50aWFsS2V5c0xlbmd0aCwgc2VxdWVuY2VMZW5ndGggPSBnZXROb25TZXF1ZW50aWFsS2V5cyh0KQogICAgbG9jYWwgbXQgICAgICAgICAgICAgICAgPSBnZXRtZXRhdGFibGUodCkKCiAgICBzZWxmOnB1dHMoJ3snKQogICAgc2VsZjpkb3duKGZ1bmN0aW9uKCkKICAgICAgbG9jYWwgY291bnQgPSAwCiAgICAgIGZvciBpPTEsIHNlcXVlbmNlTGVuZ3RoIGRvCiAgICAgICAgaWYgY291bnQgPiAwIHRoZW4gc2VsZjpwdXRzKCcsJykgZW5kCiAgICAgICAgc2VsZjpwdXRzKCcgJykKICAgICAgICBzZWxmOnB1dFZhbHVlKHRbaV0pCiAgICAgICAgY291bnQgPSBjb3VudCArIDEKICAgICAgZW5kCgogICAgICBmb3IgaT0xLCBub25TZXF1ZW50aWFsS2V5c0xlbmd0aCBkbwogICAgICAgIGxvY2FsIGsgPSBub25TZXF1ZW50aWFsS2V5c1tpXQogICAgICAgIGlmIGNvdW50ID4gMCB0aGVuIHNlbGY6cHV0cygnLCcpIGVuZAogICAgICAgIHNlbGY6dGFiaWZ5KCkKICAgICAgICBzZWxmOnB1dEtleShrKQogICAgICAgIHNlbGY6cHV0cygnID0gJykKICAgICAgICBzZWxmOnB1dFZhbHVlKHRba10pCiAgICAgICAgY291bnQgPSBjb3VudCArIDEKICAgICAgZW5kCgogICAgICBpZiB0eXBlKG10KSA9PSAndGFibGUnIHRoZW4KICAgICAgICBpZiBjb3VudCA+IDAgdGhlbiBzZWxmOnB1dHMoJywnKSBlbmQKICAgICAgICBzZWxmOnRhYmlmeSgpCiAgICAgICAgc2VsZjpwdXRzKCc8bWV0YXRhYmxlPiA9ICcpCiAgICAgICAgc2VsZjpwdXRWYWx1ZShtdCkKICAgICAgZW5kCiAgICBlbmQpCgogICAgaWYgbm9uU2VxdWVudGlhbEtleXNMZW5ndGggPiAwIG9yIHR5cGUobXQpID09ICd0YWJsZScgdGhlbiAtLSByZXN1bHQgaXMgbXVsdGktbGluZWQuIEp1c3RpZnkgY2xvc2luZyB9CiAgICAgIHNlbGY6dGFiaWZ5KCkKICAgIGVsc2VpZiBzZXF1ZW5jZUxlbmd0aCA+IDAgdGhlbiAtLSBhcnJheSB0YWJsZXMgaGF2ZSBvbmUgZXh0cmEgc3BhY2UgYmVmb3JlIGNsb3NpbmcgfQogICAgICBzZWxmOnB1dHMoJyAnKQogICAgZW5kCgogICAgc2VsZjpwdXRzKCd9JykKICBlbmQKZW5kCgpmdW5jdGlvbiBJbnNwZWN0b3I6cHV0VmFsdWUodikKICBsb2NhbCB0diA9IHR5cGUodikKCiAgaWYgdHYgPT0gJ3N0cmluZycgdGhlbgogICAgc2VsZjpwdXRzKHNtYXJ0UXVvdGUoZXNjYXBlKHYpKSkKICBlbHNlaWYgdHYgPT0gJ251bWJlcicgb3IgdHYgPT0gJ2Jvb2xlYW4nIG9yIHR2ID09ICduaWwnIG9yCiAgICAgICAgIHR2ID09ICdjZGF0YScgb3IgdHYgPT0gJ2N0eXBlJyB0aGVuCiAgICBzZWxmOnB1dHModG9zdHJpbmcodikpCiAgZWxzZWlmIHR2ID09ICd0YWJsZScgdGhlbgogICAgc2VsZjpwdXRUYWJsZSh2KQogIGVsc2UKICAgIHNlbGY6cHV0cygnPCcsIHR2LCAnICcsIHNlbGY6Z2V0SWQodiksICc+JykKICBlbmQKZW5kCgotLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tCgpmdW5jdGlvbiBpbnNwZWN0Lmluc3BlY3Qocm9vdCwgb3B0aW9ucykKICBvcHRpb25zICAgICAgID0gb3B0aW9ucyBvciB7fQoKICBsb2NhbCBkZXB0aCAgID0gb3B0aW9ucy5kZXB0aCAgIG9yIG1hdGguaHVnZQogIGxvY2FsIG5ld2xpbmUgPSBvcHRpb25zLm5ld2xpbmUgb3IgJ1xuJwogIGxvY2FsIGluZGVudCAgPSBvcHRpb25zLmluZGVudCAgb3IgJyAgJwogIGxvY2FsIHByb2Nlc3MgPSBvcHRpb25zLnByb2Nlc3MKCiAgaWYgcHJvY2VzcyB0aGVuCiAgICByb290ID0gcHJvY2Vzc1JlY3Vyc2l2ZShwcm9jZXNzLCByb290LCB7fSwge30pCiAgZW5kCgogIGxvY2FsIGluc3BlY3RvciA9IHNldG1ldGF0YWJsZSh7CiAgICBkZXB0aCAgICAgICAgICAgID0gZGVwdGgsCiAgICBsZXZlbCAgICAgICAgICAgID0gMCwKICAgIGJ1ZmZlciAgICAgICAgICAgPSB7fSwKICAgIGlkcyAgICAgICAgICAgICAgPSB7fSwKICAgIG1heElkcyAgICAgICAgICAgPSB7fSwKICAgIG5ld2xpbmUgICAgICAgICAgPSBuZXdsaW5lLAogICAgaW5kZW50ICAgICAgICAgICA9IGluZGVudCwKICAgIHRhYmxlQXBwZWFyYW5jZXMgPSBjb3VudFRhYmxlQXBwZWFyYW5jZXMocm9vdCkKICB9LCBJbnNwZWN0b3JfbXQpCgogIGluc3BlY3RvcjpwdXRWYWx1ZShyb290KQoKICByZXR1cm4gdGFibGUuY29uY2F0KGluc3BlY3Rvci5idWZmZXIpCmVuZAoKc2V0bWV0YXRhYmxlKGluc3BlY3QsIHsgX19jYWxsID0gZnVuY3Rpb24oXywgLi4uKSByZXR1cm4gaW5zcGVjdC5pbnNwZWN0KC4uLikgZW5kIH0pCgpyZXR1cm4gaW5zcGVjdAo="