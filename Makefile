#COMAKE2 edit-mode: -*- Makefile -*-
####################64Bit Mode####################
ifeq ($(shell uname -m),x86_64)
CC=gcc
CXX=g++
CXXFLAGS=
CFLAGS=
CPPFLAGS=
INCPATH=
DEP_INCPATH=-I../../../../../public/go/go13 \
  -I../../../../../public/go/go13/include \
  -I../../../../../public/go/go13/output \
  -I../../../../../public/go/go13/output/include \
  -I../../../../../thirdsrc/go-lib/src \
  -I../../../../../thirdsrc/go-lib/src/include \
  -I../../../../../thirdsrc/go-lib/src/output \
  -I../../../../../thirdsrc/go-lib/src/output/include

#============ CCP vars ============
CCHECK=@ccheck.py
CCHECK_FLAGS=
PCLINT=@pclint
PCLINT_FLAGS=
CCP=@ccp.py
CCP_FLAGS=


#COMAKE UUID
COMAKE_MD5=93b991ea10579f3c685fef8d773c2cab  COMAKE


.PHONY:all
all:comake2_makefile_check .dummy 
	@echo "[[1;32;40mCOMAKE:BUILD[0m][Target:'[1;32;40mall[0m']"
	@echo "make all done"

.PHONY:comake2_makefile_check
comake2_makefile_check:
	@echo "[[1;32;40mCOMAKE:BUILD[0m][Target:'[1;32;40mcomake2_makefile_check[0m']"
	#in case of error, update 'Makefile' by 'comake2'
	@echo "$(COMAKE_MD5)">comake2.md5
	@md5sum -c --status comake2.md5
	@rm -f comake2.md5

.PHONY:ccpclean
ccpclean:
	@echo "[[1;32;40mCOMAKE:BUILD[0m][Target:'[1;32;40mccpclean[0m']"
	@echo "make ccpclean done"

.PHONY:clean
clean:ccpclean
	@echo "[[1;32;40mCOMAKE:BUILD[0m][Target:'[1;32;40mclean[0m']"
	rm -rf .dummy
	make -f Makefile.my clean

.PHONY:dist
dist:
	@echo "[[1;32;40mCOMAKE:BUILD[0m][Target:'[1;32;40mdist[0m']"
	tar czvf output.tar.gz output
	@echo "make dist done"

.PHONY:distclean
distclean:clean
	@echo "[[1;32;40mCOMAKE:BUILD[0m][Target:'[1;32;40mdistclean[0m']"
	rm -f output.tar.gz
	@echo "make distclean done"

.PHONY:love
love:
	@echo "[[1;32;40mCOMAKE:BUILD[0m][Target:'[1;32;40mlove[0m']"
	@echo "make love done"

.dummy:
	@echo "[[1;32;40mCOMAKE:BUILD[0m][Target:'[1;32;40m.dummy[0m']"
	make -j8 -f Makefile.my

endif #ifeq ($(shell uname -m),x86_64)


