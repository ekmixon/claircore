package sqlite

type tagTab struct {
	Short      string
	Val        tag
	Type       kind
	ReturnType returnType
	Extension  bool
}

func (t tagTab) Name() string {
	return t.Val.String()
}

var tagTable = [...]tagTab{
	{"Headerimage", tagHeaderImage, typeNull, returnAny, false},
	{"Headersignatures", tagHeaderSignatures, typeNull, returnAny, false},
	{"Headerimmutable", tagHeaderImmutable, typeNull, returnAny, false},
	{"Headerregions", tagHeaderRegions, typeNull, returnAny, false},
	{"Headeri18ntable", tagHeaderI18nTable, typeStringArray, returnArray, false},
	{"Sigsize", tagSigSize, typeInt32, returnScalar, false},
	{"Sigpgp", tagSigPGP, typeBin, returnScalar, false},
	{"Sigmd5", tagSigMD5, typeBin, returnScalar, false},
	{"Siggpg", tagSigGPG, typeBin, returnScalar, false},
	{"Pubkeys", tagPubKeys, typeStringArray, returnArray, false},
	{"Dsaheader", tagDSAHeader, typeBin, returnScalar, false},
	{"Rsaheader", tagRSAHeader, typeBin, returnScalar, false},
	{"Sha1header", tagSHA1Header, typeString, returnScalar, false},
	{"Longsigsize", tagLongSigSize, typeInt64, returnScalar, false},
	{"Longarchivesize", tagLongArchiveSize, typeInt64, returnScalar, false},
	{"Sha256header", tagSHA256Header, typeString, returnScalar, false},
	{"Veritysignatures", tagVeritySignatures, typeStringArray, returnArray, false},
	{"Veritysignaturealgo", tagVeritySignatureAlgo, typeInt32, returnScalar, false},
	{"Name", tagName, typeString, returnScalar, false},
	{"Version", tagVersion, typeString, returnScalar, false},
	{"Release", tagRelease, typeString, returnScalar, false},
	{"Epoch", tagEpoch, typeInt32, returnScalar, false},
	{"Summary", tagSummary, typeI18nString, returnScalar, false},
	{"Description", tagDescription, typeI18nString, returnScalar, false},
	{"Buildtime", tagBuildTime, typeInt32, returnScalar, false},
	{"Buildhost", tagBuildHost, typeString, returnScalar, false},
	{"Installtime", tagInstallTime, typeInt32, returnScalar, false},
	{"Size", tagSize, typeInt32, returnScalar, false},
	{"Distribution", tagDistribution, typeString, returnScalar, false},
	{"Vendor", tagVendor, typeString, returnScalar, false},
	{"Gif", tagGIF, typeBin, returnScalar, false},
	{"Xpm", tagXPM, typeBin, returnScalar, false},
	{"License", tagLicense, typeString, returnScalar, false},
	{"Packager", tagPackager, typeString, returnScalar, false},
	{"Group", tagGroup, typeI18nString, returnScalar, false},
	{"Source", tagSource, typeStringArray, returnArray, false},
	{"Patch", tagPatch, typeStringArray, returnArray, false},
	{"Url", tagURL, typeString, returnScalar, false},
	{"Os", tagOS, typeString, returnScalar, false},
	{"Arch", tagArch, typeString, returnScalar, false},
	{"Prein", tagPreInstall, typeString, returnScalar, false},
	{"Postin", tagPostInstall, typeString, returnScalar, false},
	{"Preun", tagPreUninstall, typeString, returnScalar, false},
	{"Postun", tagPostUninstall, typeString, returnScalar, false},
	{"Oldfilenames", tagOldFilenames, typeStringArray, returnArray, false},
	{"Filesizes", tagFileSizes, typeInt32, returnArray, false},
	{"Filestates", tagFileStates, typeChar, returnArray, false},
	{"Filemodes", tagFileModes, typeInt16, returnArray, false},
	{"Filerdevs", tagFileRDevs, typeInt16, returnArray, false},
	{"Filemtimes", tagFileMTimes, typeInt32, returnArray, false},
	{"Filedigests", tagFileDigests, typeStringArray, returnArray, false},
	{"Filelinktos", tagFileLinkTos, typeStringArray, returnArray, false},
	{"Fileflags", tagFileFlags, typeInt32, returnArray, false},
	{"Fileusername", tagFileUsername, typeStringArray, returnArray, false},
	{"Filegroupname", tagFileGroupname, typeStringArray, returnArray, false},
	{"Icon", tagIcon, typeBin, returnScalar, false},
	{"Sourcerpm", tagSourceRPM, typeString, returnScalar, false},
	{"Fileverifyflags", tagFileVerifyFlags, typeInt32, returnArray, false},
	{"Archivesize", tagArchiveSize, typeInt32, returnScalar, false},
	{"Providename", tagProvideName, typeStringArray, returnArray, false},
	{"Requireflags", tagRequireFlags, typeInt32, returnArray, false},
	{"Requirename", tagRequireName, typeStringArray, returnArray, false},
	{"Requireversion", tagRequireVersion, typeStringArray, returnArray, false},
	{"Nosource", tagNoSource, typeInt32, returnArray, false},
	{"Nopatch", tagNoPatch, typeInt32, returnArray, false},
	{"Conflictflags", tagConflictFlags, typeInt32, returnArray, false},
	{"Conflictname", tagConflictName, typeStringArray, returnArray, false},
	{"Conflictversion", tagConflictVersion, typeStringArray, returnArray, false},
	{"Excludearch", tagExcludeArch, typeStringArray, returnArray, false},
	{"Excludeos", tagExcludeOS, typeStringArray, returnArray, false},
	{"Exclusivearch", tagExclusiveArch, typeStringArray, returnArray, false},
	{"Exclusiveos", tagExclusiveOS, typeStringArray, returnArray, false},
	{"Rpmversion", tagRPMVersion, typeString, returnScalar, false},
	{"Triggerscripts", tagTriggerScripts, typeStringArray, returnArray, false},
	{"Triggername", tagTriggerName, typeStringArray, returnArray, false},
	{"Triggerversion", tagTriggerVersion, typeStringArray, returnArray, false},
	{"Triggerflags", tagTriggerFlags, typeInt32, returnArray, false},
	{"Triggerindex", tagTriggerIndex, typeInt32, returnArray, false},
	{"Verifyscript", tagVerifyScript, typeString, returnScalar, false},
	{"Changelogtime", tagChangelogTime, typeInt32, returnArray, false},
	{"Changelogname", tagChangelogName, typeStringArray, returnArray, false},
	{"Changelogtext", tagChangelogText, typeStringArray, returnArray, false},
	{"Preinprog", tagPreInstallProg, typeStringArray, returnArray, false},
	{"Postinprog", tagPostInstallProg, typeStringArray, returnArray, false},
	{"Preunprog", tagPreUninstallProg, typeStringArray, returnArray, false},
	{"Postunprog", tagPostUninstallProg, typeStringArray, returnArray, false},
	{"Buildarchs", tagBuildArchs, typeStringArray, returnArray, false},
	{"Obsoletename", tagObsoleteName, typeStringArray, returnArray, false},
	{"Verifyscriptprog", tagVerifyScriptProg, typeStringArray, returnArray, false},
	{"Triggerscriptprog", tagTriggerScriptProg, typeStringArray, returnArray, false},
	{"Cookie", tagCookie, typeString, returnScalar, false},
	{"Filedevices", tagFileDevices, typeInt32, returnArray, false},
	{"Fileinodes", tagFileInodes, typeInt32, returnArray, false},
	{"Filelangs", tagFileLangs, typeStringArray, returnArray, false},
	{"Prefixes", tagPrefixes, typeStringArray, returnArray, false},
	{"Instprefixes", tagInstallPrefixes, typeStringArray, returnArray, false},
	{"Sourcepackage", tagSourcePackage, typeInt32, returnScalar, false},
	{"Provideflags", tagProvideFlags, typeInt32, returnArray, false},
	{"Provideversion", tagProvideVersion, typeStringArray, returnArray, false},
	{"Obsoleteflags", tagObsoleteFlags, typeInt32, returnArray, false},
	{"Obsoleteversion", tagObsoleteVersion, typeStringArray, returnArray, false},
	{"Dirindexes", tagDirindexes, typeInt32, returnArray, false},
	{"Basenames", tagBasenames, typeStringArray, returnArray, false},
	{"Dirnames", tagDirnames, typeStringArray, returnArray, false},
	{"Origdirindexes", tagOrigDirindexes, typeInt32, returnArray, false},
	{"Origbasenames", tagOrigBasenames, typeStringArray, returnArray, false},
	{"Origdirnames", tagOrigDirnames, typeStringArray, returnArray, false},
	{"Optflags", tagOptFlags, typeString, returnScalar, false},
	{"Disturl", tagDistURL, typeString, returnScalar, false},
	{"Payloadformat", tagPayloadFormat, typeString, returnScalar, false},
	{"Payloadcompressor", tagPayloadCompressor, typeString, returnScalar, false},
	{"Payloadflags", tagPayloadFlags, typeString, returnScalar, false},
	{"Installcolor", tagInstallColor, typeInt32, returnScalar, false},
	{"Installtid", tagInstallTID, typeInt32, returnScalar, false},
	{"Removetid", tagRemoveTID, typeInt32, returnScalar, false},
	{"Platform", tagPlatform, typeString, returnScalar, false},
	{"Patchesname", tagPatchesName, typeStringArray, returnArray, false},
	{"Patchesflags", tagPatchesFlags, typeInt32, returnArray, false},
	{"Patchesversion", tagPatchesVersion, typeStringArray, returnArray, false},
	{"Filecolors", tagFileColors, typeInt32, returnArray, false},
	{"Fileclass", tagFileClass, typeInt32, returnArray, false},
	{"Classdict", tagClassDict, typeStringArray, returnArray, false},
	{"Filedependsx", tagFileDependsX, typeInt32, returnArray, false},
	{"Filedependsn", tagFileDependsN, typeInt32, returnArray, false},
	{"Dependsdict", tagDependsDict, typeInt32, returnArray, false},
	{"Sourcepkgid", tagSourcePkgID, typeBin, returnScalar, false},
	{"Filecontexts", tagFileContexts, typeStringArray, returnArray, false},
	{"Fscontexts", tagFSContexts, typeStringArray, returnArray, true},
	{"Recontexts", tagREContexts, typeStringArray, returnArray, true},
	{"Policies", tagPolicies, typeStringArray, returnArray, false},
	{"Pretrans", tagPreTrans, typeString, returnScalar, false},
	{"Posttrans", tagPostTrans, typeString, returnScalar, false},
	{"Pretransprog", tagPreTransProg, typeStringArray, returnArray, false},
	{"Posttransprog", tagPostTransProg, typeStringArray, returnArray, false},
	{"Disttag", tagDistTag, typeString, returnScalar, false},
	{"Oldsuggestsname", tagOldSuggestsName, typeStringArray, returnArray, false},
	{"Oldsuggestsversion", tagOldSuggestsVersion, typeStringArray, returnArray, false},
	{"Oldsuggestsflags", tagOldSuggestsFlags, typeInt32, returnArray, false},
	{"Oldenhancesname", tagOldEnhancesName, typeStringArray, returnArray, false},
	{"Oldenhancesversion", tagOldEnhancesVersion, typeStringArray, returnArray, false},
	{"Oldenhancesflags", tagOldEnhancesFlags, typeInt32, returnArray, false},
	{"Dbinstance", tagDbInstance, typeInt32, returnScalar, true},
	{"Nvra", tagNVRA, typeString, returnScalar, true},
	{"Filenames", tagFilenames, typeStringArray, returnArray, true},
	{"Fileprovide", tagFileProvide, typeStringArray, returnArray, true},
	{"Filerequire", tagFileRequire, typeStringArray, returnArray, true},
	{"Triggerconds", tagTriggerConds, typeStringArray, returnArray, true},
	{"Triggertype", tagTriggerType, typeStringArray, returnArray, true},
	{"Origfilenames", tagOrigFileNames, typeStringArray, returnArray, true},
	{"Longfilesizes", tagLongFileSizes, typeInt64, returnArray, false},
	{"Longsize", tagLongSize, typeInt64, returnScalar, false},
	{"Filecaps", tagFileCaps, typeStringArray, returnArray, false},
	{"Filedigestalgo", tagFileDigestAlgo, typeInt32, returnScalar, false},
	{"Bugurl", tagBugURL, typeString, returnScalar, false},
	{"Evr", tagEVR, typeString, returnScalar, true},
	{"Nvr", tagNVR, typeString, returnScalar, true},
	{"Nevr", tagNEVR, typeString, returnScalar, true},
	{"Nevra", tagNEVRA, typeString, returnScalar, true},
	{"Headercolor", tagHeaderColor, typeInt32, returnScalar, true},
	{"Verbose", tagVerbose, typeInt32, returnScalar, true},
	{"Epochnum", tagEpochNum, typeInt32, returnScalar, true},
	{"Preinflags", tagPreInstallFlags, typeInt32, returnScalar, false},
	{"Postinflags", tagPostInstallFlags, typeInt32, returnScalar, false},
	{"Preunflags", tagPreUninstallFlags, typeInt32, returnScalar, false},
	{"Postunflags", tagPostUninstallFlags, typeInt32, returnScalar, false},
	{"Pretransflags", tagPreTransFlags, typeInt32, returnScalar, false},
	{"Posttransflags", tagPostTransFlags, typeInt32, returnScalar, false},
	{"Verifyscriptflags", tagVerifyScriptFlags, typeInt32, returnScalar, false},
	{"Triggerscriptflags", tagTriggerScriptFlags, typeInt32, returnArray, false},
	{"Policynames", tagPolicyNames, typeStringArray, returnArray, false},
	{"Policytypes", tagPolicyTypes, typeStringArray, returnArray, false},
	{"Policytypesindexes", tagPolicyTypesIndexes, typeInt32, returnArray, false},
	{"Policyflags", tagPolicyFlags, typeInt32, returnArray, false},
	{"Vcs", tagVCS, typeString, returnScalar, false},
	{"Ordername", tagOrderName, typeStringArray, returnArray, false},
	{"Orderversion", tagOrderVersion, typeStringArray, returnArray, false},
	{"Orderflags", tagOrderFlags, typeInt32, returnArray, false},
	{"Instfilenames", tagInstFilenames, typeStringArray, returnArray, true},
	{"Requirenevrs", tagRequireNEVRS, typeStringArray, returnArray, true},
	{"Providenevrs", tagProvideNEVRS, typeStringArray, returnArray, true},
	{"Obsoletenevrs", tagObsoleteNEVRS, typeStringArray, returnArray, true},
	{"Conflictnevrs", tagConflictNEVRS, typeStringArray, returnArray, true},
	{"Filenlinks", tagFileNLinks, typeInt32, returnArray, true},
	{"Recommendname", tagRecommendName, typeStringArray, returnArray, false},
	{"Recommendversion", tagRecommendVersion, typeStringArray, returnArray, false},
	{"Recommendflags", tagRecommendFlags, typeInt32, returnArray, false},
	{"Suggestname", tagSuggestName, typeStringArray, returnArray, false},
	{"Suggestversion", tagSuggestVersion, typeStringArray, returnArray, false},
	{"Suggestflags", tagSuggestFlags, typeInt32, returnArray, false},
	{"Supplementname", tagSupplementName, typeStringArray, returnArray, false},
	{"Supplementversion", tagSupplementVersion, typeStringArray, returnArray, false},
	{"Supplementflags", tagSupplementFlags, typeInt32, returnArray, false},
	{"Enhancename", tagEnhanceName, typeStringArray, returnArray, false},
	{"Enhanceversion", tagEnhanceVersion, typeStringArray, returnArray, false},
	{"Enhanceflags", tagEnhanceFlags, typeInt32, returnArray, false},
	{"Recommendnevrs", tagRecommendNEVRS, typeStringArray, returnArray, true},
	{"Suggestnevrs", tagSuggestNEVRS, typeStringArray, returnArray, true},
	{"Supplementnevrs", tagSupplementNEVRS, typeStringArray, returnArray, true},
	{"Enhancenevrs", tagEnhanceNEVRS, typeStringArray, returnArray, true},
	{"Encoding", tagEncoding, typeString, returnScalar, false},
	{"Filetriggerscripts", tagFileTriggerScripts, typeStringArray, returnArray, false},
	{"Filetriggerscriptprog", tagFileTriggerScriptProg, typeStringArray, returnArray, false},
	{"Filetriggerscriptflags", tagFileTriggerScriptFlags, typeInt32, returnArray, false},
	{"Filetriggername", tagFileTriggerName, typeStringArray, returnArray, false},
	{"Filetriggerindex", tagFileTriggerIndex, typeInt32, returnArray, false},
	{"Filetriggerversion", tagFileTriggerVersion, typeStringArray, returnArray, false},
	{"Filetriggerflags", tagFileTriggerFlags, typeInt32, returnArray, false},
	{"Transfiletriggerscripts", tagTransFileTriggerScripts, typeStringArray, returnArray, false},
	{"Transfiletriggerscriptprog", tagTransFileTriggerScriptProg, typeStringArray, returnArray, false},
	{"Transfiletriggerscriptflags", tagTransFileTriggerScriptFlags, typeInt32, returnArray, false},
	{"Transfiletriggername", tagTransFileTriggerName, typeStringArray, returnArray, false},
	{"Transfiletriggerindex", tagTransFileTriggerIndex, typeInt32, returnArray, false},
	{"Transfiletriggerversion", tagTransFileTriggerVersion, typeStringArray, returnArray, false},
	{"Transfiletriggerflags", tagTransFileTriggerFlags, typeInt32, returnArray, false},
	{"Filetriggerpriorities", tagFileTriggerPriorities, typeInt32, returnArray, false},
	{"Transfiletriggerpriorities", tagTransFileTriggerPriorities, typeInt32, returnArray, false},
	{"Filetriggerconds", tagFileTriggerConds, typeStringArray, returnArray, true},
	{"Filetriggertype", tagFileTriggerType, typeStringArray, returnArray, true},
	{"Transfiletriggerconds", tagTransFileTriggerConds, typeStringArray, returnArray, true},
	{"Transfiletriggertype", tagTransFileTriggerType, typeStringArray, returnArray, true},
	{"Filesignatures", tagFileSignatures, typeStringArray, returnArray, false},
	{"Filesignaturelength", tagFileSignatureLength, typeInt32, returnScalar, false},
	{"Payloaddigest", tagPayloadDigest, typeStringArray, returnArray, false},
	{"Payloaddigestalgo", tagPayloadDigestAlgo, typeInt32, returnScalar, false},
	{"Modularitylabel", tagModularityLabel, typeString, returnScalar, false},
	{"Payloaddigestalt", tagPayloadDigestAlt, typeStringArray, returnArray, false},
	{"Archsuffix", tagArchSuffix, typeString, returnScalar, true},
}

var tagByValue map[tag]int

func init() {
	tagByValue = make(map[tag]int, len(tagTable))
	for i, t := range tagTable[:] {
		tagByValue[t.Val] = i
	}
}
